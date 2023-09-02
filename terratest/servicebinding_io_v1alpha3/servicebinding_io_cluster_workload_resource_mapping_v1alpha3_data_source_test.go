/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package servicebinding_io_v1alpha3

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestServicebindingIoClusterWorkloadResourceMappingV1Alpha3DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_servicebinding_io_cluster_workload_resource_mapping_v1alpha3"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
