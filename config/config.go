package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"
)

var Attr = applicationConfiguration{}

func GenerateConfiguration() {
	fmt.Println("Env : ", loadEnv(os.Args))

	err := envconfig.Process("", &Attr)
	if err != nil {
		log.Fatalln(err)
	}
}

func loadEnv(args []string) (env string) {
	if len(os.Args) == 1 {
		godotenv.Load()
		env = "local"
	} else if len(os.Args) == 2 {
		env = os.Args[1]
		switch os.Args[1] {
		case "prod":
			godotenv.Load(".env.prod")
		default:
			err := errors.New("Specific Environment Not Found")
			log.Fatal(err)
		}
	} else {
		err := errors.New("Undefine Command to Run API with Specific Environment")
		log.Fatal(err)
	}

	return
}
