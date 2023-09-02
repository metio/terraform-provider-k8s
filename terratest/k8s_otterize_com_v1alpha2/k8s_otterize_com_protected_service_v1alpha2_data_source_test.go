/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package k8s_otterize_com_v1alpha2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestK8SOtterizeComProtectedServiceV1Alpha2DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_k8s_otterize_com_protected_service_v1alpha2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
