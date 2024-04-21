/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_marin3r_3scale_net_v1alpha1_test

import (
	"context"
	fwdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/metio/terraform-provider-k8s/internal/provider/operator_marin3r_3scale_net_v1alpha1"
	"testing"
)

func TestOperatorMarin3R3ScaleNetDiscoveryServiceV1Alpha1Manifest_ValidateSchema(t *testing.T) {
	ctx := context.Background()
	schemaRequest := fwdatasource.SchemaRequest{}
	schemaResponse := &fwdatasource.SchemaResponse{}

	operator_marin3r_3scale_net_v1alpha1.NewOperatorMarin3R3ScaleNetDiscoveryServiceV1Alpha1Manifest().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)

	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}
}
