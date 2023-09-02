/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package kustomize_toolkit_fluxcd_io_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestKustomizeToolkitFluxcdIoKustomizationV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_kustomize_toolkit_fluxcd_io_kustomization_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
