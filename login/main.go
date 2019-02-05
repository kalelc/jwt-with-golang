package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/csrf"
)

const Csrf = "800238847cd9105324275e055cfbac8b464fd036852599d3debe993a462b0013"
const Port = "3000"

func main() {

	csrfMiddleware := csrf.Protect([]byte(Csrf), csrf.Secure(false))
	router := NewRouter()
	fmt.Println("* Listen by " + Port + " port")
	http.ListenAndServe(":"+Port, csrfMiddleware(router))
}
