//go:build fetcher

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import (
	"fmt"
	"log"
)

func downloadOpenAPIv2(targetDirectory string) {
	for group, url := range openAPIv2Sources {
		targetFile := fmt.Sprintf("%s/%s/swagger.json", targetDirectory, group)
		rawUrl := githubRawUrl(url)
		rawUrl = gitlabRawUrl(rawUrl)
		err := downloadFile(targetFile, rawUrl)
		if err != nil {
			log.Printf("cannot handle [%s] because of: %s", url, err)
			continue
		}
	}
}
