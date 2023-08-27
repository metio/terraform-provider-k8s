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

var resourceTerratestTemplate *template.Template
var dataSourceTerratestTemplate *template.Template
var manifestTerratestTemplate *template.Template

func init() {
	cwd, err := currentDirectory()
	if err != nil {
		log.Fatal(err)
	}
	basePath := fmt.Sprintf("%s/generators/templates/terratest", cwd)
	resourceTerratestTemplate, err = template.ParseFiles(fmt.Sprintf("%s/resource_test.go.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
	dataSourceTerratestTemplate, err = template.ParseFiles(fmt.Sprintf("%s/data_source_test.go.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
	manifestTerratestTemplate, err = template.ParseFiles(fmt.Sprintf("%s/manifest_test.go.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
}

func generateTerratestTests(basePath string, data []*TemplateData) {
	for _, resource := range data {
		resourceTargetFile := fmt.Sprintf("%s/%s/%s", basePath, resource.Package, resource.ResourceTestFile)
		dataSourceTargetFile := fmt.Sprintf("%s/%s/%s", basePath, resource.Package, resource.DataSourceTestFile)
		manifestTargetFile := fmt.Sprintf("%s/%s/%s", basePath, resource.Package, resource.ManifestTestFile)
		resourceGeneratedFile := generateCode(resourceTargetFile, resourceTerratestTemplate, resource)
		dataSourceGeneratedFile := generateCode(dataSourceTargetFile, dataSourceTerratestTemplate, resource)
		manifestGeneratedFile := generateCode(manifestTargetFile, manifestTerratestTemplate, resource)
		formatCode(resourceGeneratedFile)
		formatCode(dataSourceGeneratedFile)
		formatCode(manifestGeneratedFile)
	}
}
