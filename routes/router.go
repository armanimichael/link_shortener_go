package routes

import (
	"github.com/armanimichael/link_shortener_go/database/dals"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Homepage",
		})
	})

	router.GET("/:short", func(context *gin.Context) {
		short := context.Param("short")
		originalLink, alreadyExists := dals.GetFullLink(short)
		if !alreadyExists {
			context.Redirect(http.StatusFound, "/")
		}
		context.Redirect(http.StatusMovedPermanently, originalLink)
	})

	api := router.Group("/api")
	{
		apiRoutes(api)
	}

	return router
}
