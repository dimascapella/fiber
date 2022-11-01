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
	var authors []models.Author
	return authors, nil
}

func (repo *authorConnection) GetById(authorId int) (models.Author, error) {
	var author models.Author
	return author, nil
}

func (repo *authorConnection) Create(author models.Author) (models.Author, error) {
	return author, nil
}

func (repo *authorConnection) Update(author models.Author) (models.Author, error) {
	return author, nil
}

func (repo *authorConnection) Delete(author models.Author) error {
	return nil
}
