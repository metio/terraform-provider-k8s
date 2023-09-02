/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package grafana_integreatly_org_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGrafanaIntegreatlyOrgGrafanaDatasourceV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_grafana_integreatly_org_grafana_datasource_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}