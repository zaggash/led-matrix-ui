package main

import (
	"flag"
	"log"

	api "github.com/zaggash/led-matrix-ui/api"
	"github.com/zaggash/led-matrix-ui/display"
)

var pixelFolder = flag.String("image-folder", "./PixelImages/", "Pixel logos folder path")

func main() {

	displayConfig, err := display.New()
	if err != nil {
		log.Panicln("Display error", err)
	}

	api.New(displayConfig).Run(*pixelFolder)
}

func init() {
	flag.Parse()
}
