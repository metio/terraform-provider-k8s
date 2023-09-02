/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package mutations_gatekeeper_sh_v1beta1_test

import (
	"context"
	fwdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/metio/terraform-provider-k8s/internal/provider/mutations_gatekeeper_sh_v1beta1"
	"github.com/metio/terraform-provider-k8s/internal/testutilities"
	"testing"
)

func TestMutationsGatekeeperShAssignMetadataV1Beta1DataSource_ValidateSchema(t *testing.T) {
	ctx := context.Background()
	schemaRequest := fwdatasource.SchemaRequest{}
	schemaResponse := &fwdatasource.SchemaResponse{}

	mutations_gatekeeper_sh_v1beta1.NewMutationsGatekeeperShAssignMetadataV1Beta1DataSource().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)

	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}
}

func TestMutationsGatekeeperShAssignMetadataV1Beta1DataSource_ConfigurationErrors(t *testing.T) {
	testCases := map[string]testutilities.ConfigurationErrorTestCase{
		"empty-name": {
			Configuration: `
				metadata = {
					name      = ""
					
				}
			`,
			ErrorRegex: "Attribute metadata.name string length must be at least 1, got: 0",
		},
		"missing-name": {
			Configuration: `
				metadata = {
					
				}
			`,
			ErrorRegex: `Inappropriate value for attribute "metadata": attribute "name" is required`,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			testutilities.VerifyConfigurationErrors(t, "data", "k8s_mutations_gatekeeper_sh_assign_metadata_v1beta1", testCase)
		})
	}
}

func TestMutationsGatekeeperShAssignMetadataV1Beta1DataSource_OfflineUsage(t *testing.T) {
	configuration := `
		metadata = {
			name = "some"
			
		}
	`
	testutilities.VerifyCannotBeUsedOffline(t, "data", "k8s_mutations_gatekeeper_sh_assign_metadata_v1beta1", configuration)
}
