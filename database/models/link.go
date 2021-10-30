package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	ID           uint
	OriginalLink string
	ShortLink    string
}
