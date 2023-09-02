/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package extensions_istio_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestExtensionsIstioIoWasmPluginV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_extensions_istio_io_wasm_plugin_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
