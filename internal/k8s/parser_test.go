//go:build k8s

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package k8s

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestParseAllCustomResourceDefinitions(t *testing.T) {
	crds := ParseAllCustomResourceDefinitions()

	assert.NotEmpty(t, crds)
}

func TestParseOpenApi(t *testing.T) {
	definitions := ParseKubernetesSwagger()

	names := make([]string, 0)
	for name, definition := range definitions {
		if _, ok := definition.Value.ExtensionProps.Extensions["x-kubernetes-group-version-kind"]; ok {
			names = append(names, name)
		}
	}
	sort.SliceStable(names, func(i, j int) bool {
		return names[i] < names[j]
	})
	for _, name := range names {
		fmt.Println(name)
	}
}
