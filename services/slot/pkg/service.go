package pkg

import (
	"net/http"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
	localconfig "github.com/JohnnyS318/RoyalAfgInGo/services/slot/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/services/slot/pkg/crypto"
	"github.com/JohnnyS318/RoyalAfgInGo/services/slot/pkg/database"
	"github.com/JohnnyS318/RoyalAfgInGo/services/slot/pkg/handlers"
	"github.com/JohnnyS318/RoyalAfgInGo/services/slot/pkg/logic"
	"github.com/Kamva/mgm"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func Start(logger *zap.SugaredLogger) {

	viper.SetEnvPrefix("slot")

	// Game Database
	cfg := &mgm.Config{CtxTimeout: viper.GetDuration(localconfig.DatabaseTimeout)}
	err := mgm.SetDefaultConfig(cfg, viper.GetString(localconfig.DatabaseName), options.Client().ApplyURI(viper.GetString(localconfig.DatabaseUrl)))

	if err != nil {
		logger.Fatalw("Could not read the mongodb config", "error", err)
	}

	_, client, _, err := mgm.DefaultConfigs()
	if err != nil {
		logger.Fatalw("Connection to mongodb failed", "error", err)
	}

	defer utils.DisconnectClient(logger, client)
	logger.Debugf("Database connection established to [%v] with database name [%v]", viper.GetString(localconfig.DatabaseUrl), viper.GetString(localconfig.DatabaseName))

	gameDatabase, err := database.NewGameDatabase(logger)

	if err != nil {
		logger.Fatalw("Could not create game database", "error", err)
	}

	// ############### Game Databank buffer ################

	buffer := database.NewGameBuffer(func(games []*models.SlotGame) error {
		err := gameDatabase.SaveGameBuffer(games)
		if err != nil {
			logger.Errorw("Could not save game buffer", "error", err)
		}
		return err
	})

	// ############### User Service Connection ################

	conn, err := grpc.Dial(viper.GetString(localconfig.UserServiceURL), grpc.WithInsecure())

	defer conn.Close()

	if err != nil {
		logger.Fatalw("Could not connect to user service", "error", err)
	}

	userServiceClient := protos.NewUserServiceClient(conn)

	// ############### Crypto keys and rng ################

	// Read crypto keys
	privateKey, publicKey, err := crypto.ReadECDSAKeys(viper.GetString(localconfig.PublicKeyPath), viper.GetString(localconfig.PrivateKeyPath))

	if err != nil {
		logger.Fatalw("Error reading crypto keys", "error", err)
	}

	// Create the crypto logic
	rng := crypto.NewVRFNumberGenerator(privateKey, publicKey)

	// ############### Game Provider ################

	gameProvider := logic.NewGameProvider(buffer, gameDatabase, rng)

	slotHandler := handlers.NewSlotServer(logger, gameProvider, userServiceClient)

	r := mux.NewRouter()

	r.HandleFunc("/api/games/slot/spin", slotHandler.Spin).Methods("POST")

	n := negroni.New(negroni.NewRecovery(), mw.NewLogger(logger.Desugar()))

	n.UseHandler(r)

	port := viper.GetString(config.HTTPPort)
	logger.Warnf("HTTP Port set to %v", port)
	srv := &http.Server{
		Addr:              ":" + port,
		WriteTimeout:      viper.GetDuration(config.WriteTimeout),
		ReadHeaderTimeout: viper.GetDuration(config.ReadTimeout),
		IdleTimeout:       viper.GetDuration(config.IdleTimeout),
		Handler:           n,
	}

	utils.StartGracefully(logger, srv, viper.GetDuration(config.GracefulShutdownTimeout))
}
