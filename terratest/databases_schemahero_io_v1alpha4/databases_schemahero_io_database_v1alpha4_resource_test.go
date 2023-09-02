/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package databases_schemahero_io_v1alpha4

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDatabasesSchemaheroIoDatabaseV1Alpha4Resource(t *testing.T) {
	path := "../../examples/resources/k8s_databases_schemahero_io_database_v1alpha4"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
