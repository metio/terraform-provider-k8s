/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package nfd_kubernetes_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNfdKubernetesIoNodeFeatureRuleV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_nfd_kubernetes_io_node_feature_rule_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
