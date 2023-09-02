/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package helm_toolkit_fluxcd_io_v2beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHelmToolkitFluxcdIoHelmReleaseV2Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_helm_toolkit_fluxcd_io_helm_release_v2beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
