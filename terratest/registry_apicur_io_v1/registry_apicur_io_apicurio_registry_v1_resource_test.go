/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package registry_apicur_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRegistryApicurIoApicurioRegistryV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_registry_apicur_io_apicurio_registry_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
