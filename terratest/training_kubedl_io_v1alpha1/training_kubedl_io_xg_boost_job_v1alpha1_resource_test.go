/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package training_kubedl_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTrainingKubedlIoXGBoostJobV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_training_kubedl_io_xg_boost_job_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
