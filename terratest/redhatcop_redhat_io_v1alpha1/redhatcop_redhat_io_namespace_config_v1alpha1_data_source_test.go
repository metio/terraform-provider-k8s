/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package redhatcop_redhat_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRedhatcopRedhatIoNamespaceConfigV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_redhatcop_redhat_io_namespace_config_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
