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

func TestClusterXK8SIoMachineSetV1Alpha4Resource(t *testing.T) {
	path := "../../examples/resources/k8s_cluster_x_k8s_io_machine_set_v1alpha4"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
