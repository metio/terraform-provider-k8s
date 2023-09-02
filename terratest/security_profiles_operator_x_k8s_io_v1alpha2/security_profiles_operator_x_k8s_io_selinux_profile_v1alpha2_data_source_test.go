/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package security_profiles_operator_x_k8s_io_v1alpha2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSecurityProfilesOperatorXK8SIoSelinuxProfileV1Alpha2DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_security_profiles_operator_x_k8s_io_selinux_profile_v1alpha2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
