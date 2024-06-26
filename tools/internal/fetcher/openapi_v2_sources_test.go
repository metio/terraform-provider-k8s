/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package fetcher_test

import (
	"github.com/metio/terraform-provider-k8s/tools/internal/fetcher"
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

func TestAlphabeticalOpenAPIs(t *testing.T) {
	var original []string
	var ordered []string

	for _, source := range fetcher.OpenAPIv2Sources {
		original = append(original, strings.ToLower(source.ProjectName))
		ordered = append(ordered, strings.ToLower(source.ProjectName))
	}

	sort.SliceStable(ordered, func(i, j int) bool {
		return ordered[i] < ordered[j]
	})

	assert.Equal(t, ordered, original)
}

func TestOpenAPIsProjectNameFormat(t *testing.T) {
	for _, source := range fetcher.OpenAPIv2Sources {
		assert.True(t, strings.Contains(source.ProjectName, "/"), source.ProjectName)
	}
}
