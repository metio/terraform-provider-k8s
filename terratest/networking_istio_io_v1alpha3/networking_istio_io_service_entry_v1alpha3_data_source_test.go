/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package networking_istio_io_v1alpha3

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNetworkingIstioIoServiceEntryV1Alpha3DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_networking_istio_io_service_entry_v1alpha3"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}