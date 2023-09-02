/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package asdb_aerospike_com_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAsdbAerospikeComAerospikeClusterV1Beta1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_asdb_aerospike_com_aerospike_cluster_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
