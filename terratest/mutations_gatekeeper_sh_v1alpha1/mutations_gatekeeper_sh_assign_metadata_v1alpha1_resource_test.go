/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package mutations_gatekeeper_sh_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMutationsGatekeeperShAssignMetadataV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_mutations_gatekeeper_sh_assign_metadata_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}