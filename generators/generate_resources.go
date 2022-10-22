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

var resourceTemplate *template.Template

func init() {
	cwd, err := currentDirectory()
	if err != nil {
		log.Fatal(err)
	}
	resourceTemplate, err = template.ParseFiles(
		fmt.Sprintf("%s/generators/templates/resource.go.tmpl", cwd),
		fmt.Sprintf("%s/generators/templates/schema_attribute.tmpl", cwd),
		fmt.Sprintf("%s/generators/templates/yaml_attribute.tmpl", cwd),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func generateResources(basePath string, data []*TemplateData) {
	for _, resource := range data {
		targetFile := fmt.Sprintf("%s/%s", basePath, resource.File)
		generatedFile := generateCode(targetFile, resourceTemplate, resource)
		formatCode(generatedFile)
	}
}
