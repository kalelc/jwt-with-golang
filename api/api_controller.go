package main

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

const JwtTokenName = "jwt_token"
const LoginHost = "http://localhost:3000"

func Index(w http.ResponseWriter, r *http.Request) {
	tokenString, err := getJwt(r)

	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, LoginHost, http.StatusSeeOther)
		return
	}

	if tokenString != "" {
		fmt.Println(tokenString)

		token, err := Sign(tokenString)

		if err != nil {
			fmt.Println(err)
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			cookie := http.Cookie{
				Name:  JwtTokenName,
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

func getJwt(r *http.Request) (string, error) {
	tokenString := r.URL.Query().Get("authentication")

	if tokenString == "" {
		cookie, err := r.Cookie(JwtTokenName)

		if err != nil {
			return "", err
		}

		tokenString = cookie.Value
	}

	return tokenString, nil

}
