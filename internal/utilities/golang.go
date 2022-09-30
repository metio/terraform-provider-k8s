/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package utilities

func Ptr[T any](v T) *T {
	return &v
}
