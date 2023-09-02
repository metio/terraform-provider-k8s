/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package executor_testkube_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestExecutorTestkubeIoExecutorV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_executor_testkube_io_executor_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
