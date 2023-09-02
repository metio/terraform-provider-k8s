/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package core_strimzi_io_v1beta2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCoreStrimziIoStrimziPodSetV1Beta2Resource(t *testing.T) {
	path := "../../examples/resources/k8s_core_strimzi_io_strimzi_pod_set_v1beta2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
