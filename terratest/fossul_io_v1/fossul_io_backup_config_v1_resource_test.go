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

func TestFossulIoBackupConfigV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_fossul_io_backup_config_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
