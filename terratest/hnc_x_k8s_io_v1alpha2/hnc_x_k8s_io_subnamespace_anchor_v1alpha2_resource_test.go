/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package hnc_x_k8s_io_v1alpha2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHncXK8SIoSubnamespaceAnchorV1Alpha2Resource(t *testing.T) {
	path := "../../examples/resources/k8s_hnc_x_k8s_io_subnamespace_anchor_v1alpha2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
