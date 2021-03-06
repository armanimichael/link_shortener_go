package dals

import (
	"github.com/armanimichael/link_shortener_go/database"
	"github.com/armanimichael/link_shortener_go/database/models"
	"github.com/armanimichael/link_shortener_go/shortener"
)

const shortLinkLen int = 10

var db = database.GetDB()

// GetShortLink fetches the short version of a given link and returns it.
// If not found it returns an empty string.
func GetShortLink(originalLink string) (short string, exists bool) {
	var link models.Link
	condition := &models.Link{OriginalLink: originalLink}
	result := db.Find(&link, condition)
	if result.Error != nil || result.RowsAffected < 1 {
		return "", false
	}

	return link.ShortLink, true
}

// GetFullLink fetches the original version of a given short-link and returns it.
// If not found it returns an empty string.
func GetFullLink(shortLink string) (originalLink string, exists bool) {
	var link models.Link
	condition := &models.Link{ShortLink: shortLink}
	db.Where(condition).First(&link)

	if link.OriginalLink != "" {
		return link.OriginalLink, true
	}
	return "", false
}

// SetShortLink generates and set a new unique short
// link for a given original link
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
	return link.OriginalLink != "", short
}

func insertShortLink(originalLink string, shortLink string) {
	db.Create(&models.Link{
		OriginalLink: originalLink,
		ShortLink:    shortLink,
	})
}
