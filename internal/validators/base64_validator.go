/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package validators

import (
	"context"
	"encoding/base64"
	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type base64Validator struct{}

var _ tfsdk.AttributeValidator = (*base64Validator)(nil)

func Base64Validator() tfsdk.AttributeValidator {
	return &base64Validator{}
}

func (validator base64Validator) Description(_ context.Context) string {
	return "value must be base64 encoded."
}

func (validator base64Validator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

func (validator base64Validator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
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
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.AttributePath,
			validator.Description(ctx),
			value.Value,
		))
	}
}
