package models

import "gorm.io/gorm"

type Shorten struct {
	gorm.Model
	Short string
	LongURL string
}