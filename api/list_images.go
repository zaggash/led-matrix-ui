package api

import (
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/zaggash/led-matrix-ui/utils"
)

type localImage struct {
	Name     string
	Path     string
	Size     int64 // Size in bytes
	MimeType string
}

func listImages(ctx *gin.Context) {
	var err error
	var images []localImage
	pixelsFolder := utils.GetPixelsFolder()
	requestedType := ctx.Param("type")

	err = filepath.WalkDir(pixelsFolder, func(filename string, file fs.DirEntry, err error) error {
		if !file.IsDir() {
			fileMime, err := mimetype.DetectFile(filename)
			if err != nil {
				return err
			}

			info, err := file.Info()
			if err != nil {
				return err
			}

			if shouldIncludeImage(requestedType, fileMime.String()) {
				images = append(images, localImage{
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

	//ctx.JSON(http.StatusOK, images)
	ctx.Data(http.StatusOK, "text/html", createGallery(requestedType, images))
}

func createGallery(requestedType string, images []localImage) []byte {
	html := ""
	for _, img := range images {
		card :=
			`<div class="col col-3 mb-3">
             <div class="card h-100 text-center">
             <div class="card-header text-nowrap overflow-auto">` + img.Name + `</div>
             <div class="ratio ratio-1x1">
             <img src="` + img.Path + `" class="card-img-top h-100 border-bottom border-1 rounded-0" alt="` + img.Name + `" />
              </div>
              <div class="card-body d-flex justify-content-center">
                <button class="btn btn-primary" hx-post="/api/draw/` + requestedType + `" hx-trigger="click" hx-swap="none" hx-vals='{"name":"` + img.Name + `", "path":"` + img.Path + `"}' >
                  Apply
                </button>
              </div>
            </div>
          </div>`
		html += card
	}
	return []byte(html)
}

// Compare images on disk to the API mimetype parameter
func shouldIncludeImage(param, mimeType string) bool {
	switch param {
	case "animated":
		return contains(mimeType, []string{"image/gif"})
	case "static":
		return (contains(mimeType, []string{"image/"}) && !contains(mimeType, []string{"image/gif"}))
	default:
		return false
	}
}

// Check if s contains at least one string value from arr
func contains(s string, arr []string) bool {
	for _, a := range arr {
		if strings.Contains(s, a) {
			return true
		}
	}
	return false
}
