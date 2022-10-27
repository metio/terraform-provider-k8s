//go:build generators

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"text/template"
)

var exampleMainTemplate *template.Template
var exampleOutputsTemplate *template.Template
var exampleResourceTemplate *template.Template

func init() {
	cwd, err := currentDirectory()
	if err != nil {
		log.Fatal(err)
	}
	exampleMainTemplate, err = template.ParseFiles(fmt.Sprintf("%s/generators/templates/main.tf.tmpl", cwd))
	if err != nil {
		log.Fatal(err)
	}
	exampleOutputsTemplate, err = template.ParseFiles(fmt.Sprintf("%s/generators/templates/outputs.tf.tmpl", cwd))
	if err != nil {
		log.Fatal(err)
	}
	exampleResourceTemplate, err = template.ParseFiles(fmt.Sprintf("%s/generators/templates/resource.tf.tmpl", cwd))
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

		mainFile := fmt.Sprintf("%s/main.tf", directory)
		generateCode(mainFile, exampleMainTemplate, nil)

		outputsFile := fmt.Sprintf("%s/outputs.tf", directory)
		if _, err := os.Stat(outputsFile); errors.Is(err, os.ErrNotExist) {
			generateCode(outputsFile, exampleOutputsTemplate, resource)
		}

		resourceFile := fmt.Sprintf("%s/resource.tf", directory)
		if _, err := os.Stat(resourceFile); errors.Is(err, os.ErrNotExist) {
			generateCode(resourceFile, exampleResourceTemplate, resource)
		}
	}
}
