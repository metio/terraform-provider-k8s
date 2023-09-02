/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package infinispan_org_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInfinispanOrgInfinispanV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_infinispan_org_infinispan_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
