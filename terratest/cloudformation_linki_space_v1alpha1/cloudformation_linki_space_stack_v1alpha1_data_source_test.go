/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package cloudformation_linki_space_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCloudformationLinkiSpaceStackV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_cloudformation_linki_space_stack_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
