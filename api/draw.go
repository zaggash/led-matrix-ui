package api

import (
	"image"
	"image/draw"
	"log"
	"net/http"
	"os"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/zaggash/led-matrix-ui/utils"
)

type Display struct {
	*utils.Display
}

func New(c *utils.Display) *Display {
	return &Display{c}
}

func (d *Display) drawImage(ctx *gin.Context) {
	imgPath := ctx.PostForm("path")
	imgName := ctx.PostForm("name")
	paramType := ctx.Param("type")

	// Clear display
	if d.Channel != nil {
		d.Channel <- true
	}
	d.Toolkit.Canvas.Clear()

	// Open requested image
	f, err := os.Open(imgPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	switch paramType {
	case "static":
		img, err := decodeImage(f)
		if err != nil {
			log.Panicln(err)
			return
		}
		err = drawImage(d, img)
		if err != nil {
			log.Panicln(err)
			return
		}
	case "animated":
		w, h := d.Matrix.Geometry()
		d.Toolkit.Transform = func(img image.Image) *image.NRGBA {
			//return imaging.Fit(img, w, h, imaging.Lanczos)
			return imaging.Fill(img, w, h, imaging.Center, imaging.Lanczos)
		}
		d.Channel, err = d.Toolkit.PlayGIF(f)
		if err != nil {
			log.Panicln(err)
			return
		}
	}

	// Send Frontend result
	send := `{"showMessage": "` + imgName + ` displayed."}`
	http_code := http.StatusOK
	if d.Matrix.Render() != nil {
		send = "Error!"
		http_code = http.StatusInternalServerError
	}

	ctx.Header("HX-Trigger", send)
	ctx.Data(http_code, "text/html", nil)
}

// Decode Static image
func decodeImage(file *os.File) (image.Image, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// Draw static image
func drawImage(d *Display, img image.Image) error {
	w, h := d.Matrix.Geometry()
	img = imaging.Fill(img, w, h, imaging.Center, imaging.Lanczos)
	draw.Draw(d.Toolkit.Canvas, d.Toolkit.Canvas.Bounds(), img, image.Point{}, draw.Over)
	return nil
}
