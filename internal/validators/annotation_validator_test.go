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

func TestAnnotationValidator(t *testing.T) {
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
		"valid annotations map": {
			val: types.Map{ElemType: types.StringType, Elems: map[string]attr.Value{
				"some":                                 types.String{Value: "value"},
				"nginx.ingress.kubernetes.io/app-root": types.String{Value: "/"},
				"sidecar.jaegertracing.io/inject":      types.String{Value: "jaeger"},
			}},
			expectError: false,
		},
		"invalid annotations": {
			val: types.Map{ElemType: types.StringType, Elems: map[string]attr.Value{
				"/some/value": types.String{Value: "value"},
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
			AnnotationValidator().Validate(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}
