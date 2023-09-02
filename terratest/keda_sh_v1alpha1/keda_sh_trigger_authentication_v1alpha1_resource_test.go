/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package keda_sh_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestKedaShTriggerAuthenticationV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_keda_sh_trigger_authentication_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
