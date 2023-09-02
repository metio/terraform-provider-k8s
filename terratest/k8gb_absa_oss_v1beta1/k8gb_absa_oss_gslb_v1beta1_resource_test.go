/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package k8gb_absa_oss_v1beta1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestK8GbAbsaOssGslbV1Beta1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_k8gb_absa_oss_gslb_v1beta1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
