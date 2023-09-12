/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package ipam_cluster_x_k8s_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestIpamClusterXK8SIoIpaddressClaimV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_ipam_cluster_x_k8s_io_ip_address_claim_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
