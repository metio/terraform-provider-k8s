/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */
package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoNameExportsEveryField(t *testing.T) {
	for input, expected := range map[string]string{
		"name":      "Name",
		"-Xms":      "DashXms",
		"-XX":       "DashXX",
		"_id":       "UnderscoreId",
		"3scale":    "Threescale",
		"$ref":      "Dollarref",
		"some-name": "Some_name",
	} {
		assert.Equal(t, expected, goName(input), input)
	}
}
