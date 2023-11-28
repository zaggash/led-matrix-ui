package api

import (
	"image"
	"image/draw"
	"image/gif"
	"log"
	"net/http"
	"os"
	"time"

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

	if d.Channel != nil {
		close(d.Channel)
	}
	d.Toolkit.Canvas.Clear()

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
		gif, err := decodeGif(f)
		if err != nil {
			log.Panicln(err)
			return
		}
		d.Channel = drawGif(d, gif)
	default:
		log.Fatalln("Nothing...")
	}

	send := `{"showMessage": "` + imgName + ` displayed."}`
	http_code := http.StatusOK
	if d.Matrix.Render() != nil {
		send = "Error!"
		http_code = http.StatusInternalServerError
	}

	ctx.Header("HX-Trigger", send)
	ctx.Data(http_code, "text/html", nil)
}

func decodeImage(file *os.File) (image.Image, error) {
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func decodeGif(file *os.File) (*gif.GIF, error) {
	gif, err := gif.DecodeAll(file)
	if err != nil {
		return nil, err
	}
	return gif, nil
}

func drawImage(d *Display, img image.Image) error {
	w, h := d.Matrix.Geometry()
	img = imaging.Fill(img, w, h, imaging.Center, imaging.Lanczos)
	draw.Draw(d.Toolkit.Canvas, d.Toolkit.Canvas.Bounds(), img, image.Point{}, draw.Over)
	return nil
}

func drawGif(d *Display, gif *gif.GIF) chan bool {
	quit := make(chan bool)

	loop := 0
	delays := make([]time.Duration, len(gif.Delay))
	images := make([]image.Image, len(gif.Image))
	for i, image := range gif.Image {
		images[i] = image
		delays[i] = time.Millisecond * time.Duration(gif.Delay[i]) * 10
	}

	go func() {
		l := len(images)
		i := 0
		for {
			select {
			case <-quit:
				return
			default:
				start := time.Now()
				defer func() { time.Sleep(delays[i] - time.Since(start)) }()
				drawImage(d, images[i])
			}

			i++
			if i > l-1 {
				if loop == 0 {
					i = 0
					continue
				}
				break
			}
		}
	}()

	return quit
}
