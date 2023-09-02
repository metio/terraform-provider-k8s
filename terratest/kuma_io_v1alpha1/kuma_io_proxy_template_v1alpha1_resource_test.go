/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package kuma_io_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestKumaIoProxyTemplateV1Alpha1Resource(t *testing.T) {
	path := "../../examples/resources/k8s_kuma_io_proxy_template_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}