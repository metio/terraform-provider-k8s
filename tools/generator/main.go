/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import (
	"flag"
	"fmt"
	"github.com/metio/terraform-provider-k8s/tools/internal/fetcher"
	"github.com/metio/terraform-provider-k8s/tools/internal/generator"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	providerDir := flag.String("provider-dir", "", "relative or absolute path to the root provider code directory when running the command outside the root provider code directory")
	schemaDir := flag.String("schema-dir", "", "relative or absolute path to the root directory for schemas")
	parseOpenAPIv2 := flag.Bool("openapi", false, "Whether to parse OpenAPIv2 schemas")
	parseCRDv1 := flag.Bool("crd", false, "Whether to parse CRDv1 schemas")
	flag.Parse()

	if *providerDir == "" {
		log.Fatalln("No --provider-dir specified!")
	}
	if *schemaDir == "" {
		log.Fatalln("No --schema-dir specified!")
	}

	var data []*generator.TemplateData
	if *parseOpenAPIv2 {
		openapi := generator.ParseOpenAPIv2Files(fmt.Sprintf("%s/openapi_v2/", *schemaDir))
		data = append(data, generator.ConvertOpenAPIv3(openapi)...)
	}
	if *parseCRDv1 {
		crd := generator.ParseCRDv1Files(fmt.Sprintf("%s/crd_v1/", *schemaDir))
		data = append(data, generator.ConvertCRDv1(crd)...)
	}

	sort.SliceStable(data, func(i, j int) bool {
		if data[i].Package != data[j].Package {
			return data[i].Package < data[j].Package
		}
		return data[i].ResourceTypeStruct < data[j].ResourceTypeStruct
	})

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	templatesDir := filepath.Join(cwd, "..", "/tools", "/internal", "/generator", "/templates")

	generator.GenerateResources(filepath.Join(templatesDir, "/go"), fmt.Sprintf("%s/internal/provider", *providerDir), data)
	generator.GenerateProvider(filepath.Join(templatesDir, "/go"), fmt.Sprintf("%s/internal/provider", *providerDir), data)
	generator.GenerateTests(filepath.Join(templatesDir, "/go"), fmt.Sprintf("%s/internal/provider", *providerDir), data)
	generator.GenerateExamples(filepath.Join(templatesDir, "/examples"), fmt.Sprintf("%s/examples", *providerDir), data)
	generator.GenerateTemplates(filepath.Join(templatesDir, "/templates"), fmt.Sprintf("%s/templates", *providerDir), data)
	generator.GenerateTerratestTests(filepath.Join(templatesDir, "/terratest"), fmt.Sprintf("%s/terratest", *providerDir), data)
	generator.GenerateGitHubWorkflows(filepath.Join(templatesDir, "/github"), fmt.Sprintf("%s/.github/workflows", *providerDir), data)
	generator.GenerateReuseFiles(filepath.Join(templatesDir, "/reuse"), fmt.Sprintf("%s/.reuse", *providerDir), fetcher.OpenAPIv2Sources, fetcher.CRDv1Sources)
}
