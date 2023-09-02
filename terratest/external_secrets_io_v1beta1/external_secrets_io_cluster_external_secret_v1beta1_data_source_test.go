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

func TestExternalSecretsIoClusterExternalSecretV1Beta1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_external_secrets_io_cluster_external_secret_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
