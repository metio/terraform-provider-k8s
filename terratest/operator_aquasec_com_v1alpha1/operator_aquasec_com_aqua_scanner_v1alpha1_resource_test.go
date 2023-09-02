/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package operator_aquasec_com_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOperatorAquasecComAquaScannerV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_operator_aquasec_com_aqua_scanner_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
