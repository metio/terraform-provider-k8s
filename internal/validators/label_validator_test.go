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
			val:         types.String{Value: "ok"},
			expectError: true,
		},
		"null map": {
			val:         types.Map{Null: true, ElemType: types.StringType},
			expectError: false,
		},
		"unknown map": {
			val:         types.Map{Unknown: true, ElemType: types.StringType},
			expectError: false,
		},
		"valid labels map": {
			val: types.Map{ElemType: types.StringType, Elems: map[string]attr.Value{
				"some":                         types.String{Value: "value"},
				"app.kubernetes.io/name":       types.String{Value: "mysql"},
				"app.kubernetes.io/instance":   types.String{Value: "mysql-abcxzy"},
				"app.kubernetes.io/version":    types.String{Value: "5.7.21"},
				"app.kubernetes.io/component":  types.String{Value: "database"},
				"app.kubernetes.io/part-of":    types.String{Value: "wordpress"},
				"app.kubernetes.io/managed-by": types.String{Value: "helm"},
			}},
			expectError: false,
		},
		"invalid label name": {
			val: types.Map{ElemType: types.StringType, Elems: map[string]attr.Value{
				"/some/value": types.String{Value: "value"},
			}},
			expectError: true,
		},
		"invalid label value": {
			val: types.Map{ElemType: types.StringType, Elems: map[string]attr.Value{
				"app.kubernetes.io/name": types.String{Value: "/"},
			}},
			expectError: true,
		},
		"wrong value type": {
			val: types.Map{ElemType: types.Int64Type, Elems: map[string]attr.Value{
				"app.kubernetes.io/name": types.Int64{Value: 123},
			}},
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
