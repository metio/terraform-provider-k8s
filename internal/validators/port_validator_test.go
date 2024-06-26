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
	"testing"
)

func TestPortValidator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         types.Int64
		expectError bool
	}
	tests := map[string]testCase{
		"valid port": {
			val:         types.Int64Value(12345),
			expectError: false,
		},
		"invalid port": {
			val:         types.Int64Value(-12345),
			expectError: true,
		},
		"null int": {
			val:         types.Int64Null(),
			expectError: false,
		},
		"unknown int": {
			val:         types.Int64Unknown(),
			expectError: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			request := validator.Int64Request{
				Path:           path.Root("test"),
				PathExpression: path.MatchRoot("test"),
				ConfigValue:    test.val,
			}
			response := validator.Int64Response{}
			PortValidator().ValidateInt64(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}
