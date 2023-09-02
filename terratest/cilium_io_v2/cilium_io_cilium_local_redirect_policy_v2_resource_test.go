/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package cilium_io_v2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCiliumIoCiliumLocalRedirectPolicyV2Resource(t *testing.T) {
	path := "../../examples/resources/k8s_cilium_io_cilium_local_redirect_policy_v2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
