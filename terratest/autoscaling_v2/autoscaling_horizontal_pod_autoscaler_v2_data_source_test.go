/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package autoscaling_v2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAutoscalingHorizontalPodAutoscalerV2DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_autoscaling_horizontal_pod_autoscaler_v2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}