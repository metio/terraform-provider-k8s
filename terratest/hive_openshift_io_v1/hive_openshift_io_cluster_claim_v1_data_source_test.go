/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package hive_openshift_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHiveOpenshiftIoClusterClaimV1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_hive_openshift_io_cluster_claim_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}