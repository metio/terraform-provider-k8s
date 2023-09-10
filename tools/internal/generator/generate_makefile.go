/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"fmt"
)

func GenerateMakefiles(templatePath string, outputPath string, data []*TemplateData) {
	projectTemplate := ParseTemplates(fmt.Sprintf("%s/project.mk.tmpl", templatePath))

	projectTargetFile := fmt.Sprintf("%s/project.mk", outputPath)
	generateCode(projectTargetFile, projectTemplate, data)
}
