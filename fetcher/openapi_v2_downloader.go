//go:build fetcher

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import (
	"fmt"
	"log"
	"strings"
)

func downloadOpenAPIv2(targetDirectory string, filter string) {
	for _, source := range openAPIv2Sources {
		for _, url := range source.URLs {
			if strings.Contains(url, filter) || filter == "" {
				log.Printf("downloading [%s]", url)
				targetFile := fmt.Sprintf("%s/%s/swagger.json", targetDirectory, source.ProjectName)
				err := downloadFile(targetFile, url)
				if err != nil {
					log.Printf("cannot handle [%s] because of: %s", url, err)
					continue
				}
			}
		}
	}
}
