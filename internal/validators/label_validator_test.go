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

func TestLabelValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         attr.Value
		expectError bool
	}
	tests := map[string]testCase{
		"wrong type": {
			val:         types.StringValue("ok"),
			expectError: true,
		},
		"null map": {
			val:         types.MapNull(types.StringType),
			expectError: false,
		},
		"unknown map": {
			val:         types.MapUnknown(types.StringType),
			expectError: false,
		},
		"valid labels map": {
			val: types.MapValueMust(types.StringType, map[string]attr.Value{
				"some":                         types.StringValue("value"),
				"app.kubernetes.io/name":       types.StringValue("mysql"),
				"app.kubernetes.io/instance":   types.StringValue("mysql-abcxzy"),
				"app.kubernetes.io/version":    types.StringValue("5.7.21"),
				"app.kubernetes.io/component":  types.StringValue("database"),
				"app.kubernetes.io/part-of":    types.StringValue("wordpress"),
				"app.kubernetes.io/managed-by": types.StringValue("helm"),
			}),
			expectError: false,
		},
		"invalid label name": {
			val: types.MapValueMust(types.StringType, map[string]attr.Value{
				"/some/value": types.StringValue("value"),
			}),
			expectError: true,
		},
		"invalid label value": {
			val: types.MapValueMust(types.StringType, map[string]attr.Value{
				"app.kubernetes.io/name": types.StringValue("/"),
			}),
			expectError: true,
		},
		"wrong value type": {
			val: types.MapValueMust(types.Int64Type, map[string]attr.Value{
				"app.kubernetes.io/name": types.Int64Value(123),
			}),
			expectError: true,
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
			LabelValidator().Validate(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}
