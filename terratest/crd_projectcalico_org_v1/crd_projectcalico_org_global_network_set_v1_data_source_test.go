/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package crd_projectcalico_org_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCrdProjectcalicoOrgGlobalNetworkSetV1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_crd_projectcalico_org_global_network_set_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
