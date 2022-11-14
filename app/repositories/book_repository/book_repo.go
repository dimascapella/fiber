package bookrepository

import (
	"github.com/fiber/app/models"
	"gorm.io/gorm"
)

type bookConnection struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookConnection{
		db: db,
	}
}

func (repo *bookConnection) GetAll() ([]models.Book, error) {
	var books []models.Book
	data := repo.db.Preload("Author").Find(&books)
	if data.Error != nil {
		return []models.Book{}, data.Error
	}
	return books, nil
}

func (repo *bookConnection) GetById(BookId int) (models.Book, error) {
	var book models.Book
	data := repo.db.Preload("Author").Where("id=?", BookId).Find(&book)
	if data.Error != nil {
		return models.Book{}, data.Error
	}
	return book, nil
}

func (repo *bookConnection) GetAuthor(BookId int) (models.Author, error) {
	data, err := repo.GetById(BookId)
	if err != nil {
		return models.Author{}, err
	}
	return data.Author, nil
}

func (repo *bookConnection) CountBook(author models.Author) (int64, error) {
	var total int64
	data := repo.db.Model(&models.Book{}).Count(&total)
	if data.Error != nil {
		return 0, data.Error
	}
	return total, nil
}

func (repo *bookConnection) Create(authorId int, book models.Book) (models.Book, error) {
	data := repo.db.Create(&models.Book{
		Name:        book.Name,
		Genre:       book.Genre,
		Pages:       book.Pages,
		Description: book.Description,
		ReleaseDate: book.ReleaseDate,
		AuthorID:    authorId,
	})

	if data.Error != nil {
		return models.Book{}, data.Error
	}
	return book, nil
}

func (repo *bookConnection) Update(authorId int, book models.Book) (models.Book, error) {
	var newData models.Book
	data := repo.db.Where("id=?", book.ID).Where("author_id=?", authorId).Find(&newData)
	newData.Name = book.Name
	newData.Genre = book.Genre
	newData.Pages = book.Pages
	newData.Description = book.Description
	newData.ReleaseDate = book.ReleaseDate
	newData.AuthorID = authorId
	repo.db.Save(&newData)

	if data.Error != nil {
		return models.Book{}, data.Error
	}
	return book, nil
}

func (repo *bookConnection) Delete(book models.Book) error {
	data := repo.db.Delete(&book)
	if data.Error != nil {
		return data.Error
	}
	return nil
}
