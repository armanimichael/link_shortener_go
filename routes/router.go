package routes

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	root := router.Group("/")
	rootRoutes(root)

	api := router.Group("/api")
	apiRoutes(api)

	return router
}
