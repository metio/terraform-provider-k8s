/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package kyverno_io_v1alpha2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestKyvernoIoClusterAdmissionReportV1Alpha2DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_kyverno_io_cluster_admission_report_v1alpha2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
