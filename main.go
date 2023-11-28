package main

import (
	"log"

	"github.com/zaggash/led-matrix-ui/api"
	"github.com/zaggash/led-matrix-ui/utils"
)

func main() {
	utils.InitFlags()

	displayConfig, err := utils.NewDisplay()
	if err != nil {
		log.Panicln("Display error", err)
	}

	api.New(displayConfig).Run()

}
