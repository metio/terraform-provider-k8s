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

var docsTemplate *template.Template

func init() {
	cwd, err := currentDirectory()
	docsTemplate, err = template.ParseFiles(fmt.Sprintf("%s/generators/templates/resource.md.tmpl", cwd))
	if err != nil {
		log.Fatal(err)
	}
}

func generateDocTemplates(basePath string, data []*TemplateData) {
	for _, resource := range data {
		targetFile := fmt.Sprintf("%s/%s.md.tmpl", basePath, resource.TerraformResourceName)
		generateCode(targetFile, docsTemplate, resource)
	}
}
