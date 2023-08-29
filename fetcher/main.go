//go:build fetcher

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var filter string
	var fetchCRDv1 bool
	var fetchOpenAPIv2 bool
	flag.StringVar(&filter, "filter", "", "")
	flag.BoolVar(&fetchCRDv1, "crd", true, "")
	flag.BoolVar(&fetchOpenAPIv2, "openapi", true, "")
	flag.Parse()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if fetchOpenAPIv2 {
		downloadOpenAPIv2(fmt.Sprintf("%s/schemas/openapi_v2", cwd), filter)
	}
	//if fetchCRDv1 {
	//	downloadCRDv1(fmt.Sprintf("%s/schemas/crd_v1", cwd), filter)
	//}
}
