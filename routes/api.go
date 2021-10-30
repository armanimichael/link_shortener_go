package routes

import (
	"github.com/armanimichael/link_shortener_go/database/dals"
	httpUtils "github.com/armanimichael/link_shortener_go/http_utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func apiRoutes(g *gin.RouterGroup) {
	g.POST("/generate-link", generateLink)
}

func generateLink(context *gin.Context) {
	url := context.PostForm("link")

	if !httpUtils.IsValidUrl(url) {
		invalidLinkResponse(context)
	} else {
		short, alreadyExists := dals.GetShortLink(url)
		if !alreadyExists {
			short = dals.SetShortLink(url)
		}
		validLinkResponse(short, url, context)
	}
}

func validLinkResponse(short string, originalLink string, context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"short":    short,
		"original": originalLink,
	})
}

func invalidLinkResponse(context *gin.Context) {
	context.JSON(http.StatusUnprocessableEntity, gin.H{
		"message": "invalid url",
	})
}
