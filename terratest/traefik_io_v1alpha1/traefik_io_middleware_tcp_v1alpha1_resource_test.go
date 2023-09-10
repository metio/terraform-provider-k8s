/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package traefik_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTraefikIoMiddlewareTcpV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_traefik_io_middleware_tcp_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
