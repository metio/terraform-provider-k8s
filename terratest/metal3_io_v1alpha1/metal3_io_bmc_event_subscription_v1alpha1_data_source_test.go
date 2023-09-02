/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package metal3_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMetal3IoBMCEventSubscriptionV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_metal3_io_bmc_event_subscription_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
