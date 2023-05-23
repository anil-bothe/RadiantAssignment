package routes

import (
	"exporting/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func MainRouting() {
	r := mux.NewRouter()

	r.HandleFunc("/register", controllers.UserRegistration).Methods("POST")
	r.HandleFunc("/login", controllers.UserLogin).Methods("POST")
	// r.HandleFunc("/refresh", controllers.RefreshToken).Methods("POST")

	r.HandleFunc("/authors", controllers.AuthorList).Methods("GET")
	r.HandleFunc("/authors/create", controllers.CreateAuthor).Methods("POST")
	r.HandleFunc("/authors/{aid}", controllers.GetAuthorById).Methods("POST")
	r.HandleFunc("/authors/{aid}/edit", controllers.UpdateAuthor).Methods("PUT")
	r.HandleFunc("/authors/{aid}/delete", controllers.DeleteAuthor).Methods("DELETE")

	r.HandleFunc("/books", controllers.BooksList).Methods("GET")
	r.HandleFunc("/books/create", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{bid}", controllers.GetBookById).Methods("POST")
	r.HandleFunc("/books/{bid}/edit", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{bid}/delete", controllers.DeleteBook).Methods("POST")

	// temp
	r.HandleFunc("/authors/books", controllers.AutherBooksList).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
