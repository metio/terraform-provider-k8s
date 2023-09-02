/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package mattermost_com_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMattermostComClusterInstallationV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_mattermost_com_cluster_installation_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
