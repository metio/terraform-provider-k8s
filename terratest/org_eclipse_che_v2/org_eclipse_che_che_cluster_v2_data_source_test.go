/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package org_eclipse_che_v2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOrgEclipseCheCheClusterV2DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_org_eclipse_che_che_cluster_v2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
