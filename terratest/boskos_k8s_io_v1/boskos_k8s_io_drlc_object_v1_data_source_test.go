/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package boskos_k8s_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBoskosK8SIoDrlcobjectV1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_boskos_k8s_io_drlc_object_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
