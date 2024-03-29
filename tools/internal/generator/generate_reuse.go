/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"fmt"
	"github.com/metio/terraform-provider-k8s/tools/internal/fetcher"
)

func GenerateReuseFiles(templatePath string, outputPath string, openapi []fetcher.UpstreamSource) {
	dep5Template := ParseTemplates(fmt.Sprintf("%s/dep5.tmpl", templatePath))

	data := dep5TemplateData{
		OpenAPI: openapi,
	}

	dep5TargetFile := fmt.Sprintf("%s/dep5", outputPath)
	generateCode(dep5TargetFile, dep5Template, data)
}

type dep5TemplateData struct {
	OpenAPI []fetcher.UpstreamSource
	CRD     []fetcher.UpstreamSource
}
