/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package scheduling_sigs_k8s_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSchedulingSigsK8SIoPodGroupV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_scheduling_sigs_k8s_io_pod_group_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
