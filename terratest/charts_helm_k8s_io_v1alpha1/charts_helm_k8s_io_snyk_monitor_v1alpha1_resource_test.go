/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package charts_helm_k8s_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestChartsHelmK8SIoSnykMonitorV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_charts_helm_k8s_io_snyk_monitor_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
