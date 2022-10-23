//go:build fetcher

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import "strings"

func githubRawUrl(url string) string {
	if !strings.HasPrefix(url, "https://github.com") || strings.HasPrefix(url, "https://raw.githubusercontent.com") {
		return url
	}

	var raw string
	if strings.HasPrefix(url, "https://github.com") {
		raw = strings.Replace(url, "github.com", "raw.githubusercontent.com", 1)
	} else if strings.HasPrefix(url, "https://www.github.com") {
		raw = strings.Replace(url, "www.github.com", "raw.githubusercontent.com", 1)
	}
	raw = strings.Replace(raw, "/blob", "", 1)

	return raw
}
