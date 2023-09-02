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

func TestCrdProjectcalicoOrgNetworkSetV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_crd_projectcalico_org_network_set_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
