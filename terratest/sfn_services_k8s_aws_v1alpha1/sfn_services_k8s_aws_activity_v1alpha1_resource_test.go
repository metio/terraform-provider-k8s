/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package sfn_services_k8s_aws_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSfnServicesK8SAwsActivityV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_sfn_services_k8s_aws_activity_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
