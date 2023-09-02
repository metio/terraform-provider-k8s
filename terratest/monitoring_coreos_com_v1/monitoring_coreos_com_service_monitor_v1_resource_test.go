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

func TestMonitoringCoreosComServiceMonitorV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_monitoring_coreos_com_service_monitor_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}