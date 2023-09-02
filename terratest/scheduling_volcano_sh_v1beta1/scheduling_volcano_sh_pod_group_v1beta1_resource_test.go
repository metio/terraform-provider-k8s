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

func TestSchedulingVolcanoShPodGroupV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_scheduling_volcano_sh_pod_group_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
