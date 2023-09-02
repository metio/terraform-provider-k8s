/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package work_karmada_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestWorkKarmadaIoClusterResourceBindingV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_work_karmada_io_cluster_resource_binding_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}