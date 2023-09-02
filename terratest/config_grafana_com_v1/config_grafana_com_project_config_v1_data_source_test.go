/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package config_grafana_com_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConfigGrafanaComProjectConfigV1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_config_grafana_com_project_config_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
