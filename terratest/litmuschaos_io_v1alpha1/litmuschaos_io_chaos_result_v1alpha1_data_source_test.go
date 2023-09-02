/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package litmuschaos_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLitmuschaosIoChaosResultV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_litmuschaos_io_chaos_result_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
