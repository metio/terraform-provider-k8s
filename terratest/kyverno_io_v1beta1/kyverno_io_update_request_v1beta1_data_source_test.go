/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package kyverno_io_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestKyvernoIoUpdateRequestV1Beta1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_kyverno_io_update_request_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
