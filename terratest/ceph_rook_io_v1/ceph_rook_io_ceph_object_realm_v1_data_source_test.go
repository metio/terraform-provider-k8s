/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package ceph_rook_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCephRookIoCephObjectRealmV1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_ceph_rook_io_ceph_object_realm_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}