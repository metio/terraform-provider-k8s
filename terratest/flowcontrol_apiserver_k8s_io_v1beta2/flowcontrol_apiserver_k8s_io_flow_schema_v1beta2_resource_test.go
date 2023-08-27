/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package flowcontrol_apiserver_k8s_io_v1beta2

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFlowcontrolApiserverK8SIoFlowSchemaV1Beta2Resource(t *testing.T) {
	path := "../../examples/resources/k8s_flowcontrol_apiserver_k8s_io_flow_schema_v1beta2"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
