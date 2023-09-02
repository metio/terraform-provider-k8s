/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package topology_node_k8s_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTopologyNodeK8SIoNodeResourceTopologyV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_topology_node_k8s_io_node_resource_topology_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
