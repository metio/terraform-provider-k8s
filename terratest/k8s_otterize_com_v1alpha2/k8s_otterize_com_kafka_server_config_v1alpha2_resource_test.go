/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package k8s_otterize_com_v1alpha2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestK8SOtterizeComKafkaServerConfigV1Alpha2Resource(t *testing.T) {
	path := "../../examples/resources/k8s_k8s_otterize_com_kafka_server_config_v1alpha2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
