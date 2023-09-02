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

func TestSchemasSchemaheroIoDataTypeV1Alpha4DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_schemas_schemahero_io_data_type_v1alpha4"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
