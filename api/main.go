package main

import (
	"fmt"
	"net/http"
)

const Port = "3001"

func main() {
	router := NewRouter()
	fmt.Println("* Listen by " + Port + " port")
	http.ListenAndServe(":"+Port, router)
}
