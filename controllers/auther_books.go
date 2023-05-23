package controllers

import (
	"encoding/json"
	"exporting/db"
	"net/http"
)

func AutherBooksList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var authorBooks []db.AutherBooks
	db.Database.Find(&authorBooks)
	json.NewEncoder(w).Encode(authorBooks)
}
