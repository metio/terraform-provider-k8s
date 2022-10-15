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
	utilValidation "k8s.io/apimachinery/pkg/util/validation"
)

type labelValidator struct{}

var _ tfsdk.AttributeValidator = (*labelValidator)(nil)

func LabelValidator() tfsdk.AttributeValidator {
	return &labelValidator{}
}

func (v labelValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v labelValidator) MarkdownDescription(_ context.Context) string {
	return "Validate metadata.annotations according to the upstream k8s spec"
}

func (v labelValidator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	elems, ok := validateMap(ctx, req, resp)
	if !ok {
		return
	}

	for key, value := range elems {
		for _, msg := range utilValidation.IsQualifiedName(key) {
			resp.Diagnostics.AddAttributeError(
				req.AttributePath,
				fmt.Sprintf("Invalid Label Name '%s'", key),
				msg,
			)
		}
		val, isString := value.(types.String)
		if !isString {
			resp.Diagnostics.AddAttributeError(
				req.AttributePath,
				"Invalid Label Value",
				fmt.Sprintf("Label values must be types.String but was %s", value.Type(ctx).String()),
			)
			continue
		}
		for _, msg := range utilValidation.IsValidLabelValue(val.Value) {
			resp.Diagnostics.AddAttributeError(
				req.AttributePath,
				fmt.Sprintf("Invalid Label Value '%s'", key),
				msg,
			)
		}
	}
}
