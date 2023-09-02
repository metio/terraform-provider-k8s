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

func TestChaosMeshOrgPhysicalMachineChaosV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_chaos_mesh_org_physical_machine_chaos_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}