package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// includes holds all the parsed templates from the /templates/includes folder
var includes *template.Template

// the base template of everything
var baseTmpl *template.Template

// layouts
var fullwidthLayout *template.Template
var sidebarRightLayout *template.Template

// pages
var indexPage *template.Template
var profilePage *template.Template

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := baseTmpl.Execute(w, nil); err != nil {
			log.Println(err)
		}
	})
	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		if err := profilePage.Execute(w, nil); err != nil {
			log.Println(err)
		}
	})
	log.Println("web server listening at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func init() {
	tmplPath := filepath.Join(defaultBase("github.com/gust1n/go-template-examples/extends"), "templates")

	// First parse all files for basic include support
	includes = parseIncludes(filepath.Join(tmplPath, "includes"))

	// The base template of everything (includes included)
	baseTmpl = template.Must(template.Must(includes.Clone()).ParseFiles(filepath.Join(tmplPath, "base.html")))
	log.Println(baseTmpl)

	// Layouts
	fullwidthLayout = template.Must(template.Must(baseTmpl.Clone()).ParseFiles(filepath.Join(tmplPath, "layouts/fullwidth.html")))
	sidebarRightLayout = template.Must(template.Must(baseTmpl.Clone()).ParseFiles(filepath.Join(tmplPath, "layouts/sidebar-right.html")))

	// Pages
	indexPage = template.Must(template.Must(fullwidthLayout.Clone()).ParseFiles(filepath.Join(tmplPath, "pages/index.html")))
	profilePage = template.Must(template.Must(fullwidthLayout.Clone()).ParseFiles(filepath.Join(tmplPath, "pages/index.html")))
}

func parseIncludes(includesPath string) *template.Template {
	t := template.New("")

	// Traverse the template dir and parse all *.html templates
	filepath.Walk(includesPath, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if fi.IsDir() || !strings.HasSuffix(path, ".html") {
			return nil
		}

		tplName := generateTemplateName(includesPath, path)
		log.Printf("found include: %s", tplName)
		_, err = t.New(tplName).ParseFiles(path)
		return err
	})

	return t
}
