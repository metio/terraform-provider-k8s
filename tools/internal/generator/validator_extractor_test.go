/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"testing"

	"github.com/pb33f/go-yaml"
	"github.com/stretchr/testify/assert"
)

func Test_escapeRegexPattern(t *testing.T) {
	tests := map[string]struct {
		pattern string
		escaped string
	}{
		"string-with-simple-pattern": {
			pattern: "abc",
			escaped: "`abc`",
		},
		"string-with-backtick-prefix": {
			pattern: "`abc",
			escaped: "\"`\"+`abc`",
		},
		"string-with-backtick-suffix": {
			pattern: "abc`",
			escaped: "`abc`+\"`\"",
		},
		"string-with-backtick-pattern": {
			pattern: "ab`c",
			escaped: "`ab`+\"`\"+`c`",
		},
		"string-with-backtick-only": {
			pattern: "`",
			escaped: "\"`\"",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, tt.escaped, escapeRegexPattern(tt.pattern), "escapeRegexPattern(%s)", tt.pattern)
		})
	}
}

func TestOpenapiv2EnumsKeepTheValuesThatParse(t *testing.T) {
	enums := []*yaml.Node{{Value: "1"}, {Value: "2.5"}, {Value: "x"}}
	assert.Equal(t, []int64{1}, openapiv2IntEnums(enums))
	assert.Equal(t, []float64{1, 2.5}, openapiv2FloatEnums(enums))
}
