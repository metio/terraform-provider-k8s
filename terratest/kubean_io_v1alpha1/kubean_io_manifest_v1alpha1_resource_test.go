/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package kubean_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestKubeanIoManifestV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_kubean_io_manifest_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
