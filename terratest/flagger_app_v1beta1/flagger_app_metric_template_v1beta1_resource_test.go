/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package flagger_app_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFlaggerAppMetricTemplateV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_flagger_app_metric_template_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}