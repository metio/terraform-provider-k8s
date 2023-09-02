/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package apm_k8s_elastic_co_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestApmK8SElasticCoApmServerV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_apm_k8s_elastic_co_apm_server_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
