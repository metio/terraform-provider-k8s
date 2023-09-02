/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package autoscaling_k8s_elastic_co_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_autoscaling_k8s_elastic_co_elasticsearch_autoscaler_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
