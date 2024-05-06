package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type ConfigDatabase struct {
	Port     string `yml:"port" env:"PORT" env-default:"5432"`
	Host     string `yml:"host" env:"HOST" env-default:"localhost"`
	Name     string `yml:"name" env:"NAME" env-default:"postgres"`
	User     string `yml:"user" env:"USER" env-default:"user"`
	Password string `yml:"password" env:"PASSWORD"`
	SSLMode  string `yml:"sslmode" env:"SSLMODE"`
}

func (c *ConfigDatabase) GetDataSourceName() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		c.Host, c.Port, c.Name, c.User, c.Password, c.SSLMode)
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

func DefaultConfigInfo() *ConfigInfo {
	return &ConfigInfo{
		Name:      "config",
		Extension: "yaml",
		Paths:     []string{".", "config", "pkg/config"},
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
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	if err = godotenv.Load(); err != nil {
		// log.Fatal()
		panic(err)
	}
	databaseConfig.Password = os.Getenv("POSTGRES_DB_PASSWORD")

	var serverConfig ConfigServer
	err = viper.Sub("server").Unmarshal(&serverConfig)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	return &ConfigApp{
		Database: databaseConfig,
		Server:   serverConfig,
	}
}
