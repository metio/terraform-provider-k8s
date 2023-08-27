/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package provider_test

import (
	"context"
	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/metio/terraform-provider-k8s/internal/provider"
	"testing"
)

func TestExampleResourceSchema(t *testing.T) {
	ctx := context.Background()
	schemaRequest := fwresource.SchemaRequest{}
	schemaResponse := &fwresource.SchemaResponse{}

	provider.NewExampleResource().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)

	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}
}

//
//func TestExampleResource_ConfigurationErrors(t *testing.T) {
//	tests := []struct {
//		name          string
//		configuration string
//		error         string
//	}{
//		{
//			name: "empty-name",
//			configuration: `
//				metadata = {
//					name      = ""
//					namespace = "somewhere"
//				}
//			`,
//			error: "Attribute metadata.name string length must be at least 1, got: 0",
//		},
//		{
//			name: "empty-namespace",
//			configuration: `
//				metadata = {
//					name      = "some"
//					namespace = ""
//				}
//			`,
//			error: "Attribute metadata.namespace string length must be at least 1, got: 0",
//		},
//		{
//			name: "missing-name",
//			configuration: `
//				metadata = {
//					namespace = "somewhere"
//				}
//			`,
//			error: `Inappropriate value for attribute "metadata": attribute "name" is required`,
//		},
//		{
//			name: "missing-namespace",
//			configuration: `
//				metadata = {
//					name = "some"
//				}
//			`,
//			error: `Inappropriate value for attribute "metadata": attribute "namespace" is\nrequired`,
//		},
//		{
//			name: "missing-spec",
//			configuration: `
//				metadata = {
//					name      = "some"
//					namespace = "somewhere"
//				}
//			`,
//			error: `sdfgfd`,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mdb := &unstructured.Unstructured{}
//			mdb.SetUnstructuredContent(map[string]interface{}{
//				"apiVersion": "some.group.com/v1alpha1",
//				"kind":       "SomeCustomResource",
//				"metadata": map[string]interface{}{
//					"name":      "some",
//					"namespace": "somewhere",
//				},
//				"spec": map[string]interface{}{
//					"image": "something:here",
//				},
//			})
//			client := fake.NewSimpleDynamicClient(runtime.NewScheme(), mdb)
//
//			resource.UnitTest(t, resource.TestCase{
//				ProtoV6ProviderFactories: testutilities.ProviderFactories(client),
//				Steps: []resource.TestStep{
//					{
//						Config: testutilities.ProviderConfig() + fmt.Sprintf(`
//							resource "k8s_example" "test" {
//								%s
//							}
//						`, tt.configuration),
//						ExpectError: regexp.MustCompile(tt.error),
//					},
//				},
//			})
//		})
//	}
//}
