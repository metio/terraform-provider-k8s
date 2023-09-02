/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package agent_k8s_elastic_co_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAgentK8SElasticCoAgentV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_agent_k8s_elastic_co_agent_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
