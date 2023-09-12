/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package jobset_x_k8s_io_v1alpha2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestJobsetXK8SIoJobSetV1Alpha2DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_jobset_x_k8s_io_job_set_v1alpha2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
