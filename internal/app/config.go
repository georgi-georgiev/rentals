package app

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfiguration
	Swagger  SwaggerConfiguration
	App      AppConfiguration
	Postgres PostgresConfiguration
}

type ServerConfiguration struct {
	RootURL         string
	Host            string
	Port            string
	ShutdownTimeout int
	Timeout         struct {
		Server time.Duration
		Read   time.Duration
		Write  time.Duration
		Idle   time.Duration
	}
}

type PostgresConfiguration struct {
	Driver       string
	Dbname       string
	Username     string
	Password     string
	Host         string
	Port         string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

type SwaggerConfiguration struct {
	Host string
}

type AppConfiguration struct {
	Name    string
	Version string
}

func NewConfig() *Config {
	var config *Config

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	path := os.Getenv("CONFIG_PATH")

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return config
}
