/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package fossul_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFossulIoFossulV1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_fossul_io_fossul_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}