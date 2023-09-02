/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package enterprise_gloo_solo_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEnterpriseGlooSoloIoAuthConfigV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_enterprise_gloo_solo_io_auth_config_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
