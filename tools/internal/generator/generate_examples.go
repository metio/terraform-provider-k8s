/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func GenerateExamples(templatePath string, outputPath string, data []*TemplateData) {
	exampleMainTemplate := ParseTemplates(fmt.Sprintf("%s/main.tf.tmpl", templatePath))
	exampleManifestMainTemplate := ParseTemplates(fmt.Sprintf("%s/manifest_main.tf.tmpl", templatePath))
	exampleResourceOutputsTemplate := ParseTemplates(fmt.Sprintf("%s/resource_outputs.tf.tmpl", templatePath))
	exampleDataSourceOutputsTemplate := ParseTemplates(fmt.Sprintf("%s/data_source_outputs.tf.tmpl", templatePath))
	exampleManifestOutputsTemplate := ParseTemplates(fmt.Sprintf("%s/manifest_outputs.tf.tmpl", templatePath))
	exampleResourceTemplate := ParseTemplates(fmt.Sprintf("%s/resource.tf.tmpl", templatePath))
	exampleDataSourceTemplate := ParseTemplates(fmt.Sprintf("%s/data_source.tf.tmpl", templatePath))
	exampleManifestTemplate := ParseTemplates(fmt.Sprintf("%s/manifest.tf.tmpl", templatePath))
	exampleImportTemplate := ParseTemplates(fmt.Sprintf("%s/import.sh.tmpl", templatePath))

	resourceExamples := fmt.Sprintf("%s/resources", outputPath)
	dataSourceExamples := fmt.Sprintf("%s/data-sources", outputPath)

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
		dataSourceMainFile := fmt.Sprintf("%s/main.tf", dataSourceDirectory)
		manifestMainFile := fmt.Sprintf("%s/main.tf", manifestDirectory)
		generateCode(resourceMainFile, exampleMainTemplate, nil)
		generateCode(dataSourceMainFile, exampleMainTemplate, nil)
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
