/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package s3_services_k8s_aws_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestS3ServicesK8SAwsBucketV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_s3_services_k8s_aws_bucket_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
