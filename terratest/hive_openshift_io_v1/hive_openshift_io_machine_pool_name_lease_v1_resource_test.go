/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package hive_openshift_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHiveOpenshiftIoMachinePoolNameLeaseV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_hive_openshift_io_machine_pool_name_lease_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
