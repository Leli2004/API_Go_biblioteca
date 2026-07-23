package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

var cfg *Config

type Config struct {
	API  APIConfig  `mapstructure:"api"`
	DB   DBConfig   `mapstructure:"database"`
	Auth AuthConfig `mapstructure:"auth"`
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

type AuthConfig struct {
	JWTSecret      string `mapstructure:"jwt_secret"`
	JWTIssuer      string `mapstructure:"jwt_issuer"`
	JWTExpireHours int    `mapstructure:"jwt_expire_hours"`
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

	CfgJWTSecret      = "auth.jwt_secret"
	CfgJWTIssuer      = "auth.jwt_issuer"
	CfgJWTExpireHours = "auth.jwt_expire_hours"

	CfgJWTIssuerValue      = "biblioteca-api"
	CfgJWTExpireHoursValue = 24
)

func Load() error {
	viper.SetConfigName(ConfigName)
	viper.SetConfigType(ConfigType)
	viper.AddConfigPath(ConfigPath)

	viper.SetDefault(CfgApiPort, CfApiPortValue)
	viper.SetDefault(CfgDbHost, CfgDbHostValue)
	viper.SetDefault(CfgDbPort, CfgDbPortValue)

	viper.SetDefault(CfgJWTIssuer, CfgJWTIssuerValue)
	viper.SetDefault(CfgJWTExpireHours, CfgJWTExpireHoursValue)

	viper.SetEnvPrefix("")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

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

func GetAuth() AuthConfig {
	return cfg.Auth
}

func GetJWTSecret() string {
	return cfg.Auth.JWTSecret
}

func GetJWTIssuer() string {
	if cfg.Auth.JWTIssuer == "" {
		return CfgJWTIssuerValue
	}
	return cfg.Auth.JWTIssuer
}

func GetJWTExpiration() time.Duration {
	expireHours := cfg.Auth.JWTExpireHours
	if expireHours <= 0 {
		expireHours = CfgJWTExpireHoursValue
	}
	return time.Duration(expireHours) * time.Hour
}
