/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package gateway_nginx_org_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGatewayNginxOrgNginxGatewayV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_gateway_nginx_org_nginx_gateway_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
