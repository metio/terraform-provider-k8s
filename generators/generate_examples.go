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
var exampleManifestMainTemplate *template.Template
var exampleResourceOutputsTemplate *template.Template
var exampleDataSourceOutputsTemplate *template.Template
var exampleManifestOutputsTemplate *template.Template
var exampleResourceTemplate *template.Template
var exampleDataSourceTemplate *template.Template
var exampleManifestTemplate *template.Template
var exampleImportTemplate *template.Template

func init() {
	cwd, err := currentDirectory()
	if err != nil {
		log.Fatal(err)
	}
	basePath := fmt.Sprintf("%s/generators/templates/examples", cwd)
	exampleMainTemplate, err = template.ParseFiles(fmt.Sprintf("%s/main.tf.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
	exampleManifestMainTemplate, err = template.ParseFiles(fmt.Sprintf("%s/manifest_main.tf.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
	exampleResourceOutputsTemplate, err = template.ParseFiles(fmt.Sprintf("%s/resource_outputs.tf.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
	exampleDataSourceOutputsTemplate, err = template.ParseFiles(fmt.Sprintf("%s/data_source_outputs.tf.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
	exampleManifestOutputsTemplate, err = template.ParseFiles(fmt.Sprintf("%s/manifest_outputs.tf.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
	exampleResourceTemplate, err = template.ParseFiles(fmt.Sprintf("%s/resource.tf.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
	exampleDataSourceTemplate, err = template.ParseFiles(fmt.Sprintf("%s/data_source.tf.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
	exampleManifestTemplate, err = template.ParseFiles(fmt.Sprintf("%s/manifest.tf.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
	exampleImportTemplate, err = template.ParseFiles(fmt.Sprintf("%s/import.sh.tmpl", basePath))
	if err != nil {
		log.Fatal(err)
	}
}

func generateExamples(basePath string, data []*TemplateData) {
	resourceExamples := fmt.Sprintf("%s/resources", basePath)
	dataSourceExamples := fmt.Sprintf("%s/data-sources", basePath)

	for _, resource := range data {
		resourceDirectory := fmt.Sprintf("%s/%s", resourceExamples, resource.FullResourceTypeName)
		dataSourceDirectory := fmt.Sprintf("%s/%s", dataSourceExamples, resource.FullDataSourceTypeName)
		manifestDirectory := fmt.Sprintf("%s/%s", dataSourceExamples, resource.FullManifestTypeName)
		err := os.MkdirAll(resourceDirectory, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		err = os.MkdirAll(dataSourceDirectory, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		err = os.MkdirAll(manifestDirectory, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		resourceMainFile := fmt.Sprintf("%s/main.tf", resourceDirectory)
		generateCode(resourceMainFile, exampleMainTemplate, nil)
		dataSourceMainFile := fmt.Sprintf("%s/main.tf", dataSourceDirectory)
		generateCode(dataSourceMainFile, exampleMainTemplate, nil)
		manifestMainFile := fmt.Sprintf("%s/main.tf", manifestDirectory)
		generateCode(manifestMainFile, exampleManifestMainTemplate, nil)

		resourceOutputsFile := fmt.Sprintf("%s/outputs.tf", resourceDirectory)
		dataSourceOutputsFile := fmt.Sprintf("%s/outputs.tf", dataSourceDirectory)
		manifestOutputsFile := fmt.Sprintf("%s/outputs.tf", manifestDirectory)
		if _, err := os.Stat(resourceOutputsFile); errors.Is(err, os.ErrNotExist) {
			generateCode(resourceOutputsFile, exampleResourceOutputsTemplate, resource)
		}
		if _, err := os.Stat(dataSourceOutputsFile); errors.Is(err, os.ErrNotExist) {
			generateCode(dataSourceOutputsFile, exampleDataSourceOutputsTemplate, resource)
		}
		if _, err := os.Stat(manifestOutputsFile); errors.Is(err, os.ErrNotExist) {
			generateCode(manifestOutputsFile, exampleManifestOutputsTemplate, resource)
		}

		resourceTfFile := fmt.Sprintf("%s/resource.tf", resourceDirectory)
		dataSourceTfFile := fmt.Sprintf("%s/data-source.tf", dataSourceDirectory)
		manifestTfFile := fmt.Sprintf("%s/data-source.tf", manifestDirectory)
		if _, err := os.Stat(resourceTfFile); errors.Is(err, os.ErrNotExist) {
			generateCode(resourceTfFile, exampleResourceTemplate, resource)
		}
		if _, err := os.Stat(dataSourceTfFile); errors.Is(err, os.ErrNotExist) {
			generateCode(dataSourceTfFile, exampleDataSourceTemplate, resource)
		}
		if _, err := os.Stat(manifestTfFile); errors.Is(err, os.ErrNotExist) {
			generateCode(manifestTfFile, exampleManifestTemplate, resource)
		}

		importFile := fmt.Sprintf("%s/import.sh", resourceDirectory)
		generateCode(importFile, exampleImportTemplate, nil)
	}
}
