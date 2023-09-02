/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package hyperfoil_io_v1alpha2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHyperfoilIoHyperfoilV1Alpha2Resource(t *testing.T) {
	path := "../../examples/resources/k8s_hyperfoil_io_hyperfoil_v1alpha2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
