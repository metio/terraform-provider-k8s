/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package validators

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	utilValidation "k8s.io/apimachinery/pkg/util/validation"
)

type annotationValidator struct{}

var _ tfsdk.AttributeValidator = (*annotationValidator)(nil)

func AnnotationValidator() tfsdk.AttributeValidator {
	return &annotationValidator{}
}

func (v annotationValidator) Description(ctx context.Context) string {
	return v.MarkdownDescription(ctx)
}

func (v annotationValidator) MarkdownDescription(_ context.Context) string {
	return "Validate metadata.annotations according to the upstream k8s spec"
}

func (v annotationValidator) Validate(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) {
	elems, ok := validateMap(ctx, req, resp)
	if !ok {
		return
	}

	for key := range elems {
		for _, msg := range utilValidation.IsQualifiedName(key) {
			resp.Diagnostics.AddAttributeError(
				req.AttributePath,
				fmt.Sprintf("Invalid Annotation Name '%s'", key),
				msg,
			)
		}
	}
}
