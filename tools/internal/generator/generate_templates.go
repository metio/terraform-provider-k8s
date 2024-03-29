/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"fmt"
)

func GenerateTemplates(templatePath string, outputPath string, data []*TemplateData) {
	//resourceDocsTemplate := ParseTemplates(fmt.Sprintf("%s/resource.md.tmpl", templatePath))
	//dataSourceDocsTemplate := ParseTemplates(fmt.Sprintf("%s/data_source.md.tmpl", templatePath))
	manifestDocsTemplate := ParseTemplates(fmt.Sprintf("%s/manifest.md.tmpl", templatePath))

	//resourceTemplates := fmt.Sprintf("%s/resources", outputPath)
	dataSourceTemplates := fmt.Sprintf("%s/data-sources", outputPath)

	for _, resource := range data {
		//resourceTemplateFile := fmt.Sprintf("%s/%s.md.tmpl", resourceTemplates, resource.ResourceTypeName)
		//generateCode(resourceTemplateFile, resourceDocsTemplate, resource)
		//dataSourceTemplateFile := fmt.Sprintf("%s/%s.md.tmpl", dataSourceTemplates, resource.DataSourceTypeName)
		//generateCode(dataSourceTemplateFile, dataSourceDocsTemplate, resource)
		manifestTemplateFile := fmt.Sprintf("%s/%s.md.tmpl", dataSourceTemplates, resource.ManifestTypeName)
		generateCode(manifestTemplateFile, manifestDocsTemplate, resource)
	}
}
