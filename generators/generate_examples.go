//go:build generators

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var exampleMainTemplate *template.Template

func init() {
	cwd, err := currentDirectory()
	if err != nil {
		log.Fatal(err)
	}
	exampleMainTemplate, err = template.ParseFiles(fmt.Sprintf("%s/generators/templates/main.tf.tmpl", cwd))
	if err != nil {
		log.Fatal(err)
	}
}

func generateResourceExamples(basePath string, data []*TemplateData) {
	for _, resource := range data {
		directory := fmt.Sprintf("%s/k8s_%s", basePath, resource.TerraformResourceName)
		err := os.MkdirAll(directory, 0750)
		if err != nil {
			log.Fatal(err)
		}
		targetFile := fmt.Sprintf("%s/main.tf", directory)
		generateCode(targetFile, exampleMainTemplate, nil)
	}
}
