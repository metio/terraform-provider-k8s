/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package schemas_schemahero_io_v1alpha4

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSchemasSchemaheroIoMigrationV1Alpha4Resource(t *testing.T) {
	path := "../../examples/resources/k8s_schemas_schemahero_io_migration_v1alpha4"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
