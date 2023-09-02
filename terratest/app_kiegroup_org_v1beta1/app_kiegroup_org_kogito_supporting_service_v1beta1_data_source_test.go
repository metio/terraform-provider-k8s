/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package app_kiegroup_org_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAppKiegroupOrgKogitoSupportingServiceV1Beta1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_app_kiegroup_org_kogito_supporting_service_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
