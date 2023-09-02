/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package templates_gatekeeper_sh_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTemplatesGatekeeperShConstraintTemplateV1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_templates_gatekeeper_sh_constraint_template_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
