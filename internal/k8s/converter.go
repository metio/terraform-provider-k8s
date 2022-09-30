//go:build k8s

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package k8s

import (
	"fmt"
	"github.com/iancoleman/strcase"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/utils/strings/slices"
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
	Properties            []Property
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
	Properties             []Property
}

func ConvertToTemplateData(crds []*apiextensionsv1.CustomResourceDefinition, pkg string) []*TemplateData {
	data := make([]*TemplateData, 0)
	for _, crd := range crds {
		if templateData := AsTemplateData(crd, pkg); templateData != nil {
			data = append(data, templateData)
		}
	}
	return data
}

func AsTemplateData(crd *apiextensionsv1.CustomResourceDefinition, pkg string) *TemplateData {
	version := crd.Spec.Versions[0]
	schema := version.Schema.OpenAPIV3Schema
	// remove manually managed or otherwise ignored properties
	delete(schema.Properties, "metadata")
	delete(schema.Properties, "status")
	delete(schema.Properties, "apiVersion")
	delete(schema.Properties, "kind")

	if len(schema.Properties) == 0 {
		return nil
	}

	return &TemplateData{
		BT:                    "`",
		Package:               pkg,
		File:                  File(&crd.Spec),
		Group:                 crd.Spec.Group,
		Version:               version.Name,
		Kind:                  crd.Spec.Names.Kind,
		Namespaced:            crd.Spec.Scope == apiextensionsv1.NamespaceScoped,
		Description:           Description(schema.Description),
		TerraformResourceType: TerraformResourceType(&crd.Spec, &version),
		TerraformModelType:    TerraformModelType(&crd.Spec, &version),
		GoModelType:           GoModelType(&crd.Spec, &version),
		TerraformResourceName: TerraformResourceName(&crd.Spec),
		Properties:            Properties(schema.Properties, schema.Required),
	}
}

func Properties(properties map[string]apiextensionsv1.JSONSchemaProps, required []string) []Property {
	props := make([]Property, 0)

	for name, prop := range properties {
		var nestedProperties []Property
		if prop.Type == "array" && prop.Items.Schema.Type == "object" {
			nestedProperties = Properties(prop.Items.Schema.Properties, prop.Items.Schema.Required)
		} else {
			nestedProperties = Properties(prop.Properties, prop.Required)
		}

		props = append(props, Property{
			BT:                     "`",
			Name:                   name,
			GoName:                 GoName(name),
			GoType:                 GoType(prop),
			TerraformAttributeName: TerraformAttributeName(name),
			TerraformAttributeType: TerraformAttributeType(prop),
			TerraformValueType:     TerraformValueType(prop),
			Description:            Description(prop.Description),
			Required:               slices.Contains(required, name),
			Optional:               !slices.Contains(required, name),
			Computed:               false,
			Properties:             nestedProperties,
		})
	}

	return props
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
var matchBackticks = regexp.MustCompile("\\x60")
var matchDoubleQuotes = regexp.MustCompile("\"")
var matchNewlines = regexp.MustCompile("\n")
var matchBackslashes = regexp.MustCompile("\\\\")
var matchDashes = regexp.MustCompile("-")
var matchDots = regexp.MustCompile("\\.")
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

func File(crdSpec *apiextensionsv1.CustomResourceDefinitionSpec) string {
	return fmt.Sprintf("resource_%s_%s_%s.go", toSnakeCase(crdSpec.Group), toSnakeCase(crdSpec.Names.Kind), crdSpec.Versions[0].Name)
}

func TerraformResourceName(crdSpec *apiextensionsv1.CustomResourceDefinitionSpec) string {
	return fmt.Sprintf("%s_%s_%s", toSnakeCase(crdSpec.Group), toSnakeCase(crdSpec.Names.Kind), crdSpec.Versions[0].Name)
}

func TerraformResourceType(crdSpec *apiextensionsv1.CustomResourceDefinitionSpec, version *apiextensionsv1.CustomResourceDefinitionVersion) string {
	return strcase.ToCamel(fmt.Sprintf("%s_%s_%s_Resource", GoName(crdSpec.Group), crdSpec.Names.Kind, version.Name))
}

func TerraformModelType(crdSpec *apiextensionsv1.CustomResourceDefinitionSpec, version *apiextensionsv1.CustomResourceDefinitionVersion) string {
	return strcase.ToCamel(fmt.Sprintf("%s_%s_%s_TerraformModel", GoName(crdSpec.Group), crdSpec.Names.Kind, version.Name))
}

func GoModelType(crdSpec *apiextensionsv1.CustomResourceDefinitionSpec, version *apiextensionsv1.CustomResourceDefinitionVersion) string {
	return strcase.ToCamel(fmt.Sprintf("%s_%s_%s_GoModel", GoName(crdSpec.Group), crdSpec.Names.Kind, version.Name))
}

func TerraformAttributeName(str string) string {
	clean := matchDashes.ReplaceAllString(str, "__")
	clean = toSnakeCase(clean)
	return clean
}

func GoName(s string) string {
	clean := upperCaseFirstLetter(s)
	clean = matchDashes.ReplaceAllString(clean, "_")
	clean = matchDots.ReplaceAllString(clean, "_")
	clean = matchSlashes.ReplaceAllString(clean, "_")
	return clean
}

func Description(description string) string {
	clean := matchBackticks.ReplaceAllString(description, "'")
	clean = matchDoubleQuotes.ReplaceAllString(clean, "'")
	clean = matchNewlines.ReplaceAllString(clean, "")
	clean = matchBackslashes.ReplaceAllString(clean, "")
	return clean
}

func TerraformAttributeType(prop apiextensionsv1.JSONSchemaProps) string {
	if prop.XIntOrString {
		return "types.StringType"
	}
	if prop.XPreserveUnknownFields != nil && *prop.XPreserveUnknownFields {
		if len(prop.Properties) > 0 {
			return "types.ObjectType"
		}
		return "types.MapType{ElemType: types.StringType}"
	}
	if prop.Type == "object" && prop.AdditionalProperties != nil && prop.AdditionalProperties.Schema.Type == "string" {
		return "types.MapType{ElemType: types.StringType}"
	}
	if prop.Type == "object" && prop.AdditionalProperties != nil && prop.AdditionalProperties.Schema.Type == "object" {
		if prop.AdditionalProperties.Schema.AdditionalProperties != nil && prop.AdditionalProperties.Schema.AdditionalProperties.Schema.Type == "string" {
			return "types.MapType{ElemType: types.MapType{ElemType: types.StringType}}"
		}
	}
	if prop.Type == "object" && prop.AdditionalProperties != nil && prop.AdditionalProperties.Schema.Type == "array" {
		if prop.AdditionalProperties.Schema.Items.Schema.Type == "string" {
			return "types.MapType{ElemType: types.ListType{ElemType: types.StringType}}"
		}
	}
	if prop.Type == "array" && prop.Items.Schema.Type == "object" {
		if prop.Items.Schema.XPreserveUnknownFields != nil && *prop.Items.Schema.XPreserveUnknownFields {
			return "types.ListType{ElemType: types.MapType{ElemType: types.StringType}}"
		}
		return "types.ListType{ElemType: types.ObjectType}"
	}
	switch prop.Type {
	case "boolean":
		return "types.BoolType"
	case "string":
		return "types.StringType"
	case "integer":
		return "types.Int64Type"
	case "number":
		return "types.NumberType"
	case "array":
		return "types.ListType{ElemType: types.StringType}"
	case "object":
		if len(prop.Properties) > 0 {
			return "types.ObjectType"
		}
		return "types.MapType{ElemType: types.StringType}"
	}
	return "UNKNOWN"
}

func GoType(prop apiextensionsv1.JSONSchemaProps) string {
	if prop.XIntOrString {
		return "string"
	}
	if prop.XPreserveUnknownFields != nil && *prop.XPreserveUnknownFields {
		if len(prop.Properties) > 0 {
			return "struct"
		}
		return "map[string]string"
	}
	if prop.Type == "object" && prop.AdditionalProperties != nil && prop.AdditionalProperties.Schema.Type == "string" {
		return "map[string]string"
	}
	if prop.Type == "object" && prop.AdditionalProperties != nil && prop.AdditionalProperties.Schema.Type == "object" {
		if prop.AdditionalProperties.Schema.AdditionalProperties != nil && prop.AdditionalProperties.Schema.AdditionalProperties.Schema.Type == "string" {
			return "map[string]map[string]string"
		}
	}
	if prop.Type == "object" && prop.AdditionalProperties != nil && prop.AdditionalProperties.Schema.Type == "array" {
		if prop.AdditionalProperties.Schema.Items.Schema.Type == "string" {
			return "map[string][]string"
		}
	}
	if prop.Type == "array" && prop.Items.Schema.Type == "object" {
		if prop.Items.Schema.XPreserveUnknownFields != nil && *prop.Items.Schema.XPreserveUnknownFields {
			return "[]map[string]string"
		}
		return "[]struct"
	}
	switch prop.Type {
	case "boolean":
		return "bool"
	case "string":
		return "string"
	case "integer":
		return "int64"
	case "number":
		return "float64"
	case "array":
		return "[]string"
	case "object":
		if len(prop.Properties) > 0 {
			return "struct"
		}
		return "map[string]string"
	}
	return "UNKNOWN"
}

func TerraformValueType(prop apiextensionsv1.JSONSchemaProps) string {
	if prop.XIntOrString {
		return "types.String"
	}
	if prop.XPreserveUnknownFields != nil && *prop.XPreserveUnknownFields {
		if len(prop.Properties) > 0 {
			return "types.Object"
		}
		return "types.Map"
	}
	if prop.Type == "object" && prop.AdditionalProperties != nil && prop.AdditionalProperties.Schema.Type == "string" {
		return "types.Map"
	}
	switch prop.Type {
	case "boolean":
		return "types.Bool"
	case "string":
		return "types.String"
	case "integer":
		return "types.Int64"
	case "number":
		return "types.Number"
	case "array":
		return "types.List"
	case "object":
		if len(prop.Properties) > 0 {
			return "types.Object"
		}
		return "types.Map"
	}
	return "UNKNOWN"
}
