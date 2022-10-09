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
	"sort"
	"strconv"
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
	UsedValidators        UsedValidators
}

type UsedValidators struct {
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

	validators := UsedValidators{}

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
		Properties:            Properties2(schema, &validators),
		//Properties:            Properties(schema.Properties, schema.Required, &validators),
		TerraformResourceName: TerraformResourceName(&crd.Spec),
		UsedValidators:        validators,
	}
}

func Properties2(schema *apiextensionsv1.JSONSchemaProps, uv *UsedValidators) []*Property {
	props := make([]*Property, 0)

	for name, prop := range schema.Properties {
		var nestedProperties []*Property
		if prop.Type == "array" && prop.Items.Schema.Type == "object" {
			nestedProperties = Properties2(prop.Items.Schema, uv)
		} else {
			nestedProperties = Properties2(&prop, uv)
		}

		props = append(props, &Property{
			BT:                     "`",
			Name:                   name,
			GoName:                 GoName(name),
			GoType:                 GoType(prop),
			TerraformAttributeName: TerraformAttributeName(name),
			TerraformAttributeType: TerraformAttributeType(prop),
			TerraformValueType:     TerraformValueType(prop),
			Description:            Description(prop.Description),
			Required:               slices.Contains(schema.Required, name),
			Optional:               !slices.Contains(schema.Required, name),
			Computed:               false,
			Properties:             nestedProperties,
			Validators:             Validators(prop, uv),
		})
	}

	sort.SliceStable(props, func(i, j int) bool {
		return props[i].Name < props[j].Name
	})

	if schema.MinProperties != nil && schema.MaxProperties != nil {
		min := *schema.MinProperties
		max := *schema.MaxProperties

		if min == 1 && max == 1 {
			uv.SchemaValidator = true

			for _, outer := range props {
				for _, inner := range props {
					if outer.Name != inner.Name {
						validator := fmt.Sprintf(`schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("%s"))`, inner.TerraformAttributeName)
						outer.Validators = append(outer.Validators, validator)
					}
				}
			}
		}
	} else if schema.MinProperties != nil && schema.MaxProperties == nil {
		min := *schema.MinProperties

		if min == 1 {
			uv.SchemaValidator = true

			for _, outer := range props {
				for _, inner := range props {
					if outer.Name != inner.Name {
						validator := fmt.Sprintf(`schemavalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("%s"))`, inner.TerraformAttributeName)
						outer.Validators = append(outer.Validators, validator)
					}
				}
			}
		} else if min > 1 && min == int64(len(props)) {
			for _, prop := range props {
				prop.Required = true
				prop.Optional = false
			}
		}
	}

	return props
}

//
//func Properties(properties map[string]apiextensionsv1.JSONSchemaProps, required []string, uv *UsedValidators) []Property {
//	props := make([]Property, 0)
//
//	for name, prop := range properties {
//		var nestedProperties []Property
//		if prop.Type == "array" && prop.Items.Schema.Type == "object" {
//			nestedProperties = Properties(prop.Items.Schema.Properties, prop.Items.Schema.Required, uv)
//		} else {
//			nestedProperties = Properties(prop.Properties, prop.Required, uv)
//		}
//
//		props = append(props, Property{
//			BT:                     "`",
//			Name:                   name,
//			GoName:                 GoName(name),
//			GoType:                 GoType(prop),
//			TerraformAttributeName: TerraformAttributeName(name),
//			TerraformAttributeType: TerraformAttributeType(prop),
//			TerraformValueType:     TerraformValueType(prop),
//			Description:            Description(prop.Description),
//			Required:               slices.Contains(required, name),
//			Optional:               !slices.Contains(required, name),
//			Computed:               false,
//			Properties:             nestedProperties,
//			Validators:             Validators(prop, uv),
//		})
//	}
//
//	sort.SliceStable(props, func(i, j int) bool {
//		return props[i].Name < props[j].Name
//	})
//
//	return props
//}

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

func Validators(prop apiextensionsv1.JSONSchemaProps, uv *UsedValidators) []string {
	validators := make([]string, 0)

	if prop.Type == "integer" && prop.Minimum != nil {
		validators = append(validators, fmt.Sprintf("int64validator.AtLeast(%v)", minValue(prop)))
		uv.Int64Validator = true
	}
	if prop.Type == "integer" && prop.Maximum != nil {
		validators = append(validators, fmt.Sprintf("int64validator.AtMost(%v)", maxValue(prop)))
		uv.Int64Validator = true
	}
	if prop.Type == "integer" && len(prop.Enum) > 0 {
		enums := intEnums(prop.Enum)
		validators = append(validators, fmt.Sprintf("int64validator.OneOf(%v)", concatEnums(enums)))
		uv.Int64Validator = true
	}
	if prop.Type == "number" && prop.Minimum != nil {
		validators = append(validators, fmt.Sprintf("float64validator.AtLeast(%v)", minValue(prop)))
		uv.Float64Validator = true
	}
	if prop.Type == "number" && prop.Maximum != nil {
		validators = append(validators, fmt.Sprintf("float64validator.AtMost(%v)", maxValue(prop)))
		uv.Float64Validator = true
	}
	if prop.Type == "number" && len(prop.Enum) > 0 {
		enums := floatEnums(prop.Enum)
		validators = append(validators, fmt.Sprintf("float64validator.OneOf(%v)", concatEnums(enums)))
		uv.Int64Validator = true
	}
	if prop.Type == "string" && prop.Format == "byte" {
		validators = append(validators, "validators.Base64Validator()")
	}
	if prop.Type == "string" && prop.Format == "date-time" {
		validators = append(validators, "validators.DateTime64Validator()")
	}
	if prop.Type == "string" && prop.MinLength != nil {
		validators = append(validators, fmt.Sprintf("stringvalidator.LengthAtLeast(%v)", *prop.MinLength))
		uv.StringValidator = true
	}
	if prop.Type == "string" && prop.MaxLength != nil {
		validators = append(validators, fmt.Sprintf("stringvalidator.LengthAtMost(%v)", *prop.MaxLength))
		uv.StringValidator = true
	}
	if prop.Type == "string" && len(prop.Enum) > 0 {
		enums := stringEnums(prop.Enum)
		validators = append(validators, fmt.Sprintf("stringvalidator.OneOf(%s)", concatEnums(enums)))
		uv.StringValidator = true
	}
	if prop.Type == "string" && prop.Pattern != "" {
		validators = append(validators, fmt.Sprintf(`stringvalidator.RegexMatches(regexp.MustCompile(%c%s%c), "")`, '`', prop.Pattern, '`'))
		uv.StringValidator = true
		uv.Regex = true
	}

	return validators
}

func stringEnums(enums []apiextensionsv1.JSON) []string {
	var values []string

	for _, val := range enums {
		if str := string(val.Raw); str != "" {
			values = append(values, str)
		}
	}

	return values
}

func intEnums(enums []apiextensionsv1.JSON) []int64 {
	var values []int64

	for _, val := range enums {
		if str := string(val.Raw); str != "" {
			i, err := strconv.ParseInt(str, 10, 64)
			if err == nil {
				values = append(values, i)
			}
		}
	}

	return values
}

func floatEnums(enums []apiextensionsv1.JSON) []float64 {
	var values []float64

	for _, val := range enums {
		if str := string(val.Raw); str != "" {
			i, err := strconv.ParseFloat(str, 64)
			if err == nil {
				values = append(values, i)
			}
		}
	}

	return values
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

func maxValue(prop apiextensionsv1.JSONSchemaProps) float64 {
	max := *prop.Maximum
	if prop.ExclusiveMaximum {
		max = max - 1
	}
	return max
}

func minValue(prop apiextensionsv1.JSONSchemaProps) float64 {
	min := *prop.Minimum
	if prop.ExclusiveMinimum {
		min = min + 1
	}
	return min
}
