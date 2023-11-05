package repository

import "github.com/kevinjuliow/golang-libraryRestApi/models"

type BooksRepository interface {
	Save(books models.Book)
	Update(books models.Book)
	Delete(id uint)
	FindById(books models.Book, id uint)
	FindAll()
}
