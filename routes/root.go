package routes

import (
	"github.com/armanimichael/link_shortener_go/database/dals"
	"github.com/gin-gonic/gin"
	"net/http"
)

func rootRoutes(g *gin.RouterGroup) {
	g.GET("/", homepage)
	g.GET("/:short", shortLinkRedirect)
}

func homepage(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Homepage",
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
