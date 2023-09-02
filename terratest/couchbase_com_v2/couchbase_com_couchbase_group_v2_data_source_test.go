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

func TestCouchbaseComCouchbaseGroupV2DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_couchbase_com_couchbase_group_v2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
