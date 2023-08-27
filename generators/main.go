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
	"sort"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	crds := parseCRDv1Files(fmt.Sprintf("%s/schemas/crd_v1/", cwd))
	openapiv2Schemas := parseOpenAPIv2Files(fmt.Sprintf("%s/schemas/openapi_v2/", cwd))

	data := convertCRDv1(crds)
	data = append(data, convertOpenAPIv3(openapiv2Schemas)...)
	sort.SliceStable(data, func(i, j int) bool {
		if data[i].Package != data[j].Package {
			return data[i].Package < data[j].Package
		}
		return data[i].ResourceTypeStruct < data[j].ResourceTypeStruct
	})

	generateResources(fmt.Sprintf("%s/internal/provider", cwd), data)
	generateProvider(fmt.Sprintf("%s/internal/provider", cwd), data)
	generateTests(fmt.Sprintf("%s/internal/provider", cwd), data)
	generateExamples(fmt.Sprintf("%s/examples", cwd), data)
	generateTemplates(fmt.Sprintf("%s/templates", cwd), data)
	generateTerratestTests(fmt.Sprintf("%s/terratest", cwd), data)
	generateGitHubWorkflows(fmt.Sprintf("%s/.github/workflows", cwd), data)
}
