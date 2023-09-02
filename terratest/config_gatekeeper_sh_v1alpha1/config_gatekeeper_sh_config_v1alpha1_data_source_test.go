/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package config_gatekeeper_sh_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConfigGatekeeperShConfigV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_config_gatekeeper_sh_config_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}