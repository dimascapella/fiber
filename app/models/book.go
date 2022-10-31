package models

import (
	"time"

	"github.com/fiber/helpers"
)

type Book struct {
	helpers.BaseModel
	Name        string    `gorm:"type:varchar(255); unique" json:"name" validate:"required"`
	Genre       string    `gorm:"type:varchar(255)" json:"genre" validate:"required"`
	Pages       int       `json:"pages" validate:"required"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	ReleaseDate time.Time `json:"releaseDate" validate:"required"`
	AuthorID    int
	Author      Author
}
