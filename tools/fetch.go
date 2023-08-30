/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package tools

// Download OpenAPIv2 schemas
//go:generate go run github.com/metio/terraform-provider-k8s/tools/fetcher --schema-dir ../schemas --openapi

// Download CRDv1 schemas
//go:generate go run github.com/metio/terraform-provider-k8s/tools/fetcher --schema-dir ../schemas --crd
