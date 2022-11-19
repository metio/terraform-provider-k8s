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

func TestAnnotationValidator(t *testing.T) {
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
		"valid annotations map": {
			val: types.MapValueMust(types.StringType, map[string]attr.Value{
				"some":                                 types.StringValue("value"),
				"nginx.ingress.kubernetes.io/app-root": types.StringValue("/"),
				"sidecar.jaegertracing.io/inject":      types.StringValue("jaeger"),
			}),
			expectError: false,
		},
		"invalid annotations": {
			val: types.MapValueMust(types.StringType, map[string]attr.Value{
				"/some/value": types.StringValue("value"),
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
