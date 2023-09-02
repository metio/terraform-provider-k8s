/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package keycloak_org_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestKeycloakOrgKeycloakClientV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_keycloak_org_keycloak_client_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}