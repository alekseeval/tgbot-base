package main

import (
	"fmt"
	"internal/configuration"
	"log"
)

func main() {

	var config *configuration.Config
	config, err := configuration.ReadYmlConfig(config)
	if err != nil {
		log.Panic(err)
	}
	config, err = configuration.ReadEnvConfig(config)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(config)
}
