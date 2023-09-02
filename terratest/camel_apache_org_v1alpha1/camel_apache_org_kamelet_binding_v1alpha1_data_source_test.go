/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package camel_apache_org_v1alpha1

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCamelApacheOrgKameletBindingV1Alpha1DataSource(t *testing.T) {
	path := "../../examples/data-sources/k8s_camel_apache_org_kamelet_binding_v1alpha1"

	_, err := os.Stat(path)
	assert.Nil(t, err)
}
