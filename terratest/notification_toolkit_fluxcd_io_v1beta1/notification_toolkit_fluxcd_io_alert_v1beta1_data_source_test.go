/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package notification_toolkit_fluxcd_io_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNotificationToolkitFluxcdIoAlertV1Beta1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_notification_toolkit_fluxcd_io_alert_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
