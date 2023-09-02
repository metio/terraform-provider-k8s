/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package elasticache_services_k8s_aws_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestElasticacheServicesK8SAwsUserGroupV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_elasticache_services_k8s_aws_user_group_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
