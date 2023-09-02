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

func TestNetworkingKarmadaIoMultiClusterServiceV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_networking_karmada_io_multi_cluster_service_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
