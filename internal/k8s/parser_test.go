//go:build k8s

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package k8s

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseAllCustomResourceDefinitions(t *testing.T) {
	crds := ParseAllCustomResourceDefinitions()

	assert.NotEmpty(t, crds)
}
