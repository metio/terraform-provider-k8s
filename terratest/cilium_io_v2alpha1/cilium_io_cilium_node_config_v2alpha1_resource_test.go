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

func TestCiliumIoCiliumNodeConfigV2Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_cilium_io_cilium_node_config_v2alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}