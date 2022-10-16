//go:build k8s

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package k8s

import (
	"go/format"
	"log"
	"os"
	"text/template"
)

var resourceTemplate *template.Template
var providerTemplate *template.Template
var resourceDocsTemplate *template.Template
var resourceMainTemplate *template.Template
var resourceTestTemplate *template.Template
var resourceVerifyTemplate *template.Template

func init() {
	var err error
	functions := template.FuncMap{
		"snake": toSnakeCase,
	}
	resourceTemplate, err = template.New("resource.go.tmpl").Funcs(functions).
		ParseFiles(
			"templates/resource.go.tmpl",
			"templates/schema_attribute.tmpl",
			"templates/yaml_attribute.tmpl",
		)
	if err != nil {
		log.Fatal(err)
	}

	providerTemplate, err = template.New("provider.go.tmpl").Funcs(functions).
		ParseFiles("templates/provider.go.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	resourceDocsTemplate, err = template.New("resource.md.tmpl").Funcs(functions).
		ParseFiles("templates/resource.md.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	resourceMainTemplate, err = template.New("main.tf.tmpl").Funcs(functions).
		ParseFiles("templates/main.tf.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	resourceTestTemplate, err = template.New("resource_test.go.tmpl").Funcs(functions).
		ParseFiles("templates/resource_test.go.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	resourceVerifyTemplate, err = template.New("verify-resource.yaml.tmpl").Funcs(functions).
		ParseFiles("templates/verify-resource.yaml.tmpl")
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateAllTheFiles(data []*TemplateData) {
	generateResources(data)
	generateProvider(data)
	generateResourcesDocTemplates(data)
	generateResourcesDocsDirectory(data)
	generateResourceTests(data)
	generateResourceGitHubWorkflows(data)
}

func generateResourceGitHubWorkflows(data []*TemplateData) {
	for _, resource := range data {
		generateCode("../../.github/workflows/verify-"+resource.TerraformResourceName+".yml", resourceVerifyTemplate, resource)
	}
}

func generateResources(crds []*TemplateData) {
	for _, crd := range crds {
		file := generateCode("../provider/"+crd.File, resourceTemplate, crd)
		formatCode(file)
	}
}

func generateResourceTests(crds []*TemplateData) {
	for _, crd := range crds {
		file := generateCode("../../terratest/k8s_"+crd.TerraformResourceName+"_test.go", resourceTestTemplate, crd)
		formatCode(file)
	}
}

func generateResourcesDocTemplates(crds []*TemplateData) {
	for _, crd := range crds {
		generateCode("../../templates/resources/"+crd.TerraformResourceName+".md.tmpl", resourceDocsTemplate, crd)
	}
}

func generateResourcesDocsDirectory(crds []*TemplateData) {
	for _, crd := range crds {
		directory := "../../examples/resources/k8s_" + crd.TerraformResourceName
		err := os.MkdirAll(directory, 0750)
		if err != nil {
			log.Printf("error creating %s", directory)
			log.Fatal(err)
		}
		generateCode(directory+"/main.tf", resourceMainTemplate, nil)
	}
}

func generateProvider(crds []*TemplateData) {
	file := generateCode("../provider/provider.go", providerTemplate, crds)
	formatCode(file)
}

func generateCode(path string, tmpl *template.Template, data any) *os.File {
	file, err := os.Create(path)
	if err != nil {
		log.Printf("error creating %s", path)
		log.Fatal(err)
	}
	err = tmpl.Execute(file, data)
	if err != nil {
		log.Printf("error templating %s", path)
		log.Fatal(err)
	}
	return file
}

func formatCode(file *os.File) {
	unformatted, err := os.ReadFile(file.Name())
	if err != nil {
		log.Printf("error reading %s", file.Name())
		log.Fatal(err)
	}
	formatted, err := format.Source(unformatted)
	if err != nil {
		log.Printf("error formatting %s", file.Name())
		log.Fatal(err)
	}
	err = os.WriteFile(file.Name(), formatted, 0644)
	if err != nil {
		log.Printf("error writing %s", file.Name())
		log.Fatal(err)
	}
}
