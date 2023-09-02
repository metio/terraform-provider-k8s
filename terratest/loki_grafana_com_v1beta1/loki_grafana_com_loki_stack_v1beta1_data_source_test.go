/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package loki_grafana_com_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLokiGrafanaComLokiStackV1Beta1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_loki_grafana_com_loki_stack_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
