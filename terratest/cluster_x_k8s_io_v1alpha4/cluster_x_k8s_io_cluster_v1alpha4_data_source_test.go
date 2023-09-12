/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package cluster_x_k8s_io_v1alpha4

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestClusterXK8SIoClusterV1Alpha4DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_cluster_x_k8s_io_cluster_v1alpha4"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
