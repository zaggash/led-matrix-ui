package handlers

import (
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
)

type image struct {
	Name     string
	Path     string
	Size     int64 // Size in bytes
	MimeType string
}

func ListImages(ctx *gin.Context) {
	rootFolder := "PixelImages/"
	whitelistedMime := []string{""}
	requestedMime := strings.ToLower(ctx.Param("mime"))
	var images []image

	err := filepath.WalkDir(rootFolder, func(filename string, file fs.DirEntry, err error) error {
		if !file.IsDir() {
			switch requestedMime {
			case "":
				whitelistedMime = []string{"image/png", "image/gif", "image/jpeg"}
			case "png":
				whitelistedMime = []string{"image/png"}
			case "gif":
				whitelistedMime = []string{"image/gif"}
			}

			fileMime, err := mimetype.DetectFile(filename)
			if err != nil {
				log.Println(err)
			}

			if mimetype.EqualsAny(fileMime.String(), whitelistedMime...) {
				info, err := os.Stat(filename)
				if err != nil {
					log.Println(err)
				}
				images = append(images, image{
					Name:     info.Name(),
					Path:     filename,
					Size:     info.Size(),
					MimeType: fileMime.String(),
				})
			}
		}
		return nil
	})

	if err != nil {
		log.Println(err)
	}
	ctx.JSON(http.StatusOK, images)
}
