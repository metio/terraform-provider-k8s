/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package rules_kubeedge_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRulesKubeedgeIoRuleEndpointV1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_rules_kubeedge_io_rule_endpoint_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
