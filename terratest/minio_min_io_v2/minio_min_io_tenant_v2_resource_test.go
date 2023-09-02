/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package minio_min_io_v2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMinioMinIoTenantV2Resource(t *testing.T) {
	path := "../../examples/resources/k8s_minio_min_io_tenant_v2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
