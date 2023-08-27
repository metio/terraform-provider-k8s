/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package validators

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"time"
)

type dateTimeValidator struct{}

var _ validator.String = dateTimeValidator{}

func DateTime64Validator() validator.String {
	return dateTimeValidator{}
}

func (validator dateTimeValidator) Description(_ context.Context) string {
	return "value must be date/time formatted as per RFC3339 5.6"
}

func (validator dateTimeValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

func (validator dateTimeValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueString()

	if value == "" {
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
		if _, err := time.Parse(format, value); err == nil {
			matched = true
			break
		}
	}
	if !matched {
		response.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			request.Path,
			validator.Description(ctx),
			value,
		))
	}
}
