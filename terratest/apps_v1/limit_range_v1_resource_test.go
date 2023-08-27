/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package apps_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLimitRangeV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_limit_range_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
