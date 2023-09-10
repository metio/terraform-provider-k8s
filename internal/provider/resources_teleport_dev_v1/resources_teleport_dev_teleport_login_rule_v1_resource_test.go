/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package resources_teleport_dev_v1_test

import (
	"context"
	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/metio/terraform-provider-k8s/internal/provider/resources_teleport_dev_v1"
	"testing"
)

func TestResourcesTeleportDevTeleportLoginRuleV1Resource_ValidateSchema(t *testing.T) {
	ctx := context.Background()
	schemaRequest := fwresource.SchemaRequest{}
	schemaResponse := &fwresource.SchemaResponse{}

	resources_teleport_dev_v1.NewResourcesTeleportDevTeleportLoginRuleV1Resource().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)

	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}
}
