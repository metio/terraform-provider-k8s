/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package nfd_kubernetes_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNfdKubernetesIoNodeFeatureDiscoveryV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_nfd_kubernetes_io_node_feature_discovery_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
