/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package flowcontrol_apiserver_k8s_io_v1beta2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFlowcontrolApiserverK8SIoPriorityLevelConfigurationV1Beta2DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_flowcontrol_apiserver_k8s_io_priority_level_configuration_v1beta2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}