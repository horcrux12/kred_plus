package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

var Attr = applicationConfiguration{}

func GenerateConfiguration() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	err = envconfig.Process("", &Attr)
	if err != nil {
		log.Fatalln(err)
	}
}
