package routes

import (
	dbConnection "github.com/armanimichael/link_shortener_go/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

type routerWithDb struct {
	*gin.Engine
	db *dbConnection.Database
}

func NewRouter(db *dbConnection.Database) *routerWithDb {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Homepage",
		})
	})

	api := router.Group("/api")
	{
		apiRoutes(api)
	}

	return &routerWithDb{
		router,
		db,
	}
}
