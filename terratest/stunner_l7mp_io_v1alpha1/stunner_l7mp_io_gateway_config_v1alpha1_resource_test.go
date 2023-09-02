/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package stunner_l7mp_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestStunnerL7MpIoGatewayConfigV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_stunner_l7mp_io_gateway_config_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
