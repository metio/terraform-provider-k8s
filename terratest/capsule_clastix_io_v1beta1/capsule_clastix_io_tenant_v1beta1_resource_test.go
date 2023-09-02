/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package capsule_clastix_io_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCapsuleClastixIoTenantV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_capsule_clastix_io_tenant_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
