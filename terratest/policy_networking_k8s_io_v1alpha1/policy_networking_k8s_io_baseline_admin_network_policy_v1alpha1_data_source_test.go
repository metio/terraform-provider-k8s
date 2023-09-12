/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package policy_networking_k8s_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_policy_networking_k8s_io_baseline_admin_network_policy_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
