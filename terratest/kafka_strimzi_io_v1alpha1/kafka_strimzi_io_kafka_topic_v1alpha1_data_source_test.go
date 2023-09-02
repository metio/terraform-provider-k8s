/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package kafka_strimzi_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestKafkaStrimziIoKafkaTopicV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_kafka_strimzi_io_kafka_topic_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
