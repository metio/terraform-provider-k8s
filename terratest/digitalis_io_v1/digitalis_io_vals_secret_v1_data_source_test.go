/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package digitalis_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDigitalisIoValsSecretV1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_digitalis_io_vals_secret_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
