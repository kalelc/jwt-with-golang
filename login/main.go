package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/csrf"
)

func main() {

	csrfMiddleware := csrf.Protect([]byte("800238847cd9105324275e055cfbac8b464fd036852599d3debe993a462b0013"), csrf.Secure(false))
	router := NewRouter()
	fmt.Println("* Listen by 3000 port")
	http.ListenAndServe(":3000", csrfMiddleware(router))
}
