package authorrepository

import (
	"github.com/fiber/app/models"
	"gorm.io/gorm"
)

type authorConnection struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &authorConnection{
		db: db,
	}
}

func (repo *authorConnection) GetAll() ([]models.Author, error) {

}

func (repo *authorConnection) GetById(AuthorId int) (models.Author, error) {

}

func (repo *authorConnection) Create(Author models.Author) (models.Author, error) {

}

func (repo *authorConnection) Update(Author models.Author) (models.Author, error) {

}

func (repo *authorConnection) Delete(Author models.Author) error {

}
