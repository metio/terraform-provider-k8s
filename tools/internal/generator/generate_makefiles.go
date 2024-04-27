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
	terratestsTemplate := ParseTemplates(fmt.Sprintf("%s/terratests.mk.tmpl", templatePath))
	testsTemplate := ParseTemplates(fmt.Sprintf("%s/tests.mk.tmpl", templatePath))

	projectTargetFile := fmt.Sprintf("%s/project.mk", outputPath)
	terratestsTargetFile := fmt.Sprintf("%s/terratests.mk", outputPath)
	testsTargetFile := fmt.Sprintf("%s/tests.mk", outputPath)
	generateCode(projectTargetFile, projectTemplate, data)
	generateCode(terratestsTargetFile, terratestsTemplate, data)
	generateCode(testsTargetFile, testsTemplate, data)
}
