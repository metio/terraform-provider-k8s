/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package core_openfeature_dev_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCoreOpenfeatureDevFeatureFlagConfigurationV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_core_openfeature_dev_feature_flag_configuration_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
