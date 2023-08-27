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
var dataSourceTemplate *template.Template
var manifestTemplate *template.Template

func init() {
	cwd, err := currentDirectory()
	if err != nil {
		log.Fatal(err)
	}
	basePath := fmt.Sprintf("%s/generators/templates/go", cwd)
	resourceTemplate, err = template.ParseFiles(
		fmt.Sprintf("%s/resource.go.tmpl", basePath),
		fmt.Sprintf("%s/read_write_schema_attribute.go.tmpl", basePath),
		fmt.Sprintf("%s/json_attribute.go.tmpl", basePath),
	)
	if err != nil {
		log.Fatal(err)
	}
	dataSourceTemplate, err = template.ParseFiles(
		fmt.Sprintf("%s/data_source.go.tmpl", basePath),
		fmt.Sprintf("%s/read_only_schema_attribute.go.tmpl", basePath),
		fmt.Sprintf("%s/json_attribute.go.tmpl", basePath),
	)
	if err != nil {
		log.Fatal(err)
	}
	manifestTemplate, err = template.ParseFiles(
		fmt.Sprintf("%s/manifest.go.tmpl", basePath),
		fmt.Sprintf("%s/read_write_schema_attribute.go.tmpl", basePath),
		fmt.Sprintf("%s/json_attribute.go.tmpl", basePath),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func generateResources(basePath string, data []*TemplateData) {
	for _, resource := range data {
		resourceTargetFile := fmt.Sprintf("%s/%s/%s", basePath, resource.Package, resource.ResourceFile)
		dataSourceTargetFile := fmt.Sprintf("%s/%s/%s", basePath, resource.Package, resource.DataSourceFile)
		manifestTargetFile := fmt.Sprintf("%s/%s/%s", basePath, resource.Package, resource.ManifestFile)
		resourceGeneratedFile := generateCode(resourceTargetFile, resourceTemplate, resource)
		dataSourceGeneratedFile := generateCode(dataSourceTargetFile, dataSourceTemplate, resource)
		manifestGeneratedFile := generateCode(manifestTargetFile, manifestTemplate, resource)
		formatCode(resourceGeneratedFile)
		formatCode(dataSourceGeneratedFile)
		formatCode(manifestGeneratedFile)
	}
}
