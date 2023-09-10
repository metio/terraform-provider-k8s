/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package nfd_k8s_sigs_io_v1alpha1_test

import (
	"context"
	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/metio/terraform-provider-k8s/internal/provider/nfd_k8s_sigs_io_v1alpha1"
	"testing"
)

func TestNfdK8SSigsIoNodeFeatureRuleV1Alpha1Resource_ValidateSchema(t *testing.T) {
	ctx := context.Background()
	schemaRequest := fwresource.SchemaRequest{}
	schemaResponse := &fwresource.SchemaResponse{}

	nfd_k8s_sigs_io_v1alpha1.NewNfdK8SSigsIoNodeFeatureRuleV1Alpha1Resource().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)

	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}
}
