/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_redhat_com_v1alpha1_test

import (
	"context"
	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/metio/terraform-provider-k8s/internal/provider/apps_redhat_com_v1alpha1"
	"testing"
)

func TestAppsRedhatComClusterImpairmentV1Alpha1Resource_ValidateSchema(t *testing.T) {
	ctx := context.Background()
	schemaRequest := fwresource.SchemaRequest{}
	schemaResponse := &fwresource.SchemaResponse{}

	apps_redhat_com_v1alpha1.NewAppsRedhatComClusterImpairmentV1Alpha1Resource().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)

	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}
}
