/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package mq_services_k8s_aws_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMqServicesK8SAwsBrokerV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_mq_services_k8s_aws_broker_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
