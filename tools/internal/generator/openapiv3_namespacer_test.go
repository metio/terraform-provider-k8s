/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isNamespacedObject_True(t *testing.T) {
	for _, tt := range namespacedObjects {
		t.Run(tt, func(t *testing.T) {
			assert.Truef(t, isNamespacedObject(tt), "isNamespacedObject(%v)", tt)
		})
	}
}

func Test_isNamespacedObject_False(t *testing.T) {
	tests := []string{
		"io.k8s.api.rbac.v1.ClusterRole",
		"io.k8s.api.rbac.v1.ClusterRoleBinding",
	}
	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			assert.Falsef(t, isNamespacedObject(tt), "isNamespacedObject(%v)", tt)
		})
	}
}
