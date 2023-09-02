/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package notification_toolkit_fluxcd_io_v1beta2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNotificationToolkitFluxcdIoReceiverV1Beta2Resource(t *testing.T) {
	path := "../../examples/resources/k8s_notification_toolkit_fluxcd_io_receiver_v1beta2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}