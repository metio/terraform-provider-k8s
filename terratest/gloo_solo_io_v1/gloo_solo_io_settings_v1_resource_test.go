/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package gloo_solo_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGlooSoloIoSettingsV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_gloo_solo_io_settings_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}