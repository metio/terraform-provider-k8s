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

func TestLokiGrafanaComRecordingRuleV1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_loki_grafana_com_recording_rule_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
