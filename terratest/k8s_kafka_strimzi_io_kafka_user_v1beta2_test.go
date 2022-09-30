/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package terratest

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestKafkaStrimziIoKafkaUserV1Beta2Resource(t *testing.T) {
	path := "../examples/resources/k8s_kafka_strimzi_io_kafka_user_v1beta2"

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

	outputMap := terraform.OutputMap(t, terraformOptions, "resources")
	for key, value := range outputMap {
		assert.NotEmpty(t, value, fmt.Sprintf("resource %s.%s did not produce an output", "k8s_kafka_strimzi_io_kafka_user_v1beta2", key))
	}
}
