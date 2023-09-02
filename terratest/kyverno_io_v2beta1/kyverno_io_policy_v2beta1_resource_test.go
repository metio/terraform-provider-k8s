/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package kyverno_io_v2beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestKyvernoIoPolicyV2Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_kyverno_io_policy_v2beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
