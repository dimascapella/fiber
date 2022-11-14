package bookrepository

import "github.com/fiber/app/models"

type BookRepository interface {
	GetAll() ([]models.Book, error)
	GetById(BookId int) (models.Book, error)
	GetAuthor(BookId int) (models.Author, error)
	CountBook(author models.Author) (int64, error)
	Create(authorId int, book models.Book) (models.Book, error)
	Update(authorId int, book models.Book) (models.Book, error)
	Delete(book models.Book) error
}
