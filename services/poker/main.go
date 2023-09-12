package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	coresdk "agones.dev/agones/pkg/sdk"
	sdk "agones.dev/agones/sdks/go"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
	"golang.org/x/net/context"

	"github.com/jonsch318/royalafg/pkg/config"
	"github.com/jonsch318/royalafg/pkg/log"
	"github.com/jonsch318/royalafg/pkg/mw"
	pokerModels "github.com/jonsch318/royalafg/pkg/poker/models"
	"github.com/jonsch318/royalafg/pkg/utils"
	"github.com/jonsch318/royalafg/services/poker/bank"
	"github.com/jonsch318/royalafg/services/poker/gameServer"
	"github.com/jonsch318/royalafg/services/poker/handlers"
	"github.com/jonsch318/royalafg/services/poker/lobby"
	"github.com/jonsch318/royalafg/services/poker/rabbit"
	"github.com/jonsch318/royalafg/services/poker/serviceconfig"
)

// main method is the entry point of the game server
func main() {
	//Register Logger
	logger := log.RegisterService()
	defer log.CleanLogger()

	//Configure
	config.ReadStandardConfig("poker", logger)
	serviceconfig.SetDefaults()

	//Bind environment variables
	viper.SetEnvPrefix("poker")
	_ = viper.BindEnv(config.RabbitMQUsername)
	_ = viper.BindEnv(config.RabbitMQPassword)

	//connect to rabbitmq to send user commands
	rabbitURL := fmt.Sprintf("amqp://%s:%s@%s", viper.GetString(config.RabbitMQUsername), viper.GetString(config.RabbitMQPassword), viper.GetString(config.RabbitMQUrl))
	rabbitConn, err := rabbit.NewRabbitMessageBroker(logger, rabbitURL)
	if err != nil {
		logger.Fatalw("Could not connect to service bus", "error", err)
	}
	defer rabbitConn.Close()

	//Register stop signal
	go gameServer.DoSignal()

	//Creating agones sdk instance to communicate with the game server orchestrator
	logger.Infof("Creating SDK instance")
	s, err := sdk.NewSDK()
	if err != nil {
		logger.Fatalf("Error during sdk connection, %v", err)
	}

	//StartHealth ping to agones.
	logger.Info("Health Ping to agones server management")
	stop := make(chan struct{})
	go gameServer.DoHealthPing(s, stop)

	//Configure logging
	lobbyConfigured := false
	shutDownStop := make(chan interface{})
	b := bank.NewBank(rabbitConn)
	lobbyInstance := lobby.NewLobby(b, s)

	//Watch Pod Labels for lobby information
	err = s.WatchGameServer(func(gs *coresdk.GameServer) {
		if !lobbyConfigured {
			err2 := SetLobby(b, lobbyInstance, gs)
			if err2 == nil {
				logger.Warnw("Lobby configured", "id", lobbyInstance.LobbyID)
				if lobbyInstance.Count() <= 0 {
					go StartShutdownTimer(shutDownStop, s)
				}
				//Lobby is configured through kubernetes labels and is assigned a unique id.
				lobbyConfigured = true
			} else if !strings.HasPrefix(err2.Error(), "key needed") {
				logger.Errorw("Error during configuration", "error", err2)
			}
		}
	})
	if err != nil {
		logger.Fatalf("Error during sdk annotation subscription: %s", err)
	}

	//Register HTTP handlers
	gameHandler := handlers.NewGame(lobbyInstance, s, shutDownStop)

	//Setup Routes
	r := mux.NewRouter()
	r.HandleFunc("/api/poker/health", gameHandler.Health).Methods(http.MethodGet)
	r.HandleFunc("/api/poker/join", gameHandler.Join) //Websocket join

	//RegisterMiddleware
	n := negroni.New(negroni.NewRecovery(), mw.NewLogger(logger.Desugar()))
	n.UseHandler(r)

	//Call Ready
	err = s.Ready()
	if err != nil {
		logger.Errorw("Error during sdk ready call", "error", err)
	}
	logger.Info("SDK Ready called")

	//Configure HTTP server
	port := viper.GetString(config.HTTPPort)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: n,
	}

	//Start HTTP server
	utils.StartGracefully(logger, srv, viper.GetDuration(config.GracefulShutdownTimeout))

}

func SetLobby(b *bank.Bank, lobbyInstance *lobby.Lobby, gs *coresdk.GameServer) error {
	labels := gs.GetObjectMeta().GetLabels()
	min, err := GetFromLabels("min-buy-in", labels)
	if err != nil {
		return err
	}

	max, err := GetFromLabels("max-buy-in", labels)
	if err != nil {
		return err
	}

	blind, err := GetFromLabels("blind", labels)
	if err != nil {
		return err
	}

	index, err := GetFromLabels("class-index", labels)
	if err != nil {
		return err
	}

	lobbyID, ok := labels["lobbyId"]
	if !ok {
		return fmt.Errorf("key needed [%v]", "lobbyId")
	}
	b.RegisterLobby(lobbyID)

	lobbyInstance.RegisterLobbyValue(&pokerModels.Class{
		Min:   min,
		Max:   max,
		Blind: blind,
	}, index, lobbyID)
	return nil
}

func GetFromLabels(key string, labels map[string]string) (int, error) {
	valString, ok := labels[key]

	if !ok {
		return 0, fmt.Errorf("key needed [%v]", key)
	}

	val, err := strconv.Atoi(valString)
	if err != nil {
		return 0, err
	}

	return val, nil
}

// StartShutdownTimer sets the status shutdown after 10 minutes.
func StartShutdownTimer(stop chan interface{}, s *sdk.SDK) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Minute))
	select {
	case _ = <-stop:
		cancel()
		//Cancel shutdown
	case <-ctx.Done():
		_ = s.Shutdown()
	}
}
