/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package validators

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
)

type regexValidator struct {
	regexp *regexp.Regexp
}

var _ validator.String = regexValidator{}

func RegexValidator(regexp *regexp.Regexp) validator.String {
	return regexValidator{regexp: regexp}
}

func (validator regexValidator) Description(_ context.Context) string {
	return fmt.Sprintf("value must match regular expression '%s'", validator.regexp)
}

func (validator regexValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

func (validator regexValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueString()

	if value == "" {
		return
	}

	if ok := validator.regexp.MatchString(value); !ok {
		response.Diagnostics.Append(validatordiag.InvalidAttributeValueMatchDiagnostic(
			request.Path,
			validator.Description(ctx),
			value,
		))
	}
}
