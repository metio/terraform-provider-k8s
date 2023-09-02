/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package wildfly_org_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestWildflyOrgWildFlyServerV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_wildfly_org_wild_fly_server_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
