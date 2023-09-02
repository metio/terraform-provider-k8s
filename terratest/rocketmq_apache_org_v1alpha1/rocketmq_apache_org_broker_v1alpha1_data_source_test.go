/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package rocketmq_apache_org_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRocketmqApacheOrgBrokerV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_rocketmq_apache_org_broker_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
