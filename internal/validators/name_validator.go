/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package validators

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	apiValidation "k8s.io/apimachinery/pkg/api/validation"
)

type nameValidator struct{}

var _ tfsdk.AttributeValidator = (*nameValidator)(nil)

func NameValidator() tfsdk.AttributeValidator {
	return &nameValidator{}
}

func (v nameValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v nameValidator) MarkdownDescription(_ context.Context) string {
	return "Validate metadata.name according to the upstream k8s spec"
}

func (v nameValidator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	var value types.String

	diags := tfsdk.ValueAs(ctx, req.AttributeConfig, &value)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	if value.IsUnknown() || value.IsNull() {
		return
	}

	for _, msg := range apiValidation.NameIsDNSSubdomain(value.Value, false) {
		resp.Diagnostics.AddError(value.Value, msg)
	}
}
