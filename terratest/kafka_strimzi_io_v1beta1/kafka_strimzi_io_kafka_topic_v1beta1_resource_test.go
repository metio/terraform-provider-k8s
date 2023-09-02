/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package kafka_strimzi_io_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestKafkaStrimziIoKafkaTopicV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_kafka_strimzi_io_kafka_topic_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
