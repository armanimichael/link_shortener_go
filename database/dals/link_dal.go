package dals

import (
	"github.com/armanimichael/link_shortener_go/database"
	"github.com/armanimichael/link_shortener_go/database/models"
	"github.com/armanimichael/link_shortener_go/shortener"
)

const shortLinkLen int = 10

var db = database.GetDB()

func GetShortLink(originalLink string) (short string, exists bool) {
	var link models.Link
	condition := &models.Link{OriginalLink: originalLink}
	db.Where(condition).First(&link)

	if link.ID > 0 {
		return link.ShortLink, true
	}
	return "", false
}

func GetFullLink(shortLink string) (originalLink string, exists bool) {
	var link models.Link
	condition := &models.Link{ShortLink: shortLink}
	db.Where(condition).First(&link)

	if link.ID > 0 {
		return link.OriginalLink, true
	}
	return "", false
}

func SetShortLink(originalLink string) string {
	var short string
	var link models.Link

	exists := true
	for exists {
		exists, short = tryGenerateShort(link, short)
	}

	insertShortLink(originalLink, short)
	return short
}

func tryGenerateShort(link models.Link, short string) (bool, string) {
	short = string(shortener.GetRandomCombination(shortLinkLen))
	condition := &models.Link{ShortLink: short}
	db.Where(condition).First(&link)
	return link.ID > 0, short
}

func insertShortLink(originalLink string, shortLink string) {
	db.Create(&models.Link{
		OriginalLink: originalLink,
		ShortLink:    shortLink,
	})
}
