/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package asdb_aerospike_com_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAsdbAerospikeComAerospikeClusterV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_asdb_aerospike_com_aerospike_cluster_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
