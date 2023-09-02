/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package ripsaw_cloudbulldozer_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRipsawCloudbulldozerIoBenchmarkV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_ripsaw_cloudbulldozer_io_benchmark_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
