/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package getambassador_io_v2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetambassadorIoKubernetesEndpointResolverV2DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_getambassador_io_kubernetes_endpoint_resolver_v2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
