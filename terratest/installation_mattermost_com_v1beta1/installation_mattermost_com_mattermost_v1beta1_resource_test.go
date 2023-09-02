/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package installation_mattermost_com_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInstallationMattermostComMattermostV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_installation_mattermost_com_mattermost_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
