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
	utilValidation "k8s.io/apimachinery/pkg/util/validation"
)

type annotationValidator struct{}

var _ tfsdk.AttributeValidator = (*annotationValidator)(nil)

func AnnotationValidator() tfsdk.AttributeValidator {
	return &annotationValidator{}
}

func (validator annotationValidator) Description(_ context.Context) string {
	return "Validate metadata.annotations according to the upstream k8s spec"
}

func (validator annotationValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

func (validator annotationValidator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	elems, ok := validateMap(ctx, req, resp)
	if !ok {
		return
	}

	for key := range elems {
		for _, msg := range utilValidation.IsQualifiedName(key) {
			resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
				req.AttributePath,
				fmt.Sprintf("Invalid Annotation Key '%s'", key),
				msg,
			))
		}
	}
}
