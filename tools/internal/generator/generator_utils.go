/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"go/format"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func ParseTemplates(filenames ...string) *template.Template {
	return template.Must(template.ParseFiles(filenames...))
}

func generateCode(path string, tmpl *template.Template, data any) *os.File {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Printf("error creating %s", dir)
		log.Fatal(err)
	}
	createdFile, err := os.Create(path)
	if err != nil {
		log.Printf("error creating %s", path)
		log.Fatal(err)
	}
	err = tmpl.Execute(createdFile, data)
	if err != nil {
		log.Printf("error templating %s", path)
		log.Fatal(err)
	}
	return createdFile
}

func formatCode(file *os.File) {
	unformatted, err := os.ReadFile(file.Name())
	if err != nil {
		log.Printf("error reading %s", file.Name())
		log.Fatal(err)
	}
	formatted, err := format.Source(unformatted)
	if err != nil {
		log.Printf("error formatting %s", file.Name())
		log.Fatal(err)
	}
	err = os.WriteFile(file.Name(), formatted, 0644)
	if err != nil {
		log.Printf("error writing %s", file.Name())
		log.Fatal(err)
	}
}
