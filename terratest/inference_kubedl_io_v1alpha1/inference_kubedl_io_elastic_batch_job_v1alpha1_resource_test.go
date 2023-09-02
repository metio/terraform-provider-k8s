/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package inference_kubedl_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInferenceKubedlIoElasticBatchJobV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_inference_kubedl_io_elastic_batch_job_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
