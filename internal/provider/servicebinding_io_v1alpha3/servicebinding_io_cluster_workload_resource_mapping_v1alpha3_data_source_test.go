/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package servicebinding_io_v1alpha3_test

import (
	"context"
	fwdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/metio/terraform-provider-k8s/internal/provider/servicebinding_io_v1alpha3"
	"github.com/metio/terraform-provider-k8s/internal/testutilities"
	"testing"
)

func TestServicebindingIoClusterWorkloadResourceMappingV1Alpha3DataSource_ValidateSchema(t *testing.T) {
	ctx := context.Background()
	schemaRequest := fwdatasource.SchemaRequest{}
	schemaResponse := &fwdatasource.SchemaResponse{}

	servicebinding_io_v1alpha3.NewServicebindingIoClusterWorkloadResourceMappingV1Alpha3DataSource().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)

	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}
}

func TestServicebindingIoClusterWorkloadResourceMappingV1Alpha3DataSource_ConfigurationErrors(t *testing.T) {
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
			testutilities.VerifyConfigurationErrors(t, "data", "k8s_servicebinding_io_cluster_workload_resource_mapping_v1alpha3", testCase)
		})
	}
}

func TestServicebindingIoClusterWorkloadResourceMappingV1Alpha3DataSource_OfflineUsage(t *testing.T) {
	configuration := `
		metadata = {
			name = "some"
			
		}
	`
	testutilities.VerifyCannotBeUsedOffline(t, "data", "k8s_servicebinding_io_cluster_workload_resource_mapping_v1alpha3", configuration)
}
