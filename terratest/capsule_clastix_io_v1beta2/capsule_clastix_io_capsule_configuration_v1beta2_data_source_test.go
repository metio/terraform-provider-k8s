/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package capsule_clastix_io_v1beta2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCapsuleClastixIoCapsuleConfigurationV1Beta2DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_capsule_clastix_io_capsule_configuration_v1beta2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
