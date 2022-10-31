package authorrepository

import "github.com/fiber/app/models"

type AuthorRepository interface {
	GetAll() ([]models.Author, error)
	GetById(AuthorId int) (models.Author, error)
	Create(Author models.Author) (models.Author, error)
	Update(Author models.Author) (models.Author, error)
	Delete(Author models.Author) error
}
