/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package cilium_io_v2alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCiliumIoCiliumEndpointSliceV2Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_cilium_io_cilium_endpoint_slice_v2alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
