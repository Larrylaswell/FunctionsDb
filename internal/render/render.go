package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func RendenderTemplate(w http.ResponseWriter, htmlPage string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + htmlPage)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}
