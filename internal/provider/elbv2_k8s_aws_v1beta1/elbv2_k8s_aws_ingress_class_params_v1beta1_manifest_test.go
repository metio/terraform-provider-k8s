/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package elbv2_k8s_aws_v1beta1_test

import (
	"context"
	fwdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/metio/terraform-provider-k8s/internal/provider/elbv2_k8s_aws_v1beta1"
	"github.com/metio/terraform-provider-k8s/internal/testutilities"
	"testing"
)

func TestElbv2K8SAwsIngressClassParamsV1Beta1Manifest_ValidateSchema(t *testing.T) {
	ctx := context.Background()
	schemaRequest := fwdatasource.SchemaRequest{}
	schemaResponse := &fwdatasource.SchemaResponse{}

	elbv2_k8s_aws_v1beta1.NewElbv2K8SAwsIngressClassParamsV1Beta1Manifest().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)

	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}
}

func TestElbv2K8SAwsIngressClassParamsV1Beta1Manifest_ConfigurationErrors(t *testing.T) {
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
			testutilities.VerifyConfigurationErrors(t, "data", "k8s_elbv2_k8s_aws_ingress_class_params_v1beta1_manifest", testCase)
		})
	}
}