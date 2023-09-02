/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package flink_apache_org_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFlinkApacheOrgFlinkSessionJobV1Beta1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_flink_apache_org_flink_session_job_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
