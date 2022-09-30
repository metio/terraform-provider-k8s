//go:build k8s

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package k8s

import (
	"testing"
)

func TestGenerateTerraformResources(t *testing.T) {
	crds := ParseAllCustomResourceDefinitions()
	data := ConvertToTemplateData(crds, "provider")
	GenerateAllTheFiles("../provider", data)
}
