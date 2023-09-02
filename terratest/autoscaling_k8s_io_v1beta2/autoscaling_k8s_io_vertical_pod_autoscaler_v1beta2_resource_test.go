/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package autoscaling_k8s_io_v1beta2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAutoscalingK8SIoVerticalPodAutoscalerV1Beta2Resource(t *testing.T) {
	path := "../../examples/resources/k8s_autoscaling_k8s_io_vertical_pod_autoscaler_v1beta2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}