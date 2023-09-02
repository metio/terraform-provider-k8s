/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package secscan_quay_redhat_com_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSecscanQuayRedhatComImageManifestVulnV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_secscan_quay_redhat_com_image_manifest_vuln_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
