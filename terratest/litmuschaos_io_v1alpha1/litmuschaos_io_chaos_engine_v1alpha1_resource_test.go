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

func TestLitmuschaosIoChaosEngineV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_litmuschaos_io_chaos_engine_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
