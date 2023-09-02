/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package mutations_gatekeeper_sh_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMutationsGatekeeperShAssignV1Beta1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_mutations_gatekeeper_sh_assign_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}