//go:build fetcher

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

var openAPIv2Sources = []UpstreamSource{
	{
		ProjectName: "kubernetes/kubernetes",
		License:     apacheV2,
		URLs: []string{
			"https://github.com/kubernetes/kubernetes/blob/master/api/openapi-spec/swagger.json",
		},
	},
}
