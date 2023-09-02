/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package expansion_gatekeeper_sh_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestExpansionGatekeeperShExpansionTemplateV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_expansion_gatekeeper_sh_expansion_template_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}