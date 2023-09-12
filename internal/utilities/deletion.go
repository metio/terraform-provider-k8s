/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package utilities

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

func MapDeletionPropagation(value string) *v1.DeletionPropagation {
	var propagation v1.DeletionPropagation
	switch strings.ToLower(value) {
	case "orphan":
		propagation = v1.DeletePropagationOrphan
	case "background":
		propagation = v1.DeletePropagationBackground
	case "foreground":
		fallthrough
	default:
		propagation = v1.DeletePropagationForeground
	}
	return &propagation
}
