package routes

import (
	"github.com/armanimichael/link_shortener_go/database"
	"github.com/armanimichael/link_shortener_go/database/dals"
	"github.com/armanimichael/link_shortener_go/database/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var db = database.GetDB()

func rootRoutes(g *gin.RouterGroup) {
	g.GET("/", homepage)
	g.GET("/:short", shortLinkRedirect)
}

func homepage(context *gin.Context) {
	var links []models.Link
	db.Find(&links)
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Homepage",
		"links": links,
	})
}

func shortLinkRedirect(context *gin.Context) {
	short := context.Param("short")
	originalLink, alreadyExists := dals.GetFullLink(short)
	if !alreadyExists {
		context.Redirect(http.StatusFound, "/")
	}
	context.Redirect(http.StatusMovedPermanently, originalLink)
}
