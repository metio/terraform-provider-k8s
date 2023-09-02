/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package scheduling_volcano_sh_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSchedulingVolcanoShQueueV1Beta1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_scheduling_volcano_sh_queue_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
