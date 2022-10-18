/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package validators

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

type dateTimeValidator struct{}

var _ tfsdk.AttributeValidator = (*dateTimeValidator)(nil)

func DateTime64Validator() tfsdk.AttributeValidator {
	return &dateTimeValidator{}
}

func (validator dateTimeValidator) Description(_ context.Context) string {
	return "value must be date/time formatted as per RFC3339 5.6"
}

func (validator dateTimeValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

func (validator dateTimeValidator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	var value types.String

	diags := tfsdk.ValueAs(ctx, req.AttributeConfig, &value)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	if value.IsUnknown() || value.IsNull() {
		return
	}

	formats := []string{
		"15:04:05",
		"15:04:05Z07:00",
		"2006-01-02",
		time.RFC3339,
		time.RFC3339Nano,
	}

	matched := false
	for _, format := range formats {
		if _, err := time.Parse(format, value.Value); err == nil {
			matched = true
			break
		}
	}
	if !matched {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.AttributePath,
			validator.Description(ctx),
			value.Value,
		))
	}
}
