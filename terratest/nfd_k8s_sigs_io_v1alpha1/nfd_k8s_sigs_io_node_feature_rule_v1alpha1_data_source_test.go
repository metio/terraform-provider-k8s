/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package nfd_k8s_sigs_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNfdK8SSigsIoNodeFeatureRuleV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_nfd_k8s_sigs_io_node_feature_rule_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}