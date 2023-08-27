/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package validators

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	utilValidation "k8s.io/apimachinery/pkg/util/validation"
)

type labelValidator struct{}

var _ validator.Map = labelValidator{}

func LabelValidator() validator.Map {
	return labelValidator{}
}

func (validator labelValidator) Description(_ context.Context) string {
	return "labels must match upstream k8s spec"
}

func (validator labelValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

func (validator labelValidator) ValidateMap(ctx context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	for key, value := range req.ConfigValue.Elements() {
		for _, msg := range utilValidation.IsQualifiedName(key) {
			resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
				req.Path,
				fmt.Sprintf("Invalid Label Key '%s'", key),
				msg,
			))
		}
		val, isString := value.(types.String)
		if !isString {
			resp.Diagnostics.Append(validatordiag.InvalidAttributeTypeDiagnostic(
				req.Path,
				fmt.Sprintf("Invalid Type in Label '%s'", key),
				fmt.Sprintf("Label values must be types.String but was %s", value.Type(ctx).String()),
			))
			continue
		}
		for _, msg := range utilValidation.IsValidLabelValue(val.ValueString()) {
			resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
				req.Path,
				fmt.Sprintf("Invalid Value in Label '%s'", key),
				msg,
			))
		}
	}
}
