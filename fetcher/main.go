//go:build fetcher

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	downloadOpenAPIv2(fmt.Sprintf("%s/schemas/openapi_v2", cwd))
	downloadCRDv1(fmt.Sprintf("%s/schemas/crd_v1", cwd))
}
