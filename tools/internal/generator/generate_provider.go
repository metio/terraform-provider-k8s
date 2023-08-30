/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"fmt"
	"k8s.io/utils/strings/slices"
)

func GenerateProvider(templatePath string, outputPath string, data []*TemplateData) {
	providerDataSourcesTemplate := ParseTemplates(fmt.Sprintf("%s/provider_data_sources.go.tmpl", templatePath))
	providerResourcesTemplate := ParseTemplates(fmt.Sprintf("%s/provider_resources.go.tmpl", templatePath))

	value := providerTemplateData{
		Resources: data,
		Packages:  uniquePackages(data),
	}
	dataSourcesTarget := fmt.Sprintf("%s/provider_data_sources.go", outputPath)
	dataSourcesGenerated := generateCode(dataSourcesTarget, providerDataSourcesTemplate, value)
	formatCode(dataSourcesGenerated)
	resourcesTarget := fmt.Sprintf("%s/provider_resources.go", outputPath)
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
