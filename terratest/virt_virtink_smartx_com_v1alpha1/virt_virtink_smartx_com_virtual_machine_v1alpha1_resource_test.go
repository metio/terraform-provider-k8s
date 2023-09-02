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

func TestVirtVirtinkSmartxComVirtualMachineV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_virt_virtink_smartx_com_virtual_machine_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
