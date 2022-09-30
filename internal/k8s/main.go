//go:build k8s

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package k8s

import "flag"

func main() {
	var input string
	var output string
	flag.StringVar(&input, "input", "", "")
	flag.StringVar(&output, "output", "", "")
	flag.Parse()
}
