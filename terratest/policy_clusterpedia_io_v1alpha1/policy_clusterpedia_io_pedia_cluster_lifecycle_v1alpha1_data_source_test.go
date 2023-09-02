/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package policy_clusterpedia_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPolicyClusterpediaIoPediaClusterLifecycleV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_policy_clusterpedia_io_pedia_cluster_lifecycle_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
