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

func TestElasticacheServicesK8SAwsCacheSubnetGroupV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_elasticache_services_k8s_aws_cache_subnet_group_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
