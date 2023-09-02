/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package app_redislabs_com_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAppRedislabsComRedisEnterpriseClusterV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_app_redislabs_com_redis_enterprise_cluster_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
