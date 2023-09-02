/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package kms_services_k8s_aws_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestKmsServicesK8SAwsAliasV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_kms_services_k8s_aws_alias_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
