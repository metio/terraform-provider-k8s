/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package source_toolkit_fluxcd_io_v1beta2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSourceToolkitFluxcdIoBucketV1Beta2Resource(t *testing.T) {
	path := "../../examples/resources/k8s_source_toolkit_fluxcd_io_bucket_v1beta2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
