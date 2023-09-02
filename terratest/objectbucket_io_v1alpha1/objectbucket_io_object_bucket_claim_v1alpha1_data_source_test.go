/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package objectbucket_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestObjectbucketIoObjectBucketClaimV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_objectbucket_io_object_bucket_claim_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
