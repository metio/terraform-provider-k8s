/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package binding_operators_coreos_com_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBindingOperatorsCoreosComServiceBindingV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_binding_operators_coreos_com_service_binding_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
