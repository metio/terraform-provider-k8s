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

func TestElbv2K8SAwsTargetGroupBindingV1Beta1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_elbv2_k8s_aws_target_group_binding_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
