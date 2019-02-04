package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
)

func Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		templates := template.Must(template.ParseFiles("views/login/index.html"))

		if err := templates.ExecuteTemplate(w, "index.html", map[string]interface{}{csrf.TemplateTag: csrf.TemplateField(r)}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodPost:
		fmt.Println("post request")

	}
}
