package models

import "gorm.io/gorm"

type Shorten struct {
	gorm.Model
	Hash uint32
	LongURL string
}