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
	BT                    string
	Package               string
	File                  string
	Group                 string
	Version               string
	Kind                  string
	Description           string
	Namespaced            bool
	TerraformResourceType string
	TerraformModelType    string
	GoModelType           string
	TerraformResourceName string
	Properties            []*Property
	AdditionalImports     AdditionalImports
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
	TerraformValueType     string
	Description            string
	Required               bool
	Optional               bool
	Computed               bool
	Properties             []*Property
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

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	snake = matchDashes.ReplaceAllString(snake, "_")
	snake = matchDots.ReplaceAllString(snake, "_")
	snake = matchSlashes.ReplaceAllString(snake, "_")
	return strings.ToLower(snake)
}

func upperCaseFirstLetter(s string) string {
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func terraformResourceFile(group string, kind string, version string) string {
	if len(group) > 0 {
		return fmt.Sprintf("resource_%s_%s_%s.go", toSnakeCase(group), toSnakeCase(kind), version)
	}
	return fmt.Sprintf("resource_%s_%s.go", toSnakeCase(kind), version)
}

func terraformResourceName(group string, kind string, version string) string {
	if len(group) > 0 {
		return fmt.Sprintf("%s_%s_%s", toSnakeCase(group), toSnakeCase(kind), version)
	}
	return fmt.Sprintf("%s_%s", toSnakeCase(kind), version)
}

func terraformResourceType(group string, kind string, version string) string {
	if len(group) > 0 {
		return strcase.ToCamel(fmt.Sprintf("%s_%s_%s_Resource", goName(group), kind, version))
	}
	return strcase.ToCamel(fmt.Sprintf("%s_%s_Resource", kind, version))
}

func terraformModelType(group string, kind string, version string) string {
	if len(group) > 0 {
		return strcase.ToCamel(fmt.Sprintf("%s_%s_%s_TerraformModel", goName(group), kind, version))
	}
	return strcase.ToCamel(fmt.Sprintf("%s_%s_TerraformModel", kind, version))
}

func goModelType(group string, kind string, version string) string {
	if len(group) > 0 {
		return strcase.ToCamel(fmt.Sprintf("%s_%s_%s_GoModel", goName(group), kind, version))
	}
	return strcase.ToCamel(fmt.Sprintf("%s_%s_GoModel", kind, version))
}

func terraformAttributeName(str string) string {
	clean := str
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
