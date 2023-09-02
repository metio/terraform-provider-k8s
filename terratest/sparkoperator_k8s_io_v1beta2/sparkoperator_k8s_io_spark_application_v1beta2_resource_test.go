/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package sparkoperator_k8s_io_v1beta2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSparkoperatorK8SIoSparkApplicationV1Beta2Resource(t *testing.T) {
	path := "../../examples/resources/k8s_sparkoperator_k8s_io_spark_application_v1beta2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
