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

func TestImagingIngestionAlvearieOrgDicomStudyBindingV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_imaging_ingestion_alvearie_org_dicom_study_binding_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
