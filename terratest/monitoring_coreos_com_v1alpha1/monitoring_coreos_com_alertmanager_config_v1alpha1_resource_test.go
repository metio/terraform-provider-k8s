/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package monitoring_coreos_com_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMonitoringCoreosComAlertmanagerConfigV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_monitoring_coreos_com_alertmanager_config_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
