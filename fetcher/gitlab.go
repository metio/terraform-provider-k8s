//go:build fetcher

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import "strings"

func gitlabRawUrl(url string) string {
	if !strings.HasPrefix(url, "https://gitlab.com") {
		return url
	}
	raw := strings.Replace(url, "/blob/", "/raw/", 1)
	return raw
}
