/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package hiveinternal_openshift_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHiveinternalOpenshiftIoFakeClusterInstallV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_hiveinternal_openshift_io_fake_cluster_install_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
