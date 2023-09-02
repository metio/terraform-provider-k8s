/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package work_karmada_io_v1alpha2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestWorkKarmadaIoResourceBindingV1Alpha2DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_work_karmada_io_resource_binding_v1alpha2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
