package main

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Cookies())
	tokenString := r.Header.Get("jwt-token")
	if tokenString != "" {
		fmt.Println(tokenString)

		token, err := Sign(tokenString)

		if err != nil {
			fmt.Println(err)
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			cookie := http.Cookie{
				Name:  "JWT_TOKEN",
				Value: tokenString,
			}
			http.SetCookie(w, &cookie)
			w.WriteHeader(http.StatusOK)
		} else {
			fmt.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
}
