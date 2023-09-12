/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package wgpolicyk8s_io_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestWgpolicyk8SIoClusterPolicyReportV1Beta1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_wgpolicyk8s_io_cluster_policy_report_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
