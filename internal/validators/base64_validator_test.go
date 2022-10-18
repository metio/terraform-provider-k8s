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
	"testing"
)

func TestBase64Validator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         attr.Value
		expectError bool
	}
	tests := map[string]testCase{
		"valid base64": {
			val:         types.String{Value: "YmFzZTY0IGVuY29kZWQgdmFsdWU="},
			expectError: false,
		},
		"invalid base64": {
			val:         types.String{Value: "not base64 encoded"},
			expectError: true,
		},
		"wrong type": {
			val:         types.Bool{Value: true},
			expectError: true,
		},
		"null string": {
			val:         types.String{Null: true},
			expectError: false,
		},
		"unknown string": {
			val:         types.String{Unknown: true},
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
			Base64Validator().Validate(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}
