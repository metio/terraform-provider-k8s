/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package databases_schemahero_io_v1alpha4_test

import (
	"context"
	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/metio/terraform-provider-k8s/internal/provider/databases_schemahero_io_v1alpha4"
	"testing"
)

func TestDatabasesSchemaheroIoDatabaseV1Alpha4Resource_ValidateSchema(t *testing.T) {
	ctx := context.Background()
	schemaRequest := fwresource.SchemaRequest{}
	schemaResponse := &fwresource.SchemaResponse{}

	databases_schemahero_io_v1alpha4.NewDatabasesSchemaheroIoDatabaseV1Alpha4Resource().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)

	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}
}
