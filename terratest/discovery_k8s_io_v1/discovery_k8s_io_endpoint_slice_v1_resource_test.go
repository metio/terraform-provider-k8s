/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package discovery_k8s_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDiscoveryK8SIoEndpointSliceV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_discovery_k8s_io_endpoint_slice_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
