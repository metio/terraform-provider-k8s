/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package flow_volcano_sh_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFlowVolcanoShJobTemplateV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_flow_volcano_sh_job_template_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
