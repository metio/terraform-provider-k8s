/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package security_profiles_operator_x_k8s_io_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSecurityProfilesOperatorXK8SIoSeccompProfileV1Beta1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
