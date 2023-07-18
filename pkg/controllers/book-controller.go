package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/code-you/Book-Management-system-api-mySQL/pkg/models"
	"github.com/code-you/Book-Management-system-api-mySQL/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("Error while parsing book: %v", err)
	}

	booksDetails, _ := models.GetBookById(Id)
	res, _ := json.Marshal(booksDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("Error while parsing book: %v", err)
	}

	booksDetails := models.DeleteBook(Id)
	res, _ := json.Marshal(booksDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)
	params := mux.Vars(r)
	bookId := params["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("Error while parsing book: %v", err)
	}
	booksDetails, db := models.GetBookById(Id)
	if updateBook.Name != "" {
		booksDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		booksDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		booksDetails.Publication = updateBook.Publication
	}
	db.Save(&booksDetails)
	res, _ := json.Marshal(booksDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}
