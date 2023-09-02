/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package sagemaker_services_k8s_aws_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSagemakerServicesK8SAwsModelPackageV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_sagemaker_services_k8s_aws_model_package_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
