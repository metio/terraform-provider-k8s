/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package networking_k8s_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNetworkingK8SIoNetworkPolicyV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_networking_k8s_io_network_policy_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}