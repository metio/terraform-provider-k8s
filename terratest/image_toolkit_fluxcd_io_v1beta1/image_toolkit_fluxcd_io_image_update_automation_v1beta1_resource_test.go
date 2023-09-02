/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package image_toolkit_fluxcd_io_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestImageToolkitFluxcdIoImageUpdateAutomationV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_image_toolkit_fluxcd_io_image_update_automation_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
