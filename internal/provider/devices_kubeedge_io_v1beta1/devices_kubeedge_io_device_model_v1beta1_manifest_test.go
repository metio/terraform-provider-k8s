/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package devices_kubeedge_io_v1beta1_test

import (
	"context"
	fwdatasource "github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/metio/terraform-provider-k8s/internal/provider/devices_kubeedge_io_v1beta1"
	"testing"
)

func TestDevicesKubeedgeIoDeviceModelV1Beta1Manifest_ValidateSchema(t *testing.T) {
	ctx := context.Background()
	schemaRequest := fwdatasource.SchemaRequest{}
	schemaResponse := &fwdatasource.SchemaResponse{}

	devices_kubeedge_io_v1beta1.NewDevicesKubeedgeIoDeviceModelV1Beta1Manifest().Schema(ctx, schemaRequest, schemaResponse)

	if schemaResponse.Diagnostics.HasError() {
		t.Fatalf("Schema method diagnostics: %+v", schemaResponse.Diagnostics)
	}

	diagnostics := schemaResponse.Schema.ValidateImplementation(ctx)

	if diagnostics.HasError() {
		t.Fatalf("Schema validation diagnostics: %+v", diagnostics)
	}
}
