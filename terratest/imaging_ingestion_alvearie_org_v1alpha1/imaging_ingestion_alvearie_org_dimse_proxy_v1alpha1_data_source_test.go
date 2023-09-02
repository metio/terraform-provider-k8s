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

func TestImagingIngestionAlvearieOrgDimseProxyV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_imaging_ingestion_alvearie_org_dimse_proxy_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
