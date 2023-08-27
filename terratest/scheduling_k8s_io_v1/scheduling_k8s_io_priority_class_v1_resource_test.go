/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package scheduling_k8s_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSchedulingK8SIoPriorityClassV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_scheduling_k8s_io_priority_class_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
