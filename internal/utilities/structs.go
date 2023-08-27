/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package utilities

import "k8s.io/client-go/dynamic"

type ResourceData struct {
	Client         dynamic.Interface
	FieldManager   string
	ForceConflicts bool
	Offline        bool
}

type DataSourceData struct {
	Client  dynamic.Interface
	Offline bool
}
