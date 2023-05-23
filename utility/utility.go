package utility

import (
	"exporting/auth"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func ValidateRequestJWT(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	tokenStr := cookie.Value
	claims := &auth.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return auth.MySignedKey, nil
		})
	if err != nil {
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	return true
}

// func createHash(key string) string {
// 	hasher := md5.New()
// 	hasher.Write([]byte(key))
// 	return hex.EncodeToString(hasher.Sum(nil))
// }

// func Encrypt(data []byte, passphrase string) []byte {
// 	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
// 	gcm, err := cipher.NewGCM(block)
// 	if err != nil {
// 		fmt.Println("Error!")
// 		// panic(err.Error())
// 	}
// 	nonce := make([]byte, gcm.NonceSize())
// 	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
// 		fmt.Println("Error!")
// 		// panic(err.Error())
// 	}
// 	ciphertext := gcm.Seal(nonce, nonce, data, nil)
// 	return ciphertext
// }

// func Decrypt(data []byte, passphrase string) []byte {
// 	key := []byte(createHash(passphrase))
// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		fmt.Println("Error!")
// 		// panic(err.Error())
// 	}
// 	gcm, err := cipher.NewGCM(block)
// 	if err != nil {
// 		fmt.Println("Error!")
// 		// panic(err.Error())
// 	}
// 	nonceSize := gcm.NonceSize()
// 	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
// 	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
// 	if err != nil {
// 		fmt.Println("Error!")
// 		// panic(err.Error())
// 	}
// 	return plaintext
// }
