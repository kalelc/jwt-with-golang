package main

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("jwt-token")

	fmt.Println("api controller")

	if tokenString != "" {
		fmt.Println(tokenString)

		token, err := Auth(tokenString)

		if err != nil {
			fmt.Println(err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			w.Header().Set("Content-Type", "application/json")
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			return
		}
	}
}
