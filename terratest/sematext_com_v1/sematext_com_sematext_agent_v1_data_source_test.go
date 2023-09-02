/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package sematext_com_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSematextComSematextAgentV1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_sematext_com_sematext_agent_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
