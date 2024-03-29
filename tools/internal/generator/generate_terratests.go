/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"fmt"
)

func GenerateTerratestTests(templatePath string, outputPath string, data []*TemplateData) {
	//resourceTerratestTemplate := ParseTemplates(fmt.Sprintf("%s/resource_test.go.tmpl", templatePath))
	//dataSourceTerratestTemplate := ParseTemplates(fmt.Sprintf("%s/data_source_test.go.tmpl", templatePath))
	manifestTerratestTemplate := ParseTemplates(fmt.Sprintf("%s/manifest_test.go.tmpl", templatePath))

	for _, resource := range data {
		//resourceTargetFile := fmt.Sprintf("%s/%s/%s", outputPath, resource.Package, resource.ResourceTestFile)
		//dataSourceTargetFile := fmt.Sprintf("%s/%s/%s", outputPath, resource.Package, resource.DataSourceTestFile)
		manifestTargetFile := fmt.Sprintf("%s/%s/%s", outputPath, resource.Package, resource.ManifestTestFile)
		//resourceGeneratedFile := generateCode(resourceTargetFile, resourceTerratestTemplate, resource)
		//dataSourceGeneratedFile := generateCode(dataSourceTargetFile, dataSourceTerratestTemplate, resource)
		manifestGeneratedFile := generateCode(manifestTargetFile, manifestTerratestTemplate, resource)
		//formatCode(resourceGeneratedFile)
		//formatCode(dataSourceGeneratedFile)
		formatCode(manifestGeneratedFile)
	}
}
