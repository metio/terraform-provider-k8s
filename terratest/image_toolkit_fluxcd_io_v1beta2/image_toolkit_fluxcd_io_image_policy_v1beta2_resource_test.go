/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package image_toolkit_fluxcd_io_v1beta2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestImageToolkitFluxcdIoImagePolicyV1Beta2Resource(t *testing.T) {
	path := "../../examples/resources/k8s_image_toolkit_fluxcd_io_image_policy_v1beta2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
