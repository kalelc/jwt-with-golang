package main

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

const Secret = "ee71a17fa8b84f99c3b99bdedb3b6e6910cfd59ed3fb0057ba96f72d2e952e0c4a26d0509702f2772ae20f8cc652124c88df5fb9a9b265ac12ee1af83e9ef0ba"

func Sign(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(Secret), nil
	})

	return token, err

}
