/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package batch_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBatchCronJobV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_batch_cron_job_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
