//go:build k8s

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package k8s

import (
	"testing"
)

func TestGenerateTerraformResourcesFromCustomResourceDefinitions(t *testing.T) {
	crds := ParseAllCustomResourceDefinitions()
	data := CRDsToTemplateData(crds, "provider")
	GenerateAllTheFiles("../provider", data)
}

func TestGenerateTerraformResourcesFromOpenApiSpec(t *testing.T) {
	definitions := ParseKubernetesOpenApi()
	data := OpenApiToTemplateData(definitions, "provider")
	GenerateAllTheFiles("../provider", data)
}

func TestGenerateAllTerraformResources(t *testing.T) {
	crds := ParseAllCustomResourceDefinitions()
	definitions := ParseKubernetesOpenApi()
	data := CRDsToTemplateData(crds, "provider")
	data = append(data, OpenApiToTemplateData(definitions, "provider")...)
	GenerateAllTheFiles("../provider", data)
}
