/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package validators

import (
	"context"
	"fmt"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

type regexValidator struct {
	regexp *regexp.Regexp
}

var _ tfsdk.AttributeValidator = regexValidator{}

func RegexValidator(regexp *regexp.Regexp) tfsdk.AttributeValidator {
	return regexValidator{regexp: regexp}
}

func (validator regexValidator) Description(_ context.Context) string {
	return fmt.Sprintf("value must match regular expression '%s'", validator.regexp)
}

func (validator regexValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

func (validator regexValidator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	expectedType := utilities.IntOrStringType{}
	if req.AttributeConfig.Type(ctx) != expectedType {
		return
	}

	givenValue := req.AttributeConfig.(utilities.IntOrString)

	if givenValue.IsZero() {
		return
	}
	if !givenValue.StringType {
		return
	}

	var stringValue string
	err := givenValue.Value.As(&stringValue)
	if err != nil {
		return
	}

	if ok := validator.regexp.MatchString(stringValue); !ok {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueMatchDiagnostic(
			req.AttributePath,
			validator.Description(ctx),
			stringValue,
		))
	}
}
