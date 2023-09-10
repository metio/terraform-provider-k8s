/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package apicodegen_apimatic_io_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestApicodegenApimaticIoApimaticV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_apicodegen_apimatic_io_api_matic_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
