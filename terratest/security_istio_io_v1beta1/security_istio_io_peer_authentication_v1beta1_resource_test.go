/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package security_istio_io_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSecurityIstioIoPeerAuthenticationV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_security_istio_io_peer_authentication_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}