/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package fetcher

var OpenAPIv2Sources = []UpstreamSource{
	{
		ProjectName: "kubernetes/kubernetes",
		License:     apacheV2,
		URLs: []string{
			"https://github.com/kubernetes/kubernetes/blob/master/api/openapi-spec/swagger.json",
		},
	},
}
