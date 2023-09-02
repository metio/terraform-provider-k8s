/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package keda_sh_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestKedaShScaledObjectV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_keda_sh_scaled_object_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}