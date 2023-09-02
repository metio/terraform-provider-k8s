/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package apps_3scale_net_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestApps3ScaleNetAPIcastV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_apps_3scale_net_ap_icast_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
