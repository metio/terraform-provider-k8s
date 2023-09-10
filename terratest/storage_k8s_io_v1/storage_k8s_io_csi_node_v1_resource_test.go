/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package storage_k8s_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestStorageK8SIoCsinodeV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_storage_k8s_io_csi_node_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
