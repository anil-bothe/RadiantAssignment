package controllers

import (
	"encoding/json"
	"exporting/auth"
	"exporting/db"
	"exporting/utility"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// TokenGeneration
	var credentials auth.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if credentials.Email == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var user db.Users
	db.Database.First(&user, "Email = ?", credentials.Email) //compare uniq field
	expectedPassword := user.Password
	err = utility.CheckPassword(credentials.Password, expectedPassword) // compare password
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Hour * 1) // 1hrs

	claims := &auth.Claims{
		Email: credentials.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(auth.MySignedKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
	var res = map[string]string{
		"Expires": expirationTime.String(),
		"Token":   tokenString,
		"user":    user.Email,
	}
	json.NewEncoder(w).Encode(res)
}

func UserRegistration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user db.Users
	errs := []string{}
	json.NewDecoder(r.Body).Decode(&user)

	// Validate
	if user.Name == "" {
		errs = append(errs, fmt.Errorf("name is required").Error())
	}
	if user.Email == "" {
		errs = append(errs, fmt.Errorf("email is required").Error())
	}
	if user.Password == "" {
		errs = append(errs, fmt.Errorf("password is required").Error())
	}
	if len(errs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs)
		return
	}
	hash, err := utility.HashPassword(user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user.Password = hash
	db.Database.Create(&user)
	json.NewEncoder(w).Encode(user)
}
