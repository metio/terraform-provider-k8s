/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package operator_open_cluster_management_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOperatorOpenClusterManagementIoKlusterletV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_operator_open_cluster_management_io_klusterlet_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
