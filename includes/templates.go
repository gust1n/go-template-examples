package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Templates embeds html/template.Template to be able to declare new funcs on the struct
type Templates struct {
	template.Template
}

// AddTemplate adds a new template (raw content) to the template collection
func (t *Templates) AddTemplate(name, tpl string) error {
	_, err := t.New(name).Parse(tpl)
	return err
}

// AddTemplateFile adds a new template (from file path) to the template collection
func (t *Templates) AddTemplateFile(name, path string) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return t.AddTemplate(name, string(b))
}

// RenderTemplate renders the named template with passed data to the passed ResponseWriter
func (t *Templates) RenderTemplate(w http.ResponseWriter, name string, data map[string]interface{}) error {
	// Make sure template exists
	if t.Lookup(name) == nil {
		return fmt.Errorf("template %s not found", name)
	}

	// Always (try to) write to buffer first to properly catch errors
	buf := new(bytes.Buffer)
	if err := t.ExecuteTemplate(buf, name, data); err != nil {
		return err
	}
	_, err := buf.WriteTo(w)
	return err
}

// NewTemplates creates a new Templates instance and initializes all templates in given directory
func NewTemplates(tmplPath string) *Templates {
	t = &Templates{
		Template: *template.New(""),
	}
	// Traverse the template dir and parse all *.html templates
	filepath.Walk(tmplPath, func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if fi.IsDir() || !strings.HasSuffix(path, ".html") {
			return nil
		}

		tplName := t.generateTemplateName(tmplPath, path)
		log.Printf("found template: %s", tplName)
		return t.AddTemplateFile(tplName, path)
	})

	return t
}
