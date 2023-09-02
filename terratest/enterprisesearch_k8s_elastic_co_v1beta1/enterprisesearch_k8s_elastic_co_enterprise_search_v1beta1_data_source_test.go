/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package enterprisesearch_k8s_elastic_co_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEnterprisesearchK8SElasticCoEnterpriseSearchV1Beta1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_enterprisesearch_k8s_elastic_co_enterprise_search_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
