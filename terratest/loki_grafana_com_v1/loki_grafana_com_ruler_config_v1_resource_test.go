/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package loki_grafana_com_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLokiGrafanaComRulerConfigV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_loki_grafana_com_ruler_config_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
