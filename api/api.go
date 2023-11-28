package api

import (
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/zaggash/led-matrix-ui/utils"
	"github.com/zaggash/led-matrix-ui/webui"
)

type httpResponse struct {
	Message     string
	Status      int
	Description string
}

func (d *Display) Run() {
	pixelsFolder := utils.GetPixelsFolder()
	// Set the router as the default one shipped with Gin
	router := gin.Default()
	//gin.SetMode(gin.ReleaseMode) // TODO : Set to Production Mode !

	// Set Middleware
	router.Use(gin.CustomRecovery(errorHandler))

	// Set MaxFile size to 8Mb
	router.MaxMultipartMemory = 8 << 20

	// Set Webui
	var webStatic = "/public/"
	//// Enable subdir from embedFS assets to /public
	subAssets, _ := fs.Sub(webui.EmbedAssets, "assets")
	router.StaticFS(webStatic, http.FS(subAssets))

	//// Enable Pixel Images folder to /pixels
	router.Static(pixelsFolder, pixelsFolder)
	router.HTMLRender = webui.LoadTemplates(webui.EmbedTemplates)

	// Define favicon as favicon.ico and use it in html templates
	router.GET("favicon.ico", func(c *gin.Context) {
		file, _ := webui.EmbedAssets.ReadFile("assets/favicon.png")
		c.Data(
			http.StatusOK,
			"image/x-png",
			file,
		)
	})
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home", nil)
	})

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/ping", healthcheck)
		api.POST("/upload", uploadFile)
		api.POST("/draw/:type", d.drawImage)
		api.GET("/images/:type", listImages)
	}
	// Start and run the server
	router.Run(":3000")

}

func errorHandler(c *gin.Context, err any) {
	goErr := errors.Wrap(err, 2)
	httpResponse := httpResponse{Message: "Internal server error", Status: 500, Description: goErr.Error()}
	c.AbortWithStatusJSON(500, httpResponse)
}
