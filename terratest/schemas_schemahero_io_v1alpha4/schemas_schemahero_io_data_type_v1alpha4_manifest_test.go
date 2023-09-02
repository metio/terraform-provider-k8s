/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package schemas_schemahero_io_v1alpha4

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSchemasSchemaheroIoDataTypeV1Alpha4Manifest(t *testing.T) {
	path := "../../examples/data-sources/k8s_schemas_schemahero_io_data_type_v1alpha4_manifest"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: path,
		NoColor:      true,
	})

	defer os.RemoveAll(path + "/.terraform.lock.hcl")
	defer os.RemoveAll(path + "/terraform.tfstate")
	defer os.RemoveAll(path + "/terraform.tfstate.backup")
	defer os.RemoveAll(path + "/.terraform")

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApplyAndIdempotent(t, terraformOptions)

	outputMap := terraform.OutputMap(t, terraformOptions, "manifests")
	for key, value := range outputMap {
		assert.NotEmpty(t, value, fmt.Sprintf("data %s.%s did not produce an output", "k8s_schemas_schemahero_io_data_type_v1alpha4_manifest", key))
	}
}