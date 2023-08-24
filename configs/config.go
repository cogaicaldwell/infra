package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	MongoDSN string `mapstructure:"MONGO_DSN"`
}

func LoadConfig(isDebugMode bool) *Config {
	viper.AddConfigPath("./configs")

	var fileName string
	if isDebugMode {
		fileName = "dev"
	} else {
		fileName = "prod"
	}

	viper.SetConfigName(fileName)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("ERROR: %c", err)
	}

	var c Config
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("ERROR: %c", err)
	}

	return &c
}
