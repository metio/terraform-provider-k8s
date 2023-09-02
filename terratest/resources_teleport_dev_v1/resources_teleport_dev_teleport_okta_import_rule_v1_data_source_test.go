/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package resources_teleport_dev_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestResourcesTeleportDevTeleportOktaImportRuleV1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_resources_teleport_dev_teleport_okta_import_rule_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
