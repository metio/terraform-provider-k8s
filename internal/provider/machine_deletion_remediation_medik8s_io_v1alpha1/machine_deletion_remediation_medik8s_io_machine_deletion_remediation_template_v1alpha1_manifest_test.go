/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package machine_deletion_remediation_medik8s_io_v1alpha1_test

import (
	"context"
	fwdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/metio/terraform-provider-k8s/internal/provider/machine_deletion_remediation_medik8s_io_v1alpha1"
	"testing"
)

func TestMachineDeletionRemediationMedik8SIoMachineDeletionRemediationTemplateV1Alpha1Manifest_ValidateSchema(t *testing.T) {
	ctx := context.Background()
	schemaRequest := fwdatasource.SchemaRequest{}
	schemaResponse := &fwdatasource.SchemaResponse{}

	machine_deletion_remediation_medik8s_io_v1alpha1.NewMachineDeletionRemediationMedik8SIoMachineDeletionRemediationTemplateV1Alpha1Manifest().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)

	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}
}
