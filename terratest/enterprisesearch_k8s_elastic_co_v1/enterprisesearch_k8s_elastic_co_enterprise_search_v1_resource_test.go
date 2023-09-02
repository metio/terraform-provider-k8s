/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package enterprisesearch_k8s_elastic_co_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEnterprisesearchK8SElasticCoEnterpriseSearchV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_enterprisesearch_k8s_elastic_co_enterprise_search_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
