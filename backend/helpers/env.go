package helpers

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

var Env, _ = LoadConfig(".")

type Config struct {
	ServerAddress     string        `mapstructure:"SERVER_ADDRESS"`
	DatabaseName      string        `mapstructure:"DATABASE_NAME"`
	PGDatabaseDriver  string        `mapstructure:"PG_DATABASE_DRIVER"`
	PGDatabaseInitURL string        `mapstructure:"PG_DATABASE_INIT_URL"`
	PGDatabaseURL     string        `mapstructure:"PG_DATABASE_URL"`
	TokenSecret       string        `mapstructure:"TOKEN_SECRET"`
	TokenIssuer       string        `mapstructure:"TOKEN_ISSUER"`
	TokenDuration     time.Duration `mapstructure:"TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	if os.Getenv("ENV") == "PRODUCTION" {
		viper.SetConfigName("local.production")
	} else {
		viper.SetConfigName("local.development")
	}
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
