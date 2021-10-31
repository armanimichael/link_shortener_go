package dals

import (
	"fmt"
	"github.com/armanimichael/link_shortener_go/database/models"
	"github.com/armanimichael/link_shortener_go/shortener"
	"testing"
)

var testLinks = make(map[string]string)

func generateTestLinks(count int) {
	for i := 0; i < count; i++ {
		short := string(shortener.GetRandomCombination(shortLinkLen))
		originalLink := fmt.Sprintf("%d", i)
		db.Create(&models.Link{
			OriginalLink: originalLink,
			ShortLink:    short,
		})
		testLinks[originalLink] = short
	}
}

func deleteTestLinks() {
	for _, short := range testLinks {
		var dbEntry models.Link
		condition := &models.Link{ShortLink: short}
		db.First(&dbEntry, condition)
		db.Unscoped().Delete(&dbEntry)
	}
	testLinks = make(map[string]string)
}

func TestDal_GetShortLink_Existing(t *testing.T) {
	generateTestLinks(10)
	for _, short := range testLinks {
		var dbEntry models.Link
		condition := &models.Link{ShortLink: short}
		db.First(&dbEntry, condition)

		foundShort, exists := GetShortLink(dbEntry.OriginalLink)
		if !exists {
			t.Errorf("link not found in db")
		} else {
			t.Logf("Link found: %-2s -> %s", dbEntry.OriginalLink, foundShort)
		}
	}
	deleteTestLinks()
}

func TestDal_GetShortLink_NonExisting(t *testing.T) {
	var dbEntry models.Link
	condition := &models.Link{ShortLink: "fake-short-link"}
	db.First(&dbEntry, condition)

	foundShort, exists := GetShortLink(dbEntry.OriginalLink)
	if !exists && foundShort == "" {
		t.Logf("link not found in db")
	} else {
		t.Error()
	}
}

func TestDal_SetShortLink(t *testing.T) {
	generateTestLinks(1)
	originalLink := "0"
	_, exists := GetShortLink(originalLink)
	if !exists {
		SetShortLink(originalLink)
	}
	deleteTestLinks()
}
