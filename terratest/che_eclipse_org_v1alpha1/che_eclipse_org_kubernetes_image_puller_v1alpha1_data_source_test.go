/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package che_eclipse_org_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCheEclipseOrgKubernetesImagePullerV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_che_eclipse_org_kubernetes_image_puller_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
