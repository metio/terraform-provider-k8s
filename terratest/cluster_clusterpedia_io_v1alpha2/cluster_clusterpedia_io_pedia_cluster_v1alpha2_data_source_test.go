/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package cluster_clusterpedia_io_v1alpha2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestClusterClusterpediaIoPediaClusterV1Alpha2DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_cluster_clusterpedia_io_pedia_cluster_v1alpha2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
