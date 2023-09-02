/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package scheduling_koordinator_sh_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSchedulingKoordinatorShPodMigrationJobV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_scheduling_koordinator_sh_pod_migration_job_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}