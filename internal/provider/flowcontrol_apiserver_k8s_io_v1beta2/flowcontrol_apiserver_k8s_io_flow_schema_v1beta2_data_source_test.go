/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package flowcontrol_apiserver_k8s_io_v1beta2_test

import (
	"context"
	fwdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/metio/terraform-provider-k8s/internal/provider/flowcontrol_apiserver_k8s_io_v1beta2"
	"github.com/metio/terraform-provider-k8s/internal/testutilities"
	"testing"
)

func TestFlowcontrolApiserverK8SIoFlowSchemaV1Beta2DataSource_ValidateSchema(t *testing.T) {
	ctx := context.Background()
	schemaRequest := fwdatasource.SchemaRequest{}
	schemaResponse := &fwdatasource.SchemaResponse{}

	flowcontrol_apiserver_k8s_io_v1beta2.NewFlowcontrolApiserverK8SIoFlowSchemaV1Beta2DataSource().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)

	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}
}

func TestFlowcontrolApiserverK8SIoFlowSchemaV1Beta2DataSource_ConfigurationErrors(t *testing.T) {
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
			testutilities.VerifyConfigurationErrors(t, "data", "k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta2", testCase)
		})
	}
}

func TestFlowcontrolApiserverK8SIoFlowSchemaV1Beta2DataSource_OfflineUsage(t *testing.T) {
	configuration := `
		metadata = {
			name = "some"
			
		}
	`
	testutilities.VerifyCannotBeUsedOffline(t, "data", "k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta2", configuration)
}