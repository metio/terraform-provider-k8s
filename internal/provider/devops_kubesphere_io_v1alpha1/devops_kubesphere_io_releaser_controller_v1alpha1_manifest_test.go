/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package devops_kubesphere_io_v1alpha1_test

import (
	"context"
	fwdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/metio/terraform-provider-k8s/internal/provider/devops_kubesphere_io_v1alpha1"
	"testing"
)

func TestDevopsKubesphereIoReleaserControllerV1Alpha1Manifest_ValidateSchema(t *testing.T) {
	ctx := context.Background()
	schemaRequest := fwdatasource.SchemaRequest{}
	schemaResponse := &fwdatasource.SchemaResponse{}

	devops_kubesphere_io_v1alpha1.NewDevopsKubesphereIoReleaserControllerV1Alpha1Manifest().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)

	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}
}
