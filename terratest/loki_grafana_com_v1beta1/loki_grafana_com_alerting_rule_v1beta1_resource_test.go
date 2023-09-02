/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package loki_grafana_com_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLokiGrafanaComAlertingRuleV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_loki_grafana_com_alerting_rule_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
