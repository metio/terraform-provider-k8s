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
	utilValidation "k8s.io/apimachinery/pkg/util/validation"
)

type portValidator struct{}

var _ tfsdk.AttributeValidator = (*portValidator)(nil)

func PortValidator() tfsdk.AttributeValidator {
	return &portValidator{}
}

func (validator portValidator) Description(_ context.Context) string {
	return "value must be a valid port number, 0 < x < 65536."
}

func (validator portValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

func (validator portValidator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	var value types.Int64

	diags := tfsdk.ValueAs(ctx, req.AttributeConfig, &value)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	if value.IsUnknown() || value.IsNull() {
		return
	}

	for _, msg := range utilValidation.IsValidPortNum(int(value.ValueInt64())) {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.AttributePath,
			validator.Description(ctx),
			msg,
		))
	}
}
