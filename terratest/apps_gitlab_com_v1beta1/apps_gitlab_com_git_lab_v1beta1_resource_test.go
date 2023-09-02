/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package apps_gitlab_com_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAppsGitlabComGitLabV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_apps_gitlab_com_git_lab_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
