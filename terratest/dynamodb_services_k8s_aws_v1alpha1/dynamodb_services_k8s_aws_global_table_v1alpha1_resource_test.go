/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package dynamodb_services_k8s_aws_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDynamodbServicesK8SAwsGlobalTableV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_dynamodb_services_k8s_aws_global_table_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
