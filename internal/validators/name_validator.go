/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package validators

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	apiValidation "k8s.io/apimachinery/pkg/api/validation"
)

type nameValidator struct{}

var _ tfsdk.AttributeValidator = (*nameValidator)(nil)

func NameValidator() tfsdk.AttributeValidator {
	return &nameValidator{}
}

func (validator nameValidator) Description(_ context.Context) string {
	return "Validate metadata.name according to the upstream k8s spec"
}

func (validator nameValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

func (validator nameValidator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	var value types.String

	diags := tfsdk.ValueAs(ctx, req.AttributeConfig, &value)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	if value.ValueString() == "" {
		return
	}

	for _, msg := range apiValidation.NameIsDNSSubdomain(value.ValueString(), false) {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.AttributePath,
			fmt.Sprintf("Invalid Object Name '%s'", value.ValueString()),
			msg,
		))
	}
}
