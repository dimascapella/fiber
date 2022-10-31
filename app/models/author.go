package models

import "github.com/fiber/helpers"

type Author struct {
	helpers.BaseModel
	Name    string `gorm:"type:varchar(255); unique" json:"name" validate:"required"`
	Address string `gorm:"type:varchar(255)" json:"address" validate:"required"`
	Phone   string `gorm:"type:varchar(255)" json:"phone" validate:"required"`
	Books   []Book `json:"books"`
}
