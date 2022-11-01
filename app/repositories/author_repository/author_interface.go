package authorrepository

import "github.com/fiber/app/models"

type AuthorRepository interface {
	GetAll() ([]models.Author, error)
	GetById(authorId int) (models.Author, error)
	Create(author models.Author) (models.Author, error)
	Update(author models.Author) (models.Author, error)
	Delete(author models.Author) error
}
