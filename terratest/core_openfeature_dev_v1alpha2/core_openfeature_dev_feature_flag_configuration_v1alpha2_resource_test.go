/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package core_openfeature_dev_v1alpha2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCoreOpenfeatureDevFeatureFlagConfigurationV1Alpha2Resource(t *testing.T) {
	path := "../../examples/resources/k8s_core_openfeature_dev_feature_flag_configuration_v1alpha2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
