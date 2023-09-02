/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package kuma_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestKumaIoContainerPatchV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_kuma_io_container_patch_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
