/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package validators

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	apiValidation "k8s.io/apimachinery/pkg/api/validation"
)

type nameValidator struct{}

var _ validator.String = &nameValidator{}

func NameValidator() validator.String {
	return nameValidator{}
}

func (validator nameValidator) Description(_ context.Context) string {
	return "name does not match upstream k8s spec"
}

func (validator nameValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

func (validator nameValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueString()

	if value == "" {
		return
	}

	for _, msg := range apiValidation.NameIsDNSSubdomain(value, false) {
		response.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			request.Path,
			msg,
			value,
		))
	}
}
