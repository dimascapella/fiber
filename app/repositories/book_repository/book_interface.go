package bookrepository

import "github.com/fiber/app/models"

type BookRepository interface {
	GetAll() ([]models.Book, error)
	GetById(BookId int) (models.Book, error)
	CountBook(author models.Author) (int, error)
	Create(author models.Author, book models.Book) (models.Book, error)
	Update(author models.Author, book models.Book) (models.Book, error)
	Delete(book models.Book) error
}
