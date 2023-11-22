package api

import (
	"image"
	"image/draw"
	"log"
	"os"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/zaggash/led-matrix-ui/display"
)

type Display struct {
	*display.Config
}

func New(c *display.Config) *Display {
	return &Display{c}
}

func drawImage(d *Display, p string) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {

		imgPath := ctx.PostForm("Path")

		f, err := os.Open(imgPath)
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()

		img, _, err := image.Decode(f)
		if err != nil {
			log.Fatalln(err)
		}

		w, h := d.Matrix.Geometry()
		d.Toolkit.Canvas.Clear() // Clear current display
		img = imaging.Fill(img, w, h, imaging.Center, imaging.Lanczos)
		draw.Draw(d.Toolkit.Canvas, d.Toolkit.Canvas.Bounds(), img, image.ZP, draw.Over)
		d.Matrix.Render()

		// DEBUG: Print POST body request
		// body, _ := io.ReadAll(ctx.Request.Body)
		// println(string(body))
	})
}
