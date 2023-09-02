/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package elbv2_k8s_aws_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestElbv2K8SAwsIngressClassParamsV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_elbv2_k8s_aws_ingress_class_params_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
