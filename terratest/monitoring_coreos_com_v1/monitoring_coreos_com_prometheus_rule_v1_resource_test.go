/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package monitoring_coreos_com_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMonitoringCoreosComPrometheusRuleV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_monitoring_coreos_com_prometheus_rule_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}