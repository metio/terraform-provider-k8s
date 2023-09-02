/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package source_toolkit_fluxcd_io_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSourceToolkitFluxcdIoGitRepositoryV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_source_toolkit_fluxcd_io_git_repository_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
