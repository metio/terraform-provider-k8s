/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package operator_victoriametrics_com_v1beta1

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOperatorVictoriametricsComVlogsV1Beta1Manifest(t *testing.T) {
	path := "../../examples/data-sources/k8s_operator_victoriametrics_com_v_logs_v1beta1_manifest"

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
		assert.NotEmpty(t, value, fmt.Sprintf("data %s.%s did not produce an output", "k8s_operator_victoriametrics_com_v_logs_v1beta1_manifest", key))
	}
}
