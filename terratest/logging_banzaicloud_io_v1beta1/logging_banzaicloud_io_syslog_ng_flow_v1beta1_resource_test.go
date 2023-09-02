/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package logging_banzaicloud_io_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoggingBanzaicloudIoSyslogNGFlowV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_logging_banzaicloud_io_syslog_ng_flow_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}