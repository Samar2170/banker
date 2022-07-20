package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

var SECRET []byte = []byte("secret")

func parseToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(SECRET), nil
}

func GenerateToken(username string, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"email":    email,
	})
	tokenString, err := token.SignedString(SECRET)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		} else if r.URL.Path == "/signup" {
			next.ServeHTTP(w, r)
			return
		} else {
			authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer")
			if len(authHeader) != 2 {
				fmt.Println("Malformed token")
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Malformed token"))
			} else {
				jwToken := authHeader[1]
				token, err := jwt.Parse(jwToken, parseToken)
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					ctx := context.WithValue(r.Context(), "props", claims)

					next.ServeHTTP(w, r.WithContext(ctx))
				} else {
					fmt.Println(err)
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Invalid token"))
				}
			}

		}
	})
}
