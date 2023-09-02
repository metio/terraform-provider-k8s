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

func TestWorkKarmadaIoResourceBindingV1Alpha2Resource(t *testing.T) {
	path := "../../examples/resources/k8s_work_karmada_io_resource_binding_v1alpha2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
