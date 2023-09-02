/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package kibana_k8s_elastic_co_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestKibanaK8SElasticCoKibanaV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_kibana_k8s_elastic_co_kibana_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
