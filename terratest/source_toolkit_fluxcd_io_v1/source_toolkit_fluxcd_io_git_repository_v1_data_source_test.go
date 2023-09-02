/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package source_toolkit_fluxcd_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSourceToolkitFluxcdIoGitRepositoryV1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_source_toolkit_fluxcd_io_git_repository_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
