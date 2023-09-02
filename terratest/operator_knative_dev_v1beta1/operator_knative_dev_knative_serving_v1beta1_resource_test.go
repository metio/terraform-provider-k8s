/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package operator_knative_dev_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOperatorKnativeDevKnativeServingV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_operator_knative_dev_knative_serving_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
