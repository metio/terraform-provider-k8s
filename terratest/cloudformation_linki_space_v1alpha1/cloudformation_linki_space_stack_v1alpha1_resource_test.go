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

func TestCloudformationLinkiSpaceStackV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_cloudformation_linki_space_stack_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
