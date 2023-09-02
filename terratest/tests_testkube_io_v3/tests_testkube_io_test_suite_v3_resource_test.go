/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package tests_testkube_io_v3

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTestsTestkubeIoTestSuiteV3Resource(t *testing.T) {
	path := "../../examples/resources/k8s_tests_testkube_io_test_suite_v3"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
