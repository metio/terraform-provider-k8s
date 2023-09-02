/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package chaos_mesh_org_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestChaosMeshOrgAzureChaosV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_chaos_mesh_org_azure_chaos_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
