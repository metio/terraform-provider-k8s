/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package cache_kubedl_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCacheKubedlIoCacheBackendV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_cache_kubedl_io_cache_backend_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
