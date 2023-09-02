/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package getambassador_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetambassadorIoModuleV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_getambassador_io_module_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
