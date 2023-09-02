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

func TestIotEclipseOrgHawkbitV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_iot_eclipse_org_hawkbit_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
