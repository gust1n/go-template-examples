package main

import (
	"go/build"
	"log"
	"os"
	"path/filepath"
)

func defaultBase(path string) string {
	p, err := build.Default.Import(path, "", build.FindOnly)
	if err != nil {
		log.Fatal(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	p.Dir, err = filepath.Rel(cwd, p.Dir)
	if err != nil {
		log.Fatal(err)
	}

	return p.Dir
}

func (t *Templates) generateTemplateName(base, path string) string {
	return filepath.ToSlash(path[len(base)+1:])
}
