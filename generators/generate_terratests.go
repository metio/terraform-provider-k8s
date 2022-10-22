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

var terratestTemplate *template.Template

func init() {
	cwd, err := currentDirectory()
	if err != nil {
		log.Fatal(err)
	}
	terratestTemplate, err = template.ParseFiles(fmt.Sprintf("%s/generators/templates/resource_test.go.tmpl", cwd))
	if err != nil {
		log.Fatal(err)
	}
}

func generateTerratests(basePath string, data []*TemplateData) {
	for _, resource := range data {
		targetFile := fmt.Sprintf("%s/k8s_%s_test.go", basePath, resource.TerraformResourceName)
		generatedFile := generateCode(targetFile, terratestTemplate, resource)
		formatCode(generatedFile)
	}
}
