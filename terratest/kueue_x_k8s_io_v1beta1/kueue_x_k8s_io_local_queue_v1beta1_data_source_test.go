/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package kueue_x_k8s_io_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestKueueXK8SIoLocalQueueV1Beta1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_kueue_x_k8s_io_local_queue_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
