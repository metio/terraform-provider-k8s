/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */
package generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseOpenAPIv2Files(t *testing.T) {
	openapiv2Schemas := ParseOpenAPIv2Files("../../../schemas/openapi_v2/")

	assert.Greater(t, len(openapiv2Schemas), 0)
}

func TestKubernetesSchema(t *testing.T) {
	openapiv2Schemas := ParseOpenAPIv2Files("../../../schemas/openapi_v2/")
	kubernetesSchema := openapiv2Schemas[0]
	statefulSet := kubernetesSchema["io.k8s.api.apps.v1.StatefulSet"]
	spec := statefulSet.Value.Properties["spec"].Value
	properties := spec.Properties
	volumeClaimTemplates := properties["volumeClaimTemplates"].Value
	templateProps := volumeClaimTemplates.Items.Value.Properties
	templateSpec := templateProps["spec"].Value
	templateSpecProps := templateSpec.Properties

	assert.Equal(t, 11, len(properties), "properties")
	assert.Equal(t, 5, len(templateProps), "templateProps")
	assert.Equal(t, 9, len(templateSpecProps), "templateSpecProps")
}
