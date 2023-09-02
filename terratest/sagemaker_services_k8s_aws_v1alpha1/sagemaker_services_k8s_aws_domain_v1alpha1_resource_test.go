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

func TestSagemakerServicesK8SAwsDomainV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_sagemaker_services_k8s_aws_domain_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
