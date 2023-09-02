/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package quay_redhat_com_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestQuayRedhatComQuayRegistryV1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_quay_redhat_com_quay_registry_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
