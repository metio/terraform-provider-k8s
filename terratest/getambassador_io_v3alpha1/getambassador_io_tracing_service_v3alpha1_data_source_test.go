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

func TestGetambassadorIoTracingServiceV3Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_getambassador_io_tracing_service_v3alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
