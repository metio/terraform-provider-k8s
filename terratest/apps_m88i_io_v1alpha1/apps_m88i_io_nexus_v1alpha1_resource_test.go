/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package apps_m88i_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAppsM88IIoNexusV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_apps_m88i_io_nexus_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
