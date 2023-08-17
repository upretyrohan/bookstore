package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/upretyrohan/bookstore/pkg/database"
	"net/http"
	"strconv"
	"io/ioutil"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := database.GetAllBooks()
	json.NewEncoder(w).Encode(books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookid := vars["id"]
	ID, _ := strconv.ParseInt(bookid, 0, 0)
	book := database.GetBookById(ID)
	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
  var book database.Book
  body, _ := ioutil.ReadAll(r.Body)
  json.Unmarshal(body, &book)
  database.CreateBook(&book)
  json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookid := vars["id"]
	ID, _ := strconv.ParseInt(bookid, 0, 0)
	book := database.DeleteBook(ID)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook database.Book
    body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &updateBook)
	vars := mux.Vars(r)
	bookid := vars["id"]
	ID, _ := strconv.ParseInt(bookid, 0, 0)
	bookdetail := database.GetBookById(ID)
	if updateBook.Name != ""{
		bookdetail.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookdetail.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookdetail.Publication = updateBook.Publication
	}

	database.GetDB().Save(&bookdetail)
	json.NewEncoder(w).Encode(bookdetail)
}