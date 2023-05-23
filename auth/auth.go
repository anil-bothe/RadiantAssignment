package auth

import (
	"github.com/dgrijalva/jwt-go"
)

// var MySignedKey = []byte(os.Getenv("JWT_SECRETE"))
var MySignedKey = []byte("supersecrete")

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
