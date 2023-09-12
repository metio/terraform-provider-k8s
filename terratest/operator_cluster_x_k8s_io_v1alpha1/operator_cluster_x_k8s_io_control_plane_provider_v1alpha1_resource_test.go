/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package operator_cluster_x_k8s_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOperatorClusterXK8SIoControlPlaneProviderV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_operator_cluster_x_k8s_io_control_plane_provider_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
