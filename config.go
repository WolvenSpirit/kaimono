package main

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type EnvConfig struct {
	GRPCPort      string `env:"GRPC_PORT"`
	GRPCAddr      string `env:"GRPC_ADDR"`
	ServerAddress string `env:"SHOP_BIND_ADDR"`
	DBDriver      string `env:"API_DB_DRIVER"`
	DBPort        string `env:"API_DB_PORT"`
	DBHost        string `env:"API_DB_HOST"`
	DBName        string `env:"API_DB_NAME"`
	DBUser        string `env:"API_DB_USER"`
	DBPassword    string `env:"API_DB_PASS"`
	DBPqSslMode   string `env:"API_PQ_SSLMODE"`
	AuthTable     string `env:"AUTH_TABLE"`
}

type YmlConfig struct {
	GRPCPort string `yaml:"GRPC_PORT" env-default:"80"`
	GRPCAddr string `yaml:"GRPC_ADDR" env-default:"0.0.0.0"`

	ServerAddress string `yaml:"SHOP_BIND_ADDR" env-default:":8080"`

	DBDriver    string `yaml:"API_DB_DRIVER" env-default:"postgres"`
	DBPort      string `yaml:"API_DB_PORT" env-default:"5432"`
	DBHost      string `yaml:"API_DB_HOST" env-default:"localhost"`
	DBName      string `yaml:"API_DB_NAME" env-default:"postgres"`
	DBUser      string `yaml:"API_DB_USER" env-default:"postgres"`
	DBPassword  string `yaml:"API_DB_PASS" env-default:""`
	DBPqSslMode string `yaml:"API_PQ_SSLMODE" env-default:"disable"`
	AuthTable   string `env:"AUTH_TABLE" env-default:"users"`
}

var env EnvConfig
var yml YmlConfig
var config interface{}
var envParseErr error
var ymlParseErr error

func parseConfig() {
	if envParseErr := cleanenv.ReadEnv(&env); envParseErr != nil {
		log.Println(envParseErr.Error())
	}
	if ymlParseErr := cleanenv.ReadConfig("config.yml", &yml); ymlParseErr != nil {
		log.Println(ymlParseErr.Error())
	}
}
