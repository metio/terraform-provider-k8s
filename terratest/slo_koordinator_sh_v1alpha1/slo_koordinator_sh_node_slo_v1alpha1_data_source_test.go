/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package slo_koordinator_sh_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSloKoordinatorShNodeSLOV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_slo_koordinator_sh_node_slo_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
