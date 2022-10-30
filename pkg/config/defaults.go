package config

import (
	"time"

	"github.com/spf13/viper"
)

const Prod = "prod"

const HTTPPort = "http_port"

const JWTSigningKey = "jwt_signing_key"
const JWTIssuer = "jwt_issuer"
const JWTExpiresAt = "jwt_expiring_time"

const SessionCookieName = "session_cookie_name"
const SessionCookieExpiration = "session_cookie_expiration"

const CorsEnabled = "cors_enabled"

const RedisAddress = "redis_address"
const RedisUsername = "redis_username"
const RedisPassword = "redis_password"

const StatusRedisAddress = "status_redis_address"
const StatusRedisUsername = "status_redis_username"
const StatusRedisPassword = "status_redis_password"

const MatchMakerJWTKey = "matchmaker_signing_key"

func RegisterDefaults() {

	viper.SetDefault(Prod, false)

	viper.SetDefault(HTTPPort, 8080)

	viper.SetDefault(JWTIssuer, "royalafg.games")
	viper.SetDefault(JWTExpiresAt, time.Hour*24*7)

	viper.SetDefault(SessionCookieName, "royalafg.session")
	viper.SetDefault(SessionCookieExpiration, time.Hour*24*7)

	viper.SetDefault(CorsEnabled, false)

	viper.SetDefault(RedisUsername, "")
	viper.SetDefault(RedisPassword, "")

	RegisterRabbitDefaults()
}
