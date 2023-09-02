/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package infinispan_org_v2alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInfinispanOrgBatchV2Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_infinispan_org_batch_v2alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
