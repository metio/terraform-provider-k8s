//go:build fetcher

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

var openAPIv2Sources = map[string]string{
	"io.kubernetes": "https://github.com/kubernetes/kubernetes/blob/master/api/openapi-spec/swagger.json",
}
