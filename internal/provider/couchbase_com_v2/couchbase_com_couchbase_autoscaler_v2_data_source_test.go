/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package couchbase_com_v2_test

import (
	"context"
	fwdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/metio/terraform-provider-k8s/internal/provider/couchbase_com_v2"
	"github.com/metio/terraform-provider-k8s/internal/testutilities"
	"testing"
)

func TestCouchbaseComCouchbaseAutoscalerV2DataSource_ValidateSchema(t *testing.T) {
	ctx := context.Background()
	schemaRequest := fwdatasource.SchemaRequest{}
	schemaResponse := &fwdatasource.SchemaResponse{}

	couchbase_com_v2.NewCouchbaseComCouchbaseAutoscalerV2DataSource().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)

	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}
}

func TestCouchbaseComCouchbaseAutoscalerV2DataSource_ConfigurationErrors(t *testing.T) {
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
			testutilities.VerifyConfigurationErrors(t, "data", "k8s_couchbase_com_couchbase_autoscaler_v2", testCase)
		})
	}
}

func TestCouchbaseComCouchbaseAutoscalerV2DataSource_OfflineUsage(t *testing.T) {
	configuration := `
		metadata = {
			name = "some"
			namespace = "somewhere"
		}
	`
	testutilities.VerifyCannotBeUsedOffline(t, "data", "k8s_couchbase_com_couchbase_autoscaler_v2", configuration)
}
