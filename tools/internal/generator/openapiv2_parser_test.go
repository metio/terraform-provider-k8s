/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */
package generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseOpenAPIv2Files(t *testing.T) {
	openapiv2Schemas := ParseOpenAPIv2Files("../../../schemas/openapi_v2/")

	assert.Greater(t, len(openapiv2Schemas), 0)
}
