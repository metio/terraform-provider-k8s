/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package hazelcast_com_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHazelcastComWanReplicationV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_hazelcast_com_wan_replication_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
