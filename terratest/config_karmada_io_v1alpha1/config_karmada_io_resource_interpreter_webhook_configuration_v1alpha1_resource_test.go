/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package config_karmada_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConfigKarmadaIoResourceInterpreterWebhookConfigurationV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_config_karmada_io_resource_interpreter_webhook_configuration_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
