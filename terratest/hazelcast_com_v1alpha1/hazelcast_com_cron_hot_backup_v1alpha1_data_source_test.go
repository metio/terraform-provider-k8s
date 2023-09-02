/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package hazelcast_com_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHazelcastComCronHotBackupV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_hazelcast_com_cron_hot_backup_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
