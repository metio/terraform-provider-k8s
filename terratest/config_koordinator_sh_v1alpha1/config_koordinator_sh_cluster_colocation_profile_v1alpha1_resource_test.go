/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package config_koordinator_sh_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConfigKoordinatorShClusterColocationProfileV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_config_koordinator_sh_cluster_colocation_profile_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
