/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package app_redislabs_com_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAppRedislabsComRedisEnterpriseDatabaseV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_app_redislabs_com_redis_enterprise_database_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
