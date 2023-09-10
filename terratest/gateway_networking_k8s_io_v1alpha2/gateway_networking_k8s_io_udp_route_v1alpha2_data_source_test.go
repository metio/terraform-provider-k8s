/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package gateway_networking_k8s_io_v1alpha2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGatewayNetworkingK8SIoUdprouteV1Alpha2DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_gateway_networking_k8s_io_udp_route_v1alpha2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
