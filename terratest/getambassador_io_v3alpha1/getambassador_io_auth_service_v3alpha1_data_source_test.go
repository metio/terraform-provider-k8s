/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package getambassador_io_v3alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetambassadorIoAuthServiceV3Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_getambassador_io_auth_service_v3alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}