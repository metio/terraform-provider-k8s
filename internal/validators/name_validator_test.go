/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
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

func TestNameValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         attr.Value
		expectError bool
	}
	tests := map[string]testCase{
		"valid name": {
			val:         types.String{Value: "ok"},
			expectError: false,
		},
		"invalid name": {
			val:         types.String{Value: "ok/or/not"},
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
			NameValidator().Validate(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}
