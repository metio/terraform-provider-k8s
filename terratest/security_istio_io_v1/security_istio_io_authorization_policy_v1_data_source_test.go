/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package security_istio_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSecurityIstioIoAuthorizationPolicyV1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_security_istio_io_authorization_policy_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
