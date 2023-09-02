/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package charts_opdev_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestChartsOpdevIoSynapseV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_charts_opdev_io_synapse_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
