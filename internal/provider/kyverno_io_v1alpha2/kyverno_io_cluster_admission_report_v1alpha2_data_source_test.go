/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kyverno_io_v1alpha2_test

import (
	"context"
	fwdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/metio/terraform-provider-k8s/internal/provider/kyverno_io_v1alpha2"
	"github.com/metio/terraform-provider-k8s/internal/testutilities"
	"testing"
)

func TestKyvernoIoClusterAdmissionReportV1Alpha2DataSource_ValidateSchema(t *testing.T) {
	ctx := context.Background()
	schemaRequest := fwdatasource.SchemaRequest{}
	schemaResponse := &fwdatasource.SchemaResponse{}

	kyverno_io_v1alpha2.NewKyvernoIoClusterAdmissionReportV1Alpha2DataSource().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)

	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}
}

func TestKyvernoIoClusterAdmissionReportV1Alpha2DataSource_ConfigurationErrors(t *testing.T) {
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
			testutilities.VerifyConfigurationErrors(t, "data", "k8s_kyverno_io_cluster_admission_report_v1alpha2", testCase)
		})
	}
}

func TestKyvernoIoClusterAdmissionReportV1Alpha2DataSource_OfflineUsage(t *testing.T) {
	configuration := `
		metadata = {
			name = "some"
			
		}
	`
	testutilities.VerifyCannotBeUsedOffline(t, "data", "k8s_kyverno_io_cluster_admission_report_v1alpha2", configuration)
}