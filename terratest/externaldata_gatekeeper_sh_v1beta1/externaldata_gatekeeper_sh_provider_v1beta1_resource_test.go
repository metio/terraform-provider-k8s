/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package externaldata_gatekeeper_sh_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestExternaldataGatekeeperShProviderV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_externaldata_gatekeeper_sh_provider_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
