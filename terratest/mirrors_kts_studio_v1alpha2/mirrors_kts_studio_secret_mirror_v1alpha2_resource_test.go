/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package mirrors_kts_studio_v1alpha2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMirrorsKtsStudioSecretMirrorV1Alpha2Resource(t *testing.T) {
	path := "../../examples/resources/k8s_mirrors_kts_studio_secret_mirror_v1alpha2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
