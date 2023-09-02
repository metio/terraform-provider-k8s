/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package flagger_app_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFlaggerAppAlertProviderV1Beta1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_flagger_app_alert_provider_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
