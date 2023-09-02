/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package metal3_io_v1alpha1_test

import (
	"context"
	fwdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/metio/terraform-provider-k8s/internal/provider/metal3_io_v1alpha1"
	"github.com/metio/terraform-provider-k8s/internal/testutilities"
	"testing"
)

func TestMetal3IoHardwareDataV1Alpha1Manifest_ValidateSchema(t *testing.T) {
	ctx := context.Background()
	schemaRequest := fwdatasource.SchemaRequest{}
	schemaResponse := &fwdatasource.SchemaResponse{}

	metal3_io_v1alpha1.NewMetal3IoHardwareDataV1Alpha1Manifest().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)

	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}
}

func TestMetal3IoHardwareDataV1Alpha1Manifest_ConfigurationErrors(t *testing.T) {
	testCases := map[string]testutilities.ConfigurationErrorTestCase{
		"empty-name": {
			Configuration: `
				metadata = {
					name      = ""
					namespace = "somewhere"
				}
			`,
			ErrorRegex: "Attribute metadata.name string length must be at least 1, got: 0",
		},
		"missing-name": {
			Configuration: `
				metadata = {
					namespace = "somewhere"
				}
			`,
			ErrorRegex: `Inappropriate value for attribute "metadata": attribute "name" is required`,
		},
		"empty-namespace": {
			Configuration: `
				metadata = {
					name      = "some"
					namespace = ""
				}
			`,
			ErrorRegex: "Attribute metadata.namespace string length must be at least 1, got: 0",
		},
		"missing-namespace": {
			Configuration: `
				metadata = {
					name = "some"
				}
			`,
			ErrorRegex: `Inappropriate value for attribute "metadata": attribute "namespace" is\nrequired`,
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			testutilities.VerifyConfigurationErrors(t, "data", "k8s_metal3_io_hardware_data_v1alpha1_manifest", testCase)
		})
	}
}
