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

func TestInstallationMattermostComMattermostV1Beta1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_installation_mattermost_com_mattermost_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
