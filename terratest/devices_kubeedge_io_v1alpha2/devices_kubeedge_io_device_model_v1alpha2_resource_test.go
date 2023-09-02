/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package devices_kubeedge_io_v1alpha2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDevicesKubeedgeIoDeviceModelV1Alpha2Resource(t *testing.T) {
	path := "../../examples/resources/k8s_devices_kubeedge_io_device_model_v1alpha2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
