/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package operator_tigera_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOperatorTigeraIoAPIServerV1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_operator_tigera_io_api_server_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
