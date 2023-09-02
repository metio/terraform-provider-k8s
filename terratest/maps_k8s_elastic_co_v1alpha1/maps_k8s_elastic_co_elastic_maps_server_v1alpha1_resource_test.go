/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package maps_k8s_elastic_co_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMapsK8SElasticCoElasticMapsServerV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_maps_k8s_elastic_co_elastic_maps_server_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
