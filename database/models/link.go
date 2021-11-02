package models

type Link struct {
	OriginalLink string `gorm:"uniqueIndex"`
	ShortLink    string `gorm:"uniqueIndex"`
}
