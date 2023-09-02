/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package external_secrets_io_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestExternalSecretsIoExternalSecretV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_external_secrets_io_external_secret_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
