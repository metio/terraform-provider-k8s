/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package resources_teleport_dev_v5

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestResourcesTeleportDevTeleportRoleV5Resource(t *testing.T) {
	path := "../../examples/resources/k8s_resources_teleport_dev_teleport_role_v5"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
