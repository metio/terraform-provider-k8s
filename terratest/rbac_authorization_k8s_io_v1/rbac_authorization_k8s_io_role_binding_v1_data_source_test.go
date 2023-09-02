/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package rbac_authorization_k8s_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRbacAuthorizationK8SIoRoleBindingV1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_rbac_authorization_k8s_io_role_binding_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}