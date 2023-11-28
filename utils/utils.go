package utils

import (
	"flag"
	"log"
	"os"
)

var (
	pixelsFolder = flag.String("images-folder", "./images/", "Logos folder path")
)

// Parse flags and error checks
func InitFlags() {
	flag.Parse()
	if _, err := os.Stat(*pixelsFolder); os.IsNotExist(err) {
		log.Fatalln(*pixelsFolder + " does not exists")
	}
}

func GetPixelsFolder() string {
	return *pixelsFolder
}
