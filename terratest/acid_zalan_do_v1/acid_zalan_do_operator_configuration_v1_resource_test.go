/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package acid_zalan_do_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAcidZalanDoOperatorConfigurationV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_acid_zalan_do_operator_configuration_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
