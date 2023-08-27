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

var resourceTestTemplate *template.Template
var dataSourceTestTemplate *template.Template
var manifestTestTemplate *template.Template

func init() {
	cwd, err := currentDirectory()
	if err != nil {
		log.Fatal(err)
	}
	basePath := fmt.Sprintf("%s/generators/templates/go", cwd)
	resourceTestTemplate, err = template.ParseFiles(fmt.Sprintf("%s/resource_test.go.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
	dataSourceTestTemplate, err = template.ParseFiles(fmt.Sprintf("%s/data_source_test.go.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
	manifestTestTemplate, err = template.ParseFiles(fmt.Sprintf("%s/manifest_test.go.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
}

func generateTests(basePath string, data []*TemplateData) {
	for _, resource := range data {
		resourceTargetFile := fmt.Sprintf("%s/%s/%s", basePath, resource.Package, resource.ResourceTestFile)
		dataSourceTargetFile := fmt.Sprintf("%s/%s/%s", basePath, resource.Package, resource.DataSourceTestFile)
		manifestTargetFile := fmt.Sprintf("%s/%s/%s", basePath, resource.Package, resource.ManifestTestFile)
		resourceGeneratedFile := generateCode(resourceTargetFile, resourceTestTemplate, resource)
		dataSourceGeneratedFile := generateCode(dataSourceTargetFile, dataSourceTestTemplate, resource)
		manifestGeneratedFile := generateCode(manifestTargetFile, manifestTestTemplate, resource)
		formatCode(resourceGeneratedFile)
		formatCode(dataSourceGeneratedFile)
		formatCode(manifestGeneratedFile)
	}
}
