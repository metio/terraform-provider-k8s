/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"fmt"
)

func GenerateTests(templatePath string, outputPath string, data []*TemplateData) {
	resourceTestTemplate := ParseTemplates(fmt.Sprintf("%s/resource_test.go.tmpl", templatePath))
	dataSourceTestTemplate := ParseTemplates(fmt.Sprintf("%s/data_source_test.go.tmpl", templatePath))
	manifestTestTemplate := ParseTemplates(fmt.Sprintf("%s/manifest_test.go.tmpl", templatePath))

	for _, resource := range data {
		resourceTargetFile := fmt.Sprintf("%s/%s/%s", outputPath, resource.Package, resource.ResourceTestFile)
		dataSourceTargetFile := fmt.Sprintf("%s/%s/%s", outputPath, resource.Package, resource.DataSourceTestFile)
		manifestTargetFile := fmt.Sprintf("%s/%s/%s", outputPath, resource.Package, resource.ManifestTestFile)
		resourceGeneratedFile := generateCode(resourceTargetFile, resourceTestTemplate, resource)
		dataSourceGeneratedFile := generateCode(dataSourceTargetFile, dataSourceTestTemplate, resource)
		manifestGeneratedFile := generateCode(manifestTargetFile, manifestTestTemplate, resource)
		formatCode(resourceGeneratedFile)
		formatCode(dataSourceGeneratedFile)
		formatCode(manifestGeneratedFile)
	}
}
