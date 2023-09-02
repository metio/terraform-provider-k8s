/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package servicebinding_io_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestServicebindingIoClusterWorkloadResourceMappingV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_servicebinding_io_cluster_workload_resource_mapping_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
