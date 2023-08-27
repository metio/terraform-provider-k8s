/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package testutilities

type ConfigurationErrorTestCase struct {
	Configuration string
	ErrorRegex    string
}
