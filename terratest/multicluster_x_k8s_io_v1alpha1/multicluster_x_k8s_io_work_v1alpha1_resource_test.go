/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package multicluster_x_k8s_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMulticlusterXK8SIoWorkV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_multicluster_x_k8s_io_work_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
