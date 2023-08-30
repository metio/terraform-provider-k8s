/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"fmt"
)

func GenerateResources(templatePath string, outputPath string, data []*TemplateData) {
	resourceTemplate := ParseTemplates(
		fmt.Sprintf("%s/resource.go.tmpl", templatePath),
		fmt.Sprintf("%s/read_write_schema_attribute.go.tmpl", templatePath),
		fmt.Sprintf("%s/json_attribute.go.tmpl", templatePath),
	)
	dataSourceTemplate := ParseTemplates(
		fmt.Sprintf("%s/data_source.go.tmpl", templatePath),
		fmt.Sprintf("%s/read_only_schema_attribute.go.tmpl", templatePath),
		fmt.Sprintf("%s/json_attribute.go.tmpl", templatePath),
	)
	manifestTemplate := ParseTemplates(
		fmt.Sprintf("%s/manifest.go.tmpl", templatePath),
		fmt.Sprintf("%s/read_write_schema_attribute.go.tmpl", templatePath),
		fmt.Sprintf("%s/json_attribute.go.tmpl", templatePath),
	)

	for _, resource := range data {
		resourceTargetFile := fmt.Sprintf("%s/%s/%s", outputPath, resource.Package, resource.ResourceFile)
		dataSourceTargetFile := fmt.Sprintf("%s/%s/%s", outputPath, resource.Package, resource.DataSourceFile)
		manifestTargetFile := fmt.Sprintf("%s/%s/%s", outputPath, resource.Package, resource.ManifestFile)
		resourceGeneratedFile := generateCode(resourceTargetFile, resourceTemplate, resource)
		dataSourceGeneratedFile := generateCode(dataSourceTargetFile, dataSourceTemplate, resource)
		manifestGeneratedFile := generateCode(manifestTargetFile, manifestTemplate, resource)
		formatCode(resourceGeneratedFile)
		formatCode(dataSourceGeneratedFile)
		formatCode(manifestGeneratedFile)
	}
}
