package main

import (
	"html/template"
	"log"
	"net/http"
)

type Gopher struct {
	Name string
}

// handler router
func helloGopherHandler(w http.ResponseWriter, r *http.Request) {
	var gopherName string
	gopherName = r.URL.Query().Get("gophername")
	if gopherName == "" {
		gopherName = "Gopher"
	}

	gopher := Gopher{Name: gopherName}

	renderTemplate(w, "./template/greeting.html", gopher)
}

// Template rendering functiom
func renderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Error whilepassing the template: ", err)
	}
	t.Execute(w, templateData)
}

func main() {
	http.HandleFunc("/hello-gopher", helloGopherHandler)
	http.ListenAndServe(":8080", nil)
}
