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

func TestGrafanaIntegreatlyOrgGrafanaDashboardV1Beta1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_grafana_integreatly_org_grafana_dashboard_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
