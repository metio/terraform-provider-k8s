//go:build generators

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	var targetPackage string
	flag.StringVar(&targetPackage, "targetPackage", "provider", "")
	flag.Parse()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	crds := parseCRDv1Files(fmt.Sprintf("%s/schemas/crd_v1/", cwd))
	openapiv2Schemas := parseOpenAPIv2Files(fmt.Sprintf("%s/schemas/openapi_v2/", cwd))

	data := convertCRDv1(crds, targetPackage)
	data = append(data, convertOpenAPIv3(openapiv2Schemas, targetPackage)...)
	sort.SliceStable(data, func(i, j int) bool {
		return data[i].TerraformResourceType < data[j].TerraformResourceType
	})

	generateResources(fmt.Sprintf("%s/internal/provider", cwd), data)
	generateK8sProvider(fmt.Sprintf("%s/internal/provider", cwd), data)
	generateResourceExamples(fmt.Sprintf("%s/examples/resources", cwd), data)
	generateDocTemplates(fmt.Sprintf("%s/templates/resources", cwd), data)
	generateTerratests(fmt.Sprintf("%s/terratest", cwd), data)
	generateGitHubWorkflows(fmt.Sprintf("%s/.github/workflows", cwd), data)
}
