/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package iot_eclipse_org_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestIotEclipseOrgDittoV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_iot_eclipse_org_ditto_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
