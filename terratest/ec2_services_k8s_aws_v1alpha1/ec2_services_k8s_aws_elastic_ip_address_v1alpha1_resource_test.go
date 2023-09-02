/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package ec2_services_k8s_aws_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEc2ServicesK8SAwsElasticIPAddressV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_ec2_services_k8s_aws_elastic_ip_address_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
