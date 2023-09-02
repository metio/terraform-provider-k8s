/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package longhorn_io_v1beta2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLonghornIoBackingImageManagerV1Beta2Resource(t *testing.T) {
	path := "../../examples/resources/k8s_longhorn_io_backing_image_manager_v1beta2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
