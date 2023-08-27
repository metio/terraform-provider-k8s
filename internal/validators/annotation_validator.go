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
	utilValidation "k8s.io/apimachinery/pkg/util/validation"
)

type annotationValidator struct{}

var _ validator.Map = annotationValidator{}

func AnnotationValidator() validator.Map {
	return annotationValidator{}
}

func (validator annotationValidator) Description(_ context.Context) string {
	return "Validate metadata.annotations according to the upstream k8s spec"
}

func (validator annotationValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

func (validator annotationValidator) ValidateMap(_ context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	for key := range req.ConfigValue.Elements() {
		for _, msg := range utilValidation.IsQualifiedName(key) {
			resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
				req.Path,
				fmt.Sprintf("Invalid Annotation Key '%s'", key),
				msg,
			))
		}
	}
}
