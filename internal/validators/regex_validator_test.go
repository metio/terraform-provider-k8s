/*
 * SPDX-FileCopyrightText: The terraform-provider-k8scr Authors
 * SPDX-License-Identifier: 0BSD
 */

package validators

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"regexp"
	"testing"
)

func TestRegexValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         attr.Value
		regexp      *regexp.Regexp
		expectError bool
	}
	tests := map[string]testCase{
		"types.String ignored": {
			val:         types.String{Value: "oook"},
			regexp:      regexp.MustCompile(`^o[j-l]?$`),
			expectError: false,
		},
		"utilities.IntOrString valid string": {
			val: utilities.IntOrString{
				Value:      tftypes.NewValue(tftypes.String, "ok"),
				StringType: true,
			},
			regexp:      regexp.MustCompile(`^o[j-l]?$`),
			expectError: false,
		},
		"utilities.IntOrString invalid string": {
			val: utilities.IntOrString{
				Value:      tftypes.NewValue(tftypes.String, "oook"),
				StringType: true,
			},
			regexp:      regexp.MustCompile(`^o[j-l]?$`),
			expectError: true,
		},
		"utilities.IntOrString int ignored": {
			val: utilities.IntOrString{
				Value:      tftypes.NewValue(tftypes.Number, 123),
				NumberType: true,
			},
			regexp:      regexp.MustCompile(`^o[j-l]?$`),
			expectError: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			request := tfsdk.ValidateAttributeRequest{
				AttributePath:           path.Root("test"),
				AttributePathExpression: path.MatchRoot("test"),
				AttributeConfig:         test.val,
			}
			response := tfsdk.ValidateAttributeResponse{}
			RegexValidator(test.regexp).Validate(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}
