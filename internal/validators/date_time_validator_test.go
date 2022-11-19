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
	"time"
)

func TestDateTime64Validator(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         attr.Value
		expectError bool
	}
	tests := map[string]testCase{
		"valid time": {
			val:         types.StringValue("23:47:38"),
			expectError: false,
		},
		"invalid time": {
			val:         types.StringValue("25:47:38"),
			expectError: true,
		},
		"valid time with timezone": {
			val:         types.StringValue("23:47:38+02:00"),
			expectError: false,
		},
		"invalid time with timezone": {
			val:         types.StringValue("25:47:38+02:00"),
			expectError: true,
		},
		"valid date": {
			val:         types.StringValue("2022-10-18"),
			expectError: false,
		},
		"invalid date": {
			val:         types.StringValue("2022-13-18"),
			expectError: true,
		},
		"valid date-time": {
			val:         types.StringValue(time.Now().Format(time.RFC3339)),
			expectError: false,
		},
		"invalid date-time": {
			val:         types.StringValue("2006-13-02T15:04:05+07:00"),
			expectError: true,
		},
		"valid date-time nano": {
			val:         types.StringValue(time.Now().Format(time.RFC3339Nano)),
			expectError: false,
		},
		"invalid date-time nano": {
			val:         types.StringValue("2006-13-02T15:04:05.999999999+07:00"),
			expectError: true,
		},
		"wrong type": {
			val:         types.BoolValue(true),
			expectError: true,
		},
		"null string": {
			val:         types.StringNull(),
			expectError: false,
		},
		"unknown string": {
			val:         types.StringUnknown(),
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
			DateTime64Validator().Validate(context.TODO(), request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", response.Diagnostics)
			}
		})
	}
}
