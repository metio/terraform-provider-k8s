/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package validators

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func validateMap(ctx context.Context, req tfsdk.ValidateAttributeRequest, resp *tfsdk.ValidateAttributeResponse) (map[string]attr.Value, bool) {
	var m types.Map

	diags := tfsdk.ValueAs(ctx, req.AttributeConfig, &m)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return nil, false
	}

	if m.IsUnknown() || m.IsNull() {
		return nil, false
	}

	return m.Elements(), true
}
