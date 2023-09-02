/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package digitalis_io_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDigitalisIoDbSecretV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_digitalis_io_db_secret_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
