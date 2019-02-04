package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := NewRouter()
	http.ListenAndServe(":3000", router)
	fmt.Println("* Listen by 3000 port")
}
