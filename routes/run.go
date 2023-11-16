package routes

import (
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/zaggash/led-matrix-ui/handlers"
	"github.com/zaggash/led-matrix-ui/webui"
)

type HttpResponse struct {
	Message     string
	Status      int
	Description string
}

func ErrorHandler(c *gin.Context, err any) {
	goErr := errors.Wrap(err, 2)
	httpResponse := HttpResponse{Message: "Internal server error", Status: 500, Description: goErr.Error()}
	c.AbortWithStatusJSON(500, httpResponse)
}

var pixelFolder = "./PixelImages"

func Run() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()
	//gin.SetMode(gin.ReleaseMode) // Set to Production Mode !

	// Set Middleware
	router.Use(gin.CustomRecovery(ErrorHandler))

	// Set MaxFile size to 8Mb
	router.MaxMultipartMemory = 8 << 20

	// Set Webui
	//// Enable subdir from embedFS assets to /public
	subAssets, _ := fs.Sub(webui.EmbedAssets, "assets")
	router.StaticFS("/public/", http.FS(subAssets))
	//// Enable Pixel Images folder to /pixels
	router.Static(pixelFolder, pixelFolder)
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
		api.GET("/ping", handlers.Healthcheck)
		//api.POST("/upload", handlers.UploadFile)
		//api.POST("/display", handlers.DisplayImage)
		//api.POST("/settings", handlers.GetSettings)
		api.GET("/images", handlers.ListImages)
		api.GET("/images/:mime", handlers.ListImages)
	}
	// Start and run the server
	router.Run(":3000")
}
