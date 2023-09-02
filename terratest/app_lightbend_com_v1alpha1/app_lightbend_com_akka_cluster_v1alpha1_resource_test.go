/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package app_lightbend_com_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAppLightbendComAkkaClusterV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_app_lightbend_com_akka_cluster_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
