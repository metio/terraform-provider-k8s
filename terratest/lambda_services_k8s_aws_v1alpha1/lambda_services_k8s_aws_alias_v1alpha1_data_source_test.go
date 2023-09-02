/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package lambda_services_k8s_aws_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLambdaServicesK8SAwsAliasV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_lambda_services_k8s_aws_alias_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}