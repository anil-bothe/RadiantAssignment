package controllers

import (
	"encoding/json"
	"exporting/db"
	"net/http"

	"github.com/gorilla/mux"
)

func BooksList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []db.Books
	db.Database.Find(&books)
	json.NewEncoder(w).Encode(books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book db.Books
	json.NewDecoder(r.Body).Decode(&book)
	// Validate
	if book.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("book name is required")
		return
	}
	// add book & auther
	db.Database.Create(&book)
	json.NewEncoder(w).Encode(book)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book db.Books
	db.Database.First(&book, mux.Vars(r)["bid"])
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book db.Books
	db.Database.First(&book, mux.Vars(r)["bid"])
	json.NewDecoder(r.Body).Decode(&book)
	// Validate
	if book.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("book name is required")
		return
	}
	db.Database.Save(&book)
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book db.Books
	db.Database.Delete(&book, mux.Vars(r)["bid"])
	json.NewEncoder(w).Encode("book is deleted successfully!")
}
