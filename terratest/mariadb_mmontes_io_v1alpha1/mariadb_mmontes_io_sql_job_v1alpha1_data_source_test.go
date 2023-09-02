/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package mariadb_mmontes_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMariadbMmontesIoSqlJobV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_mariadb_mmontes_io_sql_job_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
