/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package getambassador_io_v2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetambassadorIoModuleV2Resource(t *testing.T) {
	path := "../../examples/resources/k8s_getambassador_io_module_v2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}