/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package operator_aquasec_com_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOperatorAquasecComAquaKubeEnforcerV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_operator_aquasec_com_aqua_kube_enforcer_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
