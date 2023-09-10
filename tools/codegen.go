/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package tools

//Generate code for OpenAPIv2 schemas
// no go:generate go run github.com/metio/terraform-provider-k8s/tools/generator --provider-dir .. --schema-dir ../schemas --openapi

//Generate code for CRDv1 schemas
// no go:generate go run github.com/metio/terraform-provider-k8s/tools/generator --provider-dir .. --schema-dir ../schemas --crd

// Generate code for OpenAPIv2 & CRDv1 schemas
//go:generate go run github.com/metio/terraform-provider-k8s/tools/generator --provider-dir .. --schema-dir ../schemas --openapi --crd
