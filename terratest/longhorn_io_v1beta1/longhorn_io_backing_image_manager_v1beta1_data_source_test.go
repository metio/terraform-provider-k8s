/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package longhorn_io_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLonghornIoBackingImageManagerV1Beta1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_longhorn_io_backing_image_manager_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
