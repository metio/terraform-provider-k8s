/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package metal3_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMetal3IoHostFirmwareSettingsV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_metal3_io_host_firmware_settings_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
