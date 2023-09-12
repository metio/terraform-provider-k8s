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

func TestWgpolicyk8SIoPolicyReportV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_wgpolicyk8s_io_policy_report_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
