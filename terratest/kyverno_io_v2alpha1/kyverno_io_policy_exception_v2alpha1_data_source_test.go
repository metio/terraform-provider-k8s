/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package kyverno_io_v2alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestKyvernoIoPolicyExceptionV2Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_kyverno_io_policy_exception_v2alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
