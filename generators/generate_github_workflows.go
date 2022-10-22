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

var githubWorkflowTemplate *template.Template

func init() {
	cwd, err := currentDirectory()
	if err != nil {
		log.Fatal(err)
	}
	githubWorkflowTemplate, err = template.ParseFiles(fmt.Sprintf("%s/generators/templates/verify-resource.yaml.tmpl", cwd))
	if err != nil {
		log.Fatal(err)
	}
}

func generateGitHubWorkflows(basePath string, data []*TemplateData) {
	for _, resource := range data {
		targetFile := fmt.Sprintf("%s/verify-%s.yml", basePath, resource.TerraformResourceName)
		generateCode(targetFile, githubWorkflowTemplate, resource)
	}
}
