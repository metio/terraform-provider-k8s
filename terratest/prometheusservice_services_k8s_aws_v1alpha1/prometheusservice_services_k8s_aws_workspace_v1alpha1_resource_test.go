/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package prometheusservice_services_k8s_aws_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPrometheusserviceServicesK8SAwsWorkspaceV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_prometheusservice_services_k8s_aws_workspace_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
