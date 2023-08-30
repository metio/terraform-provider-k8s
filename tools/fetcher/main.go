/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import (
	"flag"
	"fmt"
	"github.com/metio/terraform-provider-k8s/tools/internal/fetcher"
	"log"
)

func main() {
	schemaDir := flag.String("schema-dir", "", "relative or absolute path to the root directory for schemas")
	filter := flag.String("filter", "", "Part of an URL to use as filter. Only matching URLs will be downloaded")
	fetchOpenAPIv2 := flag.Bool("openapi", false, "Whether to fetch OpenAPIv2 schemas")
	fetchCRDv1 := flag.Bool("crd", false, "Whether to fetch CRDv1 schemas")
	flag.Parse()

	if *schemaDir == "" {
		log.Fatalln("No --schema-dir specified!")
	}

	if *fetchOpenAPIv2 {
		fetcher.DownloadOpenAPIv2(fmt.Sprintf("%s/openapi_v2", *schemaDir), *filter)
	}
	if *fetchCRDv1 {
		fetcher.DownloadCRDv1(fmt.Sprintf("%s/crd_v1", *schemaDir), *filter)
	}
}
