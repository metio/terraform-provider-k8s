/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package ibmcloud_ibm_com_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestIbmcloudIbmComComposableV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_ibmcloud_ibm_com_composable_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}