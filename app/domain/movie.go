package domain

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title string `json:"title"`
}

type MovieInteractor interface {
	Test() string
}
