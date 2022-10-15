/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package validators

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type base64Validator struct{}

var _ tfsdk.AttributeValidator = (*base64Validator)(nil)

func Base64Validator() tfsdk.AttributeValidator {
	return &base64Validator{}
}

func (v base64Validator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v base64Validator) MarkdownDescription(_ context.Context) string {
	return "Validates that strings hold base64 values."
}

func (v base64Validator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	var value types.String

	diags := tfsdk.ValueAs(ctx, req.AttributeConfig, &value)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	if value.IsUnknown() || value.IsNull() {
		return
	}

	_, err := base64.StdEncoding.DecodeString(value.Value)
	if err != nil {
		resp.Diagnostics.AddAttributeError(
			req.AttributePath,
			"Invalid Base64 Value",
			fmt.Sprintf("The value '%s' is not a valid base64 value", value.Value),
		)
	}
}
