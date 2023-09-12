/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1alpha3

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInfrastructureClusterXK8SIoVsphereDeploymentZoneV1Alpha3DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_infrastructure_cluster_x_k8s_io_v_sphere_deployment_zone_v1alpha3"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
