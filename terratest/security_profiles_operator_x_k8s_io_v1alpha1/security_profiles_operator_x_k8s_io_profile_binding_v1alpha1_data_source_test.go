/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package security_profiles_operator_x_k8s_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSecurityProfilesOperatorXK8SIoProfileBindingV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_security_profiles_operator_x_k8s_io_profile_binding_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
