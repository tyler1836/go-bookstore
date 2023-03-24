package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/mediocrerlplayer/go-bookstore/pkg/models"
	"github.com/mediocrerlplayer/go-bookstore/pkg/utils"
)

var NewBook bookmodels.Book
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Bookstore!")
}
func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := bookmodels.GetBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing bookId")
	}
	bookDetails := bookmodels.GetBook(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &bookmodels.Book{}
	utils.ParseBody(r, CreateBook)
	b:= CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing bookId")
	}
	bookDetails := bookmodels.DeleteBook(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}