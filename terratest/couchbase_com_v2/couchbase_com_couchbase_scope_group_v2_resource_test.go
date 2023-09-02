/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package couchbase_com_v2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCouchbaseComCouchbaseScopeGroupV2Resource(t *testing.T) {
	path := "../../examples/resources/k8s_couchbase_com_couchbase_scope_group_v2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}