/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_gitlab_com_v1beta2_test

import (
	"context"
	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/metio/terraform-provider-k8s/internal/provider/apps_gitlab_com_v1beta2"
	"testing"
)

func TestAppsGitlabComRunnerV1Beta2Resource_ValidateSchema(t *testing.T) {
	ctx := context.Background()
	schemaRequest := fwresource.SchemaRequest{}
	schemaResponse := &fwresource.SchemaResponse{}

	apps_gitlab_com_v1beta2.NewAppsGitlabComRunnerV1Beta2Resource().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)

	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}
}
