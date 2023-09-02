/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package gateway_solo_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGatewaySoloIoRouteOptionV1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_gateway_solo_io_route_option_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
