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

func TestMonitoringCoreosComAlertmanagerV1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_monitoring_coreos_com_alertmanager_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
