/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */
package generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCRDv1Files(t *testing.T) {
	crds := ParseCRDv1Files("../../../schemas/crd_v1/")

	assert.Greater(t, len(crds), 0)
}
