/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package reliablesyncs_kubeedge_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReliablesyncsKubeedgeIoClusterObjectSyncV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_reliablesyncs_kubeedge_io_cluster_object_sync_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
