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

var githubResourceTemplate *template.Template
var githubDataSourceTemplate *template.Template
var githubManifestTemplate *template.Template

func init() {
	cwd, err := currentDirectory()
	if err != nil {
		log.Fatal(err)
	}
	basePath := fmt.Sprintf("%s/generators/templates/github", cwd)
	githubResourceTemplate, err = template.ParseFiles(fmt.Sprintf("%s/resource.yaml.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
	githubDataSourceTemplate, err = template.ParseFiles(fmt.Sprintf("%s/data_source.yaml.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
	githubManifestTemplate, err = template.ParseFiles(fmt.Sprintf("%s/manifest.yaml.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
}

func generateGitHubWorkflows(basePath string, data []*TemplateData) {
	for _, resource := range data {
		resourceTargetFile := fmt.Sprintf("%s/%s", basePath, resource.ResourceWorkflowFile)
		dataSourceTargetFile := fmt.Sprintf("%s/%s", basePath, resource.DataSourceWorkflowFile)
		manifestTargetFile := fmt.Sprintf("%s/%s", basePath, resource.ManifestWorkflowFile)
		generateCode(resourceTargetFile, githubResourceTemplate, resource)
		generateCode(dataSourceTargetFile, githubDataSourceTemplate, resource)
		generateCode(manifestTargetFile, githubManifestTemplate, resource)
	}
}
