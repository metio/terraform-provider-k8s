/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package beat_k8s_elastic_co_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBeatK8SElasticCoBeatV1Beta1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_beat_k8s_elastic_co_beat_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
