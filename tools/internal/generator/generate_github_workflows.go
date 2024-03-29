/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"fmt"
)

func GenerateGitHubWorkflows(templatePath string, outputPath string, data []*TemplateData) {
	//githubResourceTemplate := ParseTemplates(fmt.Sprintf("%s/resource.yaml.tmpl", templatePath))
	//githubDataSourceTemplate := ParseTemplates(fmt.Sprintf("%s/data_source.yaml.tmpl", templatePath))
	githubManifestTemplate := ParseTemplates(fmt.Sprintf("%s/manifest.yaml.tmpl", templatePath))

	for _, resource := range data {
		//resourceTargetFile := fmt.Sprintf("%s/%s", outputPath, resource.ResourceWorkflowFile)
		//dataSourceTargetFile := fmt.Sprintf("%s/%s", outputPath, resource.DataSourceWorkflowFile)
		manifestTargetFile := fmt.Sprintf("%s/%s", outputPath, resource.ManifestWorkflowFile)
		//generateCode(resourceTargetFile, githubResourceTemplate, resource)
		//generateCode(dataSourceTargetFile, githubDataSourceTemplate, resource)
		generateCode(manifestTargetFile, githubManifestTemplate, resource)
	}
}
