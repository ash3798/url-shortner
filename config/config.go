package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	ApplicationPort int `split_words:"true" default:"9999"`
}

var (
	Manager config
)

func InitEnv() bool {
	err := envconfig.Process("", &Manager)
	if err != nil {
		log.Println("Error while initializing the environment , Error : ", err.Error())
		return false
	}
	return true
}
