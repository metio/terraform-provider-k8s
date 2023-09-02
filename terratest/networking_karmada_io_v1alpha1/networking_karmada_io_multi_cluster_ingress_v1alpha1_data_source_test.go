/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package networking_karmada_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNetworkingKarmadaIoMultiClusterIngressV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_networking_karmada_io_multi_cluster_ingress_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
