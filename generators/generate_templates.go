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

var resourceDocsTemplate *template.Template
var dataSourceDocsTemplate *template.Template
var manifestDocsTemplate *template.Template

func init() {
	cwd, err := currentDirectory()
	if err != nil {
		log.Fatal(err)
	}
	basePath := fmt.Sprintf("%s/generators/templates/templates", cwd)
	resourceDocsTemplate, err = template.ParseFiles(fmt.Sprintf("%s/resource.md.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
	dataSourceDocsTemplate, err = template.ParseFiles(fmt.Sprintf("%s/data_source.md.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
	manifestDocsTemplate, err = template.ParseFiles(fmt.Sprintf("%s/manifest.md.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
}

func generateTemplates(basePath string, data []*TemplateData) {
	resourceTemplates := fmt.Sprintf("%s/resources", basePath)
	dataSourceTemplates := fmt.Sprintf("%s/data-sources", basePath)

	for _, resource := range data {
		resourceTemplateFile := fmt.Sprintf("%s/%s.md.tmpl", resourceTemplates, resource.FullResourceTypeName)
		generateCode(resourceTemplateFile, resourceDocsTemplate, resource)
		dataSourceTemplateFile := fmt.Sprintf("%s/%s.md.tmpl", dataSourceTemplates, resource.FullDataSourceTypeName)
		generateCode(dataSourceTemplateFile, dataSourceDocsTemplate, resource)
		manifestTemplateFile := fmt.Sprintf("%s/%s.md.tmpl", dataSourceTemplates, resource.FullManifestTypeName)
		generateCode(manifestTemplateFile, manifestDocsTemplate, resource)
	}
}
