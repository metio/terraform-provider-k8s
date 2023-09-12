/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1beta2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInfrastructureClusterXK8SIoIbmpowerVsmachineV1Beta2Resource(t *testing.T) {
	path := "../../examples/resources/k8s_infrastructure_cluster_x_k8s_io_ibm_power_vs_machine_v1beta2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
