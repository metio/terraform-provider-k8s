//go:build generators

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"regexp"
	"strings"
	"unicode"
)

type TemplateData struct {
	BT      string
	Package string

	// ResourceFile is the file name of the resource
	ResourceFile string

	// ResourceTestFile is the file name of the test for the resource
	ResourceTestFile string

	// ResourceWorkflowFile is the file name of the GitHub Actions workflow for the resource
	ResourceWorkflowFile string

	// ManifestFile is the file name of the data source
	DataSourceFile string

	// DataSourceTestFile is the file name of the test for the data source
	DataSourceTestFile string

	// DataSourceWorkflowFile is the file name of the GitHub Actions workflow for the data source
	DataSourceWorkflowFile string

	// ManifestFile is the file name of the manifest data source
	ManifestFile string

	// ManifestTestFile is the file name of the test for the manifest data source
	ManifestTestFile string

	// ManifestWorkflowFile is the file name of the GitHub Actions workflow for the manifest data source
	ManifestWorkflowFile string

	Group       string
	Version     string
	Kind        string
	Description string
	Namespaced  bool

	// ResourceTypeStruct is the CamelCase version of the name used by resources for the Terraform type
	ResourceTypeStruct string

	// ResourceDataStruct is the CamelCase version of the name used by resources for the data struct
	ResourceDataStruct string

	// ResourceTypeStruct is the CamelCase version of the name used by tests for a resource
	ResourceTypeTest string

	// DataSourceTypeStruct is the CamelCase version of the name used by data sources for the Terraform type
	DataSourceTypeStruct string

	// DataSourceDataStruct is the CamelCase version of the name used by data sources for the data struct
	DataSourceDataStruct string

	// DataSourceTypeTest is the CamelCase version of the name used by tests for a data source
	DataSourceTypeTest string

	// ManifestTypeStruct is the CamelCase version of the name used by manifest data sources for the Terraform type
	ManifestTypeStruct string

	// ManifestDataStruct is the CamelCase version of the name used by manifest data sources for data struct
	ManifestDataStruct string

	// ManifestTypeTest is the CamelCase version of the name used by tests for a manifest data source
	ManifestTypeTest string

	TerraformModelType string

	// ResourceTypeName is the snake_case version of the name as used by Terraform without the provider prefix
	ResourceTypeName string

	// FullResourceTypeName is the snake_case version of the name as used by Terraform with the provider prefix
	FullResourceTypeName string

	// DataSourceTypeName is the snake_case version of the name as used by Terraform without the provider prefix
	DataSourceTypeName string

	// FullDataSourceTypeName is the snake_case version of the name as used by Terraform with the provider prefix
	FullDataSourceTypeName string

	// ManifestTypeName is the snake_case version of the name as used by Terraform without the provider prefix
	ManifestTypeName string

	// FullManifestTypeName is the snake_case version of the name as used by Terraform with the provider prefix
	FullManifestTypeName string

	Properties        []*Property
	AdditionalImports AdditionalImports
}

type AdditionalImports struct {
	Int64Validator   bool
	Float64Validator bool
	StringValidator  bool
	Regex            bool
	SchemaValidator  bool
}

type Property struct {
	BT                     string
	Name                   string
	GoName                 string
	GoType                 string
	TerraformAttributeName string
	TerraformAttributeType string
	TerraformElementType   string
	TerraformValueType     string
	Description            string
	Required               bool
	Optional               bool
	Computed               bool
	Properties             []*Property
	ValidatorsType         string
	Validators             []string
}

func propertyPath(path string, attributeName string) string {
	if len(path) == 0 {
		return attributeName
	}
	return fmt.Sprintf("%s.%s", path, attributeName)
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
var matchBackticks = regexp.MustCompile(`\x60`)
var matchDoubleQuotes = regexp.MustCompile("\"")
var matchNewlines = regexp.MustCompile("\n")
var matchBackslashes = regexp.MustCompile(`\\`)
var matchDashes = regexp.MustCompile("-")
var matchDots = regexp.MustCompile(`\.`)
var matchSlashes = regexp.MustCompile("/")
var matchColons = regexp.MustCompile(":")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	snake = matchDashes.ReplaceAllString(snake, "_")
	snake = matchDots.ReplaceAllString(snake, "_")
	snake = matchSlashes.ReplaceAllString(snake, "_")
	snake = matchColons.ReplaceAllString(snake, "_")
	return strings.ToLower(snake)
}

func upperCaseFirstLetter(s string) string {
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func resourceFile(group string, kind string, version string) string {
	return goFilename(group, kind, version, "resource")
}

func resourceTestFile(group string, kind string, version string) string {
	return goFilename(group, kind, version, "resource_test")
}

func resourceWorkflowFile(group string, kind string, version string) string {
	return githubActionTerratestFilename(group, kind, version, "resource")
}

func dataSourceFile(group string, kind string, version string) string {
	return goFilename(group, kind, version, "data_source")
}

func dataSourceTestFile(group string, kind string, version string) string {
	return goFilename(group, kind, version, "data_source_test")
}

func dataSourceWorkflowFile(group string, kind string, version string) string {
	return githubActionTerratestFilename(group, kind, version, "data_source")
}

func manifestFile(group string, kind string, version string) string {
	return goFilename(group, kind, version, "manifest")
}

func manifestTestFile(group string, kind string, version string) string {
	return goFilename(group, kind, version, "manifest_test")
}

func manifestWorkflowFile(group string, kind string, version string) string {
	return githubActionTerratestFilename(group, kind, version, "manifest")
}

func goFilename(group string, kind string, version string, suffix string) string {
	if len(group) > 0 {
		return fmt.Sprintf("%s_%s_%s_%s.go", toSnakeCase(group), toSnakeCase(kind), version, suffix)
	}
	return fmt.Sprintf("%s_%s_%s.go", toSnakeCase(kind), version, suffix)
}

func githubActionTerratestFilename(group string, kind string, version string, suffix string) string {
	if len(group) > 0 {
		return fmt.Sprintf("terratest-%s_%s_%s_%s.yml", toSnakeCase(group), toSnakeCase(kind), version, suffix)
	}
	return fmt.Sprintf("terratest-%s_%s_%s.yml", toSnakeCase(kind), version, suffix)
}

func resourceTypeName(group string, kind string, version string) string {
	if len(group) > 0 {
		return fmt.Sprintf("%s_%s_%s", toSnakeCase(group), toSnakeCase(kind), version)
	}
	return fmt.Sprintf("%s_%s", toSnakeCase(kind), version)
}

func resourceTypeStruct(group string, kind string, version string) string {
	return typeStruct(group, kind, version, "Resource")
}

func resourceDataStruct(group string, kind string, version string) string {
	return typeStruct(group, kind, version, "ResourceData")
}

func dataSourceTypeStruct(group string, kind string, version string) string {
	return typeStruct(group, kind, version, "DataSource")
}

func dataSourceDataStruct(group string, kind string, version string) string {
	return typeStruct(group, kind, version, "DataSourceData")
}

func manifestTypeStruct(group string, kind string, version string) string {
	return typeStruct(group, kind, version, "Manifest")
}

func manifestDataStruct(group string, kind string, version string) string {
	return typeStruct(group, kind, version, "ManifestData")
}

func typeStruct(group string, kind string, version string, suffix string) string {
	if len(group) > 0 {
		return strcase.ToCamel(fmt.Sprintf("%s_%s_%s_%s", goName(group), kind, version, suffix))
	}
	return strcase.ToCamel(fmt.Sprintf("%s_%s_%s", kind, version, suffix))
}

func resourceTypeTest(group string, kind string, version string) string {
	return typeTest(group, kind, version, "Resource")
}

func dataSourceTypeTest(group string, kind string, version string) string {
	return typeTest(group, kind, version, "DataSource")
}

func manifestTypeTest(group string, kind string, version string) string {
	return typeTest(group, kind, version, "Manifest")
}

func typeTest(group string, kind string, version string, suffix string) string {
	if len(group) > 0 {
		return strcase.ToCamel(fmt.Sprintf("Test%s_%s_%s_%s", goName(group), kind, version, suffix))
	}
	return strcase.ToCamel(fmt.Sprintf("Test%s_%s_%s", kind, version, suffix))
}

func terraformModelType(group string, kind string, version string) string {
	if len(group) > 0 {
		return strcase.ToCamel(fmt.Sprintf("%s_%s_%s_TerraformModel", goName(group), kind, version))
	}
	return strcase.ToCamel(fmt.Sprintf("%s_%s_TerraformModel", kind, version))
}

func goPackageName(group string, version string) string {
	if len(group) > 0 {
		return toSnakeCase(fmt.Sprintf("%s_%s", goName(group), version))
	}
	return toSnakeCase(fmt.Sprintf("%s_%s", "core", version))
}

func terraformAttributeName(str string, rootPath bool) string {
	clean := str
	if rootPath && clean == "provisioner" {
		clean = "k8s_provisioner"
	}
	if strings.HasPrefix(clean, "-") {
		clean = strings.Replace(clean, "-", "", 1)
	}
	if strings.HasPrefix(clean, "3") {
		clean = strings.Replace(clean, "3", "Three", 1)
	}
	if strings.HasPrefix(clean, "$") {
		clean = strings.Replace(clean, "$", "Dollar", 1)
	}
	clean = strings.ReplaceAll(clean, "URL", "Url")
	clean = toSnakeCase(clean)
	return clean
}

func goName(s string) string {
	clean := upperCaseFirstLetter(s)
	if strings.HasPrefix(clean, "3") {
		clean = strings.Replace(clean, "3", "Three", 1)
	}
	if strings.HasPrefix(clean, "$") {
		clean = strings.Replace(clean, "$", "Dollar", 1)
	}
	clean = matchDashes.ReplaceAllString(clean, "_")
	clean = matchDots.ReplaceAllString(clean, "_")
	clean = matchSlashes.ReplaceAllString(clean, "_")
	clean = matchColons.ReplaceAllString(clean, "_")
	return clean
}

func description(description string) string {
	clean := matchBackticks.ReplaceAllString(description, "'")
	clean = matchDoubleQuotes.ReplaceAllString(clean, "'")
	clean = matchNewlines.ReplaceAllString(clean, "")
	clean = matchBackslashes.ReplaceAllString(clean, "")
	return clean
}

func concatEnums[T any](enums []T) string {
	output := ""

	for index, value := range enums {
		if index < len(enums) {
			output = output + fmt.Sprintf("%v, ", value)
		} else {
			output = output + fmt.Sprintf("%v", value)
		}
	}

	return output
}
