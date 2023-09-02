/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package postgres_operator_crunchydata_com_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPostgresOperatorCrunchydataComPostgresClusterV1Beta1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_postgres_operator_crunchydata_com_postgres_cluster_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
