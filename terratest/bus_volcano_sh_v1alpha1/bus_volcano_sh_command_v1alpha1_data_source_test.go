/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package bus_volcano_sh_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBusVolcanoShCommandV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_bus_volcano_sh_command_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
