/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package jobsmanager_raczylo_com_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestJobsmanagerRaczyloComManagedJobV1Beta1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_jobsmanager_raczylo_com_managed_job_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
