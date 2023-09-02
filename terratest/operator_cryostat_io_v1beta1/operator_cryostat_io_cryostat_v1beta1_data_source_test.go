/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package operator_cryostat_io_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOperatorCryostatIoCryostatV1Beta1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_operator_cryostat_io_cryostat_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
