/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package networking_k8s_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNetworkingK8SIoIngressClassV1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_networking_k8s_io_ingress_class_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
