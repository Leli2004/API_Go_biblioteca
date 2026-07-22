package config

import (
	"github.com/spf13/viper"
)

var cfg *Config

type Config struct {
	API APIConfig `mapstructure:"api"`
	DB  DBConfig  `mapstructure:"database"`
}

type APIConfig struct {
	Port string `mapstructure:"port"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Pass     string `mapstructure:"pass"`
	Database string `mapstructure:"name"`
}

const (
	ConfigName = "config"
	ConfigType = "yaml"
	ConfigPath = "./config"

	CfgApiPort     = "api.port"
	CfApiPortValue = "8001"

	CfgDbHost      = "database.host"
	CfgDbHostValue = "localhost"

	CfgDbPort      = "database.port"
	CfgDbPortValue = "5433"

	CfgDbUser     = "database.user"
	CfgDbPass     = "database.pass"
	CfgDbDatabase = "database.name"
)

func Load() error {
	viper.SetConfigName(ConfigName)
	viper.SetConfigType(ConfigType)
	viper.AddConfigPath(ConfigPath)

	viper.SetDefault(CfgApiPort, CfApiPortValue)
	viper.SetDefault(CfgDbHost, CfgDbHostValue)
	viper.SetDefault(CfgDbPort, CfgDbPortValue)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	cfg = &Config{}

	if err := viper.Unmarshal(cfg); err != nil {
		return err
	}

	return nil
}

func Get() *Config {
	return cfg
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
