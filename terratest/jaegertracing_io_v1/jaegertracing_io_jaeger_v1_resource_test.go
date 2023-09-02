/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package jaegertracing_io_v1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestJaegertracingIoJaegerV1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_jaegertracing_io_jaeger_v1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
