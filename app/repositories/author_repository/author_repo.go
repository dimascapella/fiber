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
	data := repo.db.Find(&authors)
	if data.Error != nil {
		return []models.Author{}, data.Error
	}
	return authors, nil
}

func (repo *authorConnection) GetById(authorId int) (models.Author, error) {
	var author models.Author
	data := repo.db.Where("id=?", authorId).Find(&author)
	if data.Error != nil {
		return models.Author{}, data.Error
	}
	return author, nil
}

func (repo *authorConnection) Create(author models.Author) (models.Author, error) {
	data := repo.db.Create(&models.Author{
		Name:    author.Name,
		Phone:   author.Phone,
		Address: author.Address,
	})
	if data.Error != nil {
		return models.Author{}, data.Error
	}
	return author, nil
}

func (repo *authorConnection) Update(author models.Author) (models.Author, error) {
	data := repo.db.Model(&models.Author{}).Where("id=?", author.ID).Updates(author)
	if data.Error != nil {
		return models.Author{}, data.Error
	}
	return author, nil
}

func (repo *authorConnection) Delete(author models.Author) error {
	data := repo.db.Delete(&author)
	if data.Error != nil {
		return data.Error
	}
	return nil
}

func (repo *authorConnection) Count() (int64, error) {
	var total int64
	data := repo.db.Model(&models.Author{}).Count(&total)
	if data.Error != nil {
		return 0, data.Error
	}
	return total, nil
}
