/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package runtime_cluster_x_k8s_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRuntimeClusterXK8SIoExtensionConfigV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_runtime_cluster_x_k8s_io_extension_config_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
