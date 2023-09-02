/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package imaging_ingestion_alvearie_org_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestImagingIngestionAlvearieOrgDimseProxyV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_imaging_ingestion_alvearie_org_dimse_proxy_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
