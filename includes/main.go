package main

import (
	"log"
	"net/http"
	"path/filepath"
)

// t is a global reference of templates
var t *Templates

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := t.RenderTemplate(w, "pages/index.html", nil); err != nil {
			log.Println(err)
		}
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if err := t.RenderTemplate(w, "pages/user/login.html", nil); err != nil {
			log.Println(err)
		}
	})
	log.Println("web server listening at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func init() {
	t = NewTemplates(filepath.Join(defaultBase("github.com/gust1n/go-template-examples/includes"), "templates"))
}
