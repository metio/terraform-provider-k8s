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

func TestPrometheusserviceServicesK8SAwsRuleGroupsNamespaceV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_prometheusservice_services_k8s_aws_rule_groups_namespace_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
