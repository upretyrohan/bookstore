package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
)

var db * gorm.DB
var  err error

type Book struct{
	gorm.Model
	Name string 
	Author string 
	Publication string 
}

func Init() {

	dsn := "root:password@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Book{})

	fmt.Println("database init successfull")
}

func CreateBook(book* Book)  {
	db.Create(&book)
}

func GetAllBooks() []Book {
	var Books[] Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) *Book {
	var getBook Book
	db.Where("ID=?", Id).Find(&getBook)
	return &getBook
}
func DeleteBook(ID int64) Book{
	var book Book
	db.Where("ID=?", ID).Delete(&book)
	return book
}

func GetDB() *gorm.DB{
	return db
}