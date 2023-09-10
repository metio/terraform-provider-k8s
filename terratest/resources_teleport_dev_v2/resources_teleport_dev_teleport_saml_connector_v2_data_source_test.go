/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package resources_teleport_dev_v2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestResourcesTeleportDevTeleportSamlconnectorV2DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_resources_teleport_dev_teleport_saml_connector_v2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
