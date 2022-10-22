//go:build generators

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import (
	"fmt"
	"log"
	"text/template"
)

var providerTemplate *template.Template

func init() {
	cwd, err := currentDirectory()
	providerTemplate, err = template.ParseFiles(fmt.Sprintf("%s/generators/templates/provider.go.tmpl", cwd))
	if err != nil {
		log.Fatal(err)
	}
}

func generateK8sProvider(basePath string, data []*TemplateData) {
	targetFile := fmt.Sprintf("%s/provider.go", basePath)
	generatedFile := generateCode(targetFile, providerTemplate, data)
	formatCode(generatedFile)
}
