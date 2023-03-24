package routes

import (
	"github.com/gorilla/mux"
	"github.com/mediocrerlplayer/go-bookstore/pkg/controllers"
)
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	//router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
}