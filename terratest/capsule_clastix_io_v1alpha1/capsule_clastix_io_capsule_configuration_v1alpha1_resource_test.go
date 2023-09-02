/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package capsule_clastix_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCapsuleClastixIoCapsuleConfigurationV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_capsule_clastix_io_capsule_configuration_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
