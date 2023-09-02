/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package data_fluid_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDataFluidIoGooseFSRuntimeV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_data_fluid_io_goose_fs_runtime_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
