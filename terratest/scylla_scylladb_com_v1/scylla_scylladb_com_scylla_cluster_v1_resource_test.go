/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package scylla_scylladb_com_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestScyllaScylladbComScyllaClusterV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_scylla_scylladb_com_scylla_cluster_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
