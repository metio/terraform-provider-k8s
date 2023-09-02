/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package policy_karmada_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPolicyKarmadaIoOverridePolicyV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_policy_karmada_io_override_policy_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}