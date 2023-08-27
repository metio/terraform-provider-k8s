/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package admissionregistration_k8s_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAdmissionregistrationK8SIoValidatingWebhookConfigurationV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_admissionregistration_k8s_io_validating_webhook_configuration_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
