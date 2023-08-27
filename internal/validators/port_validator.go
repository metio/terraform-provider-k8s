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

type portValidator struct{}

var _ validator.Int64 = portValidator{}

func PortValidator() validator.Int64 {
	return portValidator{}
}

func (validator portValidator) Description(_ context.Context) string {
	return "port number must be within 0 < x < 65536."
}

func (validator portValidator) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

func (validator portValidator) ValidateInt64(ctx context.Context, request validator.Int64Request, response *validator.Int64Response) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueInt64()

	for _, msg := range utilValidation.IsValidPortNum(int(value)) {
		response.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			request.Path,
			msg,
			fmt.Sprintf("%v", value),
		))
	}
}
