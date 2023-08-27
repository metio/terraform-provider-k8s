//go:build generators

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import (
	"fmt"
	"k8s.io/utils/strings/slices"
	"log"
	"text/template"
)

var providerDataSourcesTemplate *template.Template
var providerResourcesTemplate *template.Template

func init() {
	cwd, err := currentDirectory()
	if err != nil {
		log.Fatal(err)
	}
	basePath := fmt.Sprintf("%s/generators/templates/go", cwd)
	providerDataSourcesTemplate, err = template.ParseFiles(fmt.Sprintf("%s/provider_data_sources.go.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
	providerResourcesTemplate, err = template.ParseFiles(fmt.Sprintf("%s/provider_resources.go.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
}

func generateProvider(basePath string, data []*TemplateData) {
	value := providerTemplateData{
		Resources: data,
		Packages:  uniquePackages(data),
	}
	dataSourcesTarget := fmt.Sprintf("%s/provider_data_sources.go", basePath)
	dataSourcesGenerated := generateCode(dataSourcesTarget, providerDataSourcesTemplate, value)
	formatCode(dataSourcesGenerated)
	resourcesTarget := fmt.Sprintf("%s/provider_resources.go", basePath)
	resourcesGenerated := generateCode(resourcesTarget, providerResourcesTemplate, value)
	formatCode(resourcesGenerated)
}

func uniquePackages(data []*TemplateData) []string {
	packages := make([]string, 0)
	for _, d := range data {
		if !slices.Contains(packages, d.Package) {
			packages = append(packages, d.Package)
		}
	}
	return packages
}

type providerTemplateData struct {
	Resources []*TemplateData
	Packages  []string
}
