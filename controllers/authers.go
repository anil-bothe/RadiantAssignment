package controllers

import (
	"encoding/json"
	"exporting/db"
	"exporting/utility"
	"net/http"

	"github.com/gorilla/mux"
)

func AuthorList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	isValidJWT := utility.ValidateRequestJWT(w, r)
	if !isValidJWT {
		json.NewEncoder(w).Encode("JWT is not valid!")
		return
	}
	// now send authors list
	var author []db.Authors
	db.Database.Find(&author)

	json.NewEncoder(w).Encode(author)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	isValidJWT := utility.ValidateRequestJWT(w, r)
	if !isValidJWT {
		json.NewEncoder(w).Encode("JWT is not valid!")
		return
	}
	var author db.Authors
	json.NewDecoder(r.Body).Decode(&author)
	// Validate
	if author.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("author name is required")
		return
	}
	db.Database.Create(&author)
	json.NewEncoder(w).Encode(author)
}

func GetAuthorById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	isValidJWT := utility.ValidateRequestJWT(w, r)
	if !isValidJWT {
		json.NewEncoder(w).Encode("JWT is not valid!")
		return
	}
	var author db.Authors
	db.Database.First(&author, mux.Vars(r)["aid"])
	json.NewEncoder(w).Encode(author)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	isValidJWT := utility.ValidateRequestJWT(w, r)
	if !isValidJWT {
		json.NewEncoder(w).Encode("JWT is not valid!")
		return
	}
	var author db.Authors
	db.Database.First(&author, mux.Vars(r)["aid"])
	json.NewDecoder(r.Body).Decode(&author)

	// Validate
	if author.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("author name is required")
		return
	}
	db.Database.Save(&author)
	json.NewEncoder(w).Encode(author)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	isValidJWT := utility.ValidateRequestJWT(w, r)

	if !isValidJWT {
		json.NewEncoder(w).Encode("JWT is not valid!")
		return
	}

	var author db.Authors
	db.Database.Delete(&author, mux.Vars(r)["aid"])
	json.NewEncoder(w).Encode("Author is deleted successfully!")
}
