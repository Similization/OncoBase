package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type ConfigDatabase struct {
	Port     string `yml:"port" env:"PORT" env-default:"5432"`
	Host     string `yml:"host" env:"HOST" env-default:"localhost"`
	Name     string `yml:"name" env:"NAME" env-default:"postgres"`
	User     string `yml:"user" env:"USER" env-default:"user"`
	Password string `yml:"password" env:"PASSWORD"`
}

type ConfigServer struct {
	Port string `yml:"port" env:"PORT" env-default:"8080"`
	Host string `yml:"host" env:"HOST" env-default:"localhost"`
}

type ConfigApp struct {
	Database ConfigDatabase
	Server   ConfigServer
}

type ConfigInfo struct {
	Name      string
	Extension string
	Paths     []string
}

func NewConfigInfo() *ConfigInfo {
	return &ConfigInfo{
		Name:      "config",
		Extension: "yaml",
		Paths:     []string{".", "configs"},
	}
}

func InitConfig(configInfo ConfigInfo) *ConfigApp {
	viper.SetConfigName(configInfo.Name)      // name of config file (without extension)
	viper.SetConfigType(configInfo.Extension) // REQUIRED if the config file does not have the extension in the name
	for _, path := range configInfo.Paths {
		viper.AddConfigPath(path)
	}

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var databaseConfig ConfigDatabase
	err = viper.Sub("database").Unmarshal(&databaseConfig)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}

	var serverConfig ConfigServer
	err = viper.Sub("server").Unmarshal(&serverConfig)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}

	return &ConfigApp{
		Database: databaseConfig,
		Server:   serverConfig,
	}
}
