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

func TestGatewayNetworkingK8SIoGatewayClassV1Alpha2Resource(t *testing.T) {
	path := "../../examples/resources/k8s_gateway_networking_k8s_io_gateway_class_v1alpha2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}