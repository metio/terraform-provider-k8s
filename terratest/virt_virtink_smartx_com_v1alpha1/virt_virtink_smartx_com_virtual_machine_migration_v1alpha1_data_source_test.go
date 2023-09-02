/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package virt_virtink_smartx_com_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestVirtVirtinkSmartxComVirtualMachineMigrationV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_virt_virtink_smartx_com_virtual_machine_migration_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
