package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/csrf"
)

const Secret = "ee71a17fa8b84f99c3b99bdedb3b6e6910cfd59ed3fb0057ba96f72d2e952e0c4a26d0509702f2772ae20f8cc652124c88df5fb9a9b265ac12ee1af83e9ef0ba"

func Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		templates := template.Must(template.ParseFiles("views/login/index.html"))

		if err := templates.ExecuteTemplate(w, "index.html", map[string]interface{}{csrf.TemplateTag: csrf.TemplateField(r)}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodPost:
		name := r.FormValue("name")

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": name,
			"iss":      "login.app",
			"aud":      "api",
			"exp":      time.Now(),
		})

		tokenString, err := token.SignedString([]byte(Secret))

		if err != nil {
			fmt.Println(err)
		}

		http.Redirect(w, r, "http://localhost:3001?authentication="+tokenString, http.StatusSeeOther)
		return
	}
}
