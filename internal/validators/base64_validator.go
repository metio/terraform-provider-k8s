/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package validators

import (
	"context"
	"encoding/base64"
	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type base64Validator struct{}

var _ validator.String = base64Validator{}

func Base64Validator() validator.String {
	return base64Validator{}
}

func (validator base64Validator) Description(_ context.Context) string {
	return "value must be base64 encoded."
}

func (validator base64Validator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

func (validator base64Validator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueString()

	if value == "" {
		return
	}

	_, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		response.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			request.Path,
			validator.Description(ctx),
			value,
		))
	}
}
