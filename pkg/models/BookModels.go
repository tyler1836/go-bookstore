package bookmodels

import (
	"github.com/jinzhu/gorm"
	"github.com/mediocrerlplayer/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func GetBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBook(id int64) *Book {
	var book Book
	db.Where("ID=?", id).Find(&book)
	return &book
}

//receive type book and then return type book
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}