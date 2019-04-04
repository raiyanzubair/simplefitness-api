package auth

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"time"
)

func getJWTKey() string {
	signingKey := os.Getenv("JWT_KEY")
	if signingKey == "" {
		signingKey = "hellothere"
	}
	return signingKey
}

// GenerateJWT generates a jwt for us to use
func GenerateJWT(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  "simplfitness",
		"sub":  "user",
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
		"name": "Raiyan",
	})

	tokenString, err := token.SignedString([]byte(getJWTKey()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWTMiddleware verifies JWTs in requests for certain routes
func ValidateJWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] == nil {
			http.Error(w, "Authorization error", http.StatusUnauthorized)
			return
		}
		token, err := jwt.Parse(r.Header["Token"][0],
			func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					return nil, fmt.Errorf("Error Validating JWT")
				}
				return getJWTKey(), nil
			})

		if err != nil {
			http.Error(w, "Authorization error", http.StatusUnauthorized)
			return
		}

		if token.Valid {
			next.ServeHTTP(w, r)
		}
	})
}
