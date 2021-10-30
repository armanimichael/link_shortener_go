package routes

import (
	"github.com/armanimichael/link_shortener_go/database/dals"
	"github.com/gin-gonic/gin"
	"net/http"
)

func apiRoutes(g *gin.RouterGroup) {
	g.POST("/generate-link", generateLink)
}

func generateLink(context *gin.Context) {
	url := context.PostForm("link")

	short, alreadyExists := dals.GetShortLink(url)
	if !alreadyExists {
		short = dals.SetShortLink(url)
	}

	context.JSON(http.StatusOK, gin.H{
		"short":    short,
		"original": url,
	})
}
