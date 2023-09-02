/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package resources_teleport_dev_v6

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestResourcesTeleportDevTeleportRoleV6DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_resources_teleport_dev_teleport_role_v6"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
