/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package validators

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

type dateTimeValidator struct{}

var _ tfsdk.AttributeValidator = (*dateTimeValidator)(nil)

func DateTime64Validator() tfsdk.AttributeValidator {
	return &dateTimeValidator{}
}

func (v dateTimeValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v dateTimeValidator) MarkdownDescription(_ context.Context) string {
	return "Validates that strings are date/time values formatted as per RFC3339 5.6"
}

func (v dateTimeValidator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
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
	for _, format := range formats {
		if _, err := time.Parse(format, value.Value); err != nil {
			resp.Diagnostics.AddAttributeError(
				req.AttributePath,
				"Invalid Date/Time Value",
				fmt.Sprintf("The value '%s' is not a valid date/time value", value.Value),
			)
			return
		}
	}
}
