/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
