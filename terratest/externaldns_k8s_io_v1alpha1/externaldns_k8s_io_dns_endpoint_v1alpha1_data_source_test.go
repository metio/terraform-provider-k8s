/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package externaldns_k8s_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestExternaldnsK8SIoDnsendpointV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_externaldns_k8s_io_dns_endpoint_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
