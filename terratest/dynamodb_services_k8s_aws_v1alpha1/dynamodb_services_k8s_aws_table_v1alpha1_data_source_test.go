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

func TestDynamodbServicesK8SAwsTableV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_dynamodb_services_k8s_aws_table_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
