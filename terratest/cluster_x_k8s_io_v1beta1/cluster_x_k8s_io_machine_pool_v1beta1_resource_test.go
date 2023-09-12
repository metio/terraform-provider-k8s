/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package cluster_x_k8s_io_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestClusterXK8SIoMachinePoolV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_cluster_x_k8s_io_machine_pool_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
