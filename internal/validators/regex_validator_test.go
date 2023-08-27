/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package validators

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"regexp"
	"testing"
)

func TestRegexValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         types.String
		regexp      *regexp.Regexp
		expectError bool
	}
	tests := map[string]testCase{
		"not-matched": {
			val:         types.StringValue("oook"),
			regexp:      regexp.MustCompile(`^o[j-l]?$`),
			expectError: true,
		},
		"matched": {
			val:         types.StringValue("ok"),
			regexp:      regexp.MustCompile(`^o[j-l]?$`),
			expectError: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			request := validator.StringRequest{
				Path:           path.Root("test"),
				PathExpression: path.MatchRoot("test"),
				ConfigValue:    test.val,
			}
			response := validator.StringResponse{}
			RegexValidator(test.regexp).ValidateString(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}
