package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Mysql  Mysql
	Redis  Redis
	Server Server
	Jwt    Jwt
	Mongo  Mongo
}

var EnvConfig *Config

func LoadConfig() *Config {

	path, err := os.Getwd() // get curent path
	if err != nil {
		panic(err)
	}
	log.Println(path)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path + "/configs")

	if err := viper.ReadInConfig(); err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	// export config
	EnvConfig = config

	return config
}
