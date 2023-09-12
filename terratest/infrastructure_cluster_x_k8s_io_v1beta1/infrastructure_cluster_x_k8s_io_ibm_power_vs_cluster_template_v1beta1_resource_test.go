/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInfrastructureClusterXK8SIoIbmpowerVsclusterTemplateV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_template_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
