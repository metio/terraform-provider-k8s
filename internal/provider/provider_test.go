/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package provider_test

import (
	"context"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/metio/terraform-provider-k8s/internal/provider"
	"testing"
)

func TestProvider_Schema(t *testing.T) {
	ctx := context.Background()
	schemaRequest := fwprovider.SchemaRequest{}
	schemaResponse := &fwprovider.SchemaResponse{}

	provider.New().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)
	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}
}
