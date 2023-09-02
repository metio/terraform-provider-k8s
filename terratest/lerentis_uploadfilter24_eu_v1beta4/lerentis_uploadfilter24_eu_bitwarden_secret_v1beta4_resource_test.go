/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package lerentis_uploadfilter24_eu_v1beta4

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLerentisUploadfilter24EuBitwardenSecretV1Beta4Resource(t *testing.T) {
	path := "../../examples/resources/k8s_lerentis_uploadfilter24_eu_bitwarden_secret_v1beta4"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
