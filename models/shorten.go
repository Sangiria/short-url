package models

import "gorm.io/gorm"

type Shorten struct {
	gorm.Model
	Token string
	LongURL string
}