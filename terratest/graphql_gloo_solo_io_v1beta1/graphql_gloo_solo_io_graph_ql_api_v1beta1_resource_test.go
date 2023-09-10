/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package graphql_gloo_solo_io_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGraphqlGlooSoloIoGraphQlapiV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_graphql_gloo_solo_io_graph_ql_api_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
