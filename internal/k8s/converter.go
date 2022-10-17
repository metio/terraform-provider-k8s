//go:build k8s

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package k8s

import (
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
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

func CRDsToTemplateData(crds []*apiextensionsv1.CustomResourceDefinition, pkg string) []*TemplateData {
	data := make([]*TemplateData, 0)
	for _, crd := range crds {
		if templateData := CrdAsTemplateData(crd, pkg); templateData != nil {
			data = append(data, templateData)
		}
	}
	sort.SliceStable(data, func(i, j int) bool {
		return data[i].Group < data[j].Group
	})
	sort.SliceStable(data, func(i, j int) bool {
		return data[i].Version < data[j].Version
	})
	sort.SliceStable(data, func(i, j int) bool {
		return data[i].Kind < data[j].Kind
	})
	return data
}

var allowedDefinitions = []string{
	"io.k8s.api.admissionregistration.v1.MutatingWebhookConfiguration",
	"io.k8s.api.admissionregistration.v1.ValidatingWebhookConfiguration",
	"io.k8s.api.apps.v1.DaemonSet",
	"io.k8s.api.apps.v1.Deployment",
	"io.k8s.api.apps.v1.ReplicaSet",
	"io.k8s.api.apps.v1.StatefulSet",
	"io.k8s.api.autoscaling.v1.HorizontalPodAutoscaler",
	"io.k8s.api.autoscaling.v2.HorizontalPodAutoscaler",
	"io.k8s.api.batch.v1.CronJob",
	"io.k8s.api.batch.v1.Job",
	"io.k8s.api.certificates.v1.CertificateSigningRequest",
	"io.k8s.api.core.v1.ConfigMap",
	"io.k8s.api.core.v1.Endpoints",
	"io.k8s.api.core.v1.LimitRange",
	"io.k8s.api.core.v1.Namespace",
	"io.k8s.api.core.v1.PersistentVolume",
	"io.k8s.api.core.v1.PersistentVolumeClaim",
	"io.k8s.api.core.v1.Pod",
	"io.k8s.api.core.v1.ReplicationController",
	"io.k8s.api.core.v1.Secret",
	"io.k8s.api.core.v1.Service",
	"io.k8s.api.core.v1.ServiceAccount",
	"io.k8s.api.discovery.v1.EndpointSlice",
	"io.k8s.api.events.v1.Event",
	"io.k8s.api.flowcontrol.v1beta2.FlowSchema",
	"io.k8s.api.flowcontrol.v1beta2.PriorityLevelConfiguration",
	"io.k8s.api.flowcontrol.v1beta3.FlowSchema",
	"io.k8s.api.flowcontrol.v1beta3.PriorityLevelConfiguration",
	"io.k8s.api.networking.v1.Ingress",
	"io.k8s.api.networking.v1.IngressClass",
	"io.k8s.api.networking.v1.NetworkPolicy",
	"io.k8s.api.policy.v1.PodDisruptionBudget",
	"io.k8s.api.rbac.v1.ClusterRole",
	"io.k8s.api.rbac.v1.ClusterRoleBinding",
	"io.k8s.api.rbac.v1.Role",
	"io.k8s.api.rbac.v1.RoleBinding",
	"io.k8s.api.scheduling.v1.PriorityClass",
	"io.k8s.api.storage.v1.CSIDriver",
	"io.k8s.api.storage.v1.CSINode",
	"io.k8s.api.storage.v1.StorageClass",
	"io.k8s.api.storage.v1.VolumeAttachment",
}

func OpenApiToTemplateData(definitions map[string]*openapi3.SchemaRef, pkg string) []*TemplateData {
	data := make([]*TemplateData, 0)
	for name, definition := range definitions {
		if slices.Contains(allowedDefinitions, name) {
			if _, ok := definition.Value.ExtensionProps.Extensions["x-kubernetes-group-version-kind"]; ok {
				if templateData := OpenApiAsTemplateData(definition, pkg); templateData != nil {
					data = append(data, templateData)
				}
			}
		}
	}
	sort.SliceStable(data, func(i, j int) bool {
		return data[i].Group < data[j].Group
	})
	sort.SliceStable(data, func(i, j int) bool {
		return data[i].Version < data[j].Version
	})
	sort.SliceStable(data, func(i, j int) bool {
		return data[i].Kind < data[j].Kind
	})
	return data
}

func CrdAsTemplateData(crd *apiextensionsv1.CustomResourceDefinition, pkg string) *TemplateData {
	group := crd.Spec.Group
	version := crd.Spec.Versions[0]
	kind := crd.Spec.Names.Kind

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
		File:                  File(group, kind, version.Name),
		Group:                 group,
		Version:               version.Name,
		Kind:                  kind,
		Namespaced:            crd.Spec.Scope == apiextensionsv1.NamespaceScoped,
		Description:           Description(schema.Description),
		TerraformResourceType: TerraformResourceType(group, kind, version.Name),
		TerraformModelType:    TerraformModelType(group, kind, version.Name),
		GoModelType:           GoModelType(group, kind, version.Name),
		Properties:            CrdProperties(schema, &validators),
		TerraformResourceName: TerraformResourceName(group, kind, version.Name),
		UsedValidators:        validators,
	}
}

type GVK struct {
	Group   string `json:"group"`
	Version string `json:"version"`
	Kind    string `json:"kind"`
}

func OpenApiAsTemplateData(definition *openapi3.SchemaRef, pkg string) *TemplateData {
	var group string
	var version string
	var kind string
	if gvkExt, ok := definition.Value.ExtensionProps.Extensions["x-kubernetes-group-version-kind"]; ok {
		raw := gvkExt.(json.RawMessage)
		var gvks []GVK
		if err := json.Unmarshal(raw, &gvks); err != nil {
			return nil
		}
		if len(gvks) != 1 {
			return nil
		}
		gvk := gvks[0]
		group = gvk.Group
		version = gvk.Version
		kind = gvk.Kind
	}
	schema := definition.Value
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
		File:                  File(group, kind, version),
		Group:                 group,
		Version:               version,
		Kind:                  kind,
		Namespaced:            true,
		Description:           Description(schema.Description),
		TerraformResourceType: TerraformResourceType(group, kind, version),
		TerraformModelType:    TerraformModelType(group, kind, version),
		GoModelType:           GoModelType(group, kind, version),
		Properties:            OpenApiProperties(schema, &validators),
		TerraformResourceName: TerraformResourceName(group, kind, version),
		UsedValidators:        validators,
	}
}

func CrdProperties(schema *apiextensionsv1.JSONSchemaProps, uv *UsedValidators) []*Property {
	props := make([]*Property, 0)

	for name, prop := range schema.Properties {
		var nestedProperties []*Property
		if prop.Type == "array" && prop.Items.Schema.Type == "object" {
			nestedProperties = CrdProperties(prop.Items.Schema, uv)
		} else {
			nestedProperties = CrdProperties(&prop, uv)
		}

		attributeType, valueType, goType := CRDv1Types(prop)
		props = append(props, &Property{
			BT:                     "`",
			Name:                   name,
			GoName:                 GoName(name),
			GoType:                 goType,
			TerraformAttributeName: TerraformAttributeName(name),
			TerraformAttributeType: attributeType,
			TerraformValueType:     valueType,
			Description:            Description(prop.Description),
			Required:               slices.Contains(schema.Required, name),
			Optional:               !slices.Contains(schema.Required, name),
			Computed:               false,
			Properties:             nestedProperties,
			Validators:             CRDv1Validators(prop, uv),
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

func OpenApiProperties(schema *openapi3.Schema, uv *UsedValidators) []*Property {
	props := make([]*Property, 0)

	if schema != nil {
		for name, prop := range schema.Properties {
			if prop.Value != nil {
				var nestedProperties []*Property
				if prop.Value.Type == "array" && prop.Value.Items != nil && prop.Value.Items.Value != nil && prop.Value.Items.Value.Type == "object" {
					nestedProperties = OpenApiProperties(prop.Value.Items.Value, uv)
				} else {
					nestedProperties = OpenApiProperties(prop.Value, uv)
				}

				attributeType, valueType, goType := OpenApiTypes(prop.Value)
				props = append(props, &Property{
					BT:                     "`",
					Name:                   name,
					GoName:                 GoName(name),
					GoType:                 goType,
					TerraformAttributeName: TerraformAttributeName(name),
					TerraformAttributeType: attributeType,
					TerraformValueType:     valueType,
					Description:            Description(prop.Value.Description),
					Required:               slices.Contains(schema.Required, name),
					Optional:               !slices.Contains(schema.Required, name),
					Computed:               false,
					Properties:             nestedProperties,
					Validators:             OpenApiValidators(prop.Value, uv),
				})
			}
		}
	}

	sort.SliceStable(props, func(i, j int) bool {
		return props[i].Name < props[j].Name
	})

	return props
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

func File(group string, kind string, version string) string {
	if len(group) > 0 {
		return fmt.Sprintf("resource_%s_%s_%s.go", toSnakeCase(group), toSnakeCase(kind), version)
	}
	return fmt.Sprintf("resource_%s_%s.go", toSnakeCase(kind), version)
}

func TerraformResourceName(group string, kind string, version string) string {
	if len(group) > 0 {
		return fmt.Sprintf("%s_%s_%s", toSnakeCase(group), toSnakeCase(kind), version)
	}
	return fmt.Sprintf("%s_%s", toSnakeCase(kind), version)
}

func TerraformResourceType(group string, kind string, version string) string {
	if len(group) > 0 {
		return strcase.ToCamel(fmt.Sprintf("%s_%s_%s_Resource", GoName(group), kind, version))
	}
	return strcase.ToCamel(fmt.Sprintf("%s_%s_Resource", kind, version))
}

func TerraformModelType(group string, kind string, version string) string {
	if len(group) > 0 {
		return strcase.ToCamel(fmt.Sprintf("%s_%s_%s_TerraformModel", GoName(group), kind, version))
	}
	return strcase.ToCamel(fmt.Sprintf("%s_%s_TerraformModel", kind, version))
}

func GoModelType(group string, kind string, version string) string {
	if len(group) > 0 {
		return strcase.ToCamel(fmt.Sprintf("%s_%s_%s_GoModel", GoName(group), kind, version))
	}
	return strcase.ToCamel(fmt.Sprintf("%s_%s_GoModel", kind, version))
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

func CRDv1Types(prop apiextensionsv1.JSONSchemaProps) (attributeType string, valueType string, goType string) {
	if prop.XIntOrString {
		attributeType = "types.StringType"
		valueType = "types.String"
		goType = "string"
		return
	}
	if prop.XPreserveUnknownFields != nil && *prop.XPreserveUnknownFields {
		if len(prop.Properties) > 0 {
			attributeType = "types.ObjectType"
			valueType = "types.Object"
			goType = "struct"
			return
		}
		attributeType = "utilities.DynamicType{}"
		valueType = "utilities.Dynamic"
		goType = "utilities.Dynamic"
		return
	}
	if prop.Type == "object" && prop.AdditionalProperties != nil && prop.AdditionalProperties.Schema.Type == "string" {
		attributeType = "types.MapType{ElemType: types.StringType}"
		valueType = "types.Map"
		goType = "map[string]string"
		return
	}
	if prop.Type == "object" && prop.AdditionalProperties != nil && prop.AdditionalProperties.Schema.Type == "object" {
		if prop.AdditionalProperties.Schema.AdditionalProperties != nil && prop.AdditionalProperties.Schema.AdditionalProperties.Schema.Type == "string" {
			attributeType = "types.MapType{ElemType: types.MapType{ElemType: types.StringType}}"
			valueType = "types.Map"
			goType = "map[string]map[string]string"
			return
		}
	}
	if prop.Type == "object" && prop.AdditionalProperties != nil && prop.AdditionalProperties.Schema.Type == "array" {
		if prop.AdditionalProperties.Schema.Items.Schema.Type == "string" {
			attributeType = "types.MapType{ElemType: types.ListType{ElemType: types.StringType}}"
			valueType = "types.Map"
			goType = "map[string][]string"
			return
		}
	}
	if prop.Type == "array" && prop.Items.Schema.Type == "object" {
		if prop.Items.Schema.XPreserveUnknownFields != nil && *prop.Items.Schema.XPreserveUnknownFields {
			attributeType = "types.ListType{ElemType: types.MapType{ElemType: types.StringType}}"
			valueType = "types.List"
			goType = "[]map[string]string"
			return
		}
		attributeType = "types.ListType{ElemType: types.ObjectType}"
		valueType = "types.List"
		goType = "[]struct"
		return
	}
	switch prop.Type {
	case "boolean":
		attributeType = "types.BoolType"
		valueType = "types.Bool"
		goType = "bool"
		return
	case "string":
		attributeType = "types.StringType"
		valueType = "types.String"
		goType = "string"
		return
	case "integer":
		attributeType = "types.Int64Type"
		valueType = "types.Int64"
		goType = "int64"
		return
	case "number":
		if prop.Format == "float" || prop.Format == "double" {
			attributeType = "types.Float64Type"
			valueType = "types.Float64"
			goType = "float64"
			return
		}
		attributeType = "utilities.DynamicNumberType{}"
		valueType = "utilities.DynamicNumber"
		goType = "utilities.DynamicNumber"
		return
	case "array":
		attributeType = "types.ListType{ElemType: types.StringType}"
		valueType = "types.List"
		goType = "[]string"
		return
	case "object":
		if len(prop.Properties) > 0 {
			attributeType = "types.ObjectType"
			valueType = "types.Object"
			goType = "struct"
			return
		}
		attributeType = "types.MapType{ElemType: types.StringType}"
		valueType = "types.Map"
		goType = "map[string]string"
		return
	}
	attributeType = "UNKNOWN"
	valueType = "UNKNOWN"
	goType = "UNKNOWN"
	return
}

func OpenApiTypes(prop *openapi3.Schema) (attributeType string, valueType string, goType string) {
	if _, ok := prop.Extensions["x-kubernetes-int-or-string"]; ok {
		attributeType = "types.StringType"
		valueType = "types.String"
		goType = "string"
		return
	}
	if _, ok := prop.Extensions["x-kubernetes-preserve-unknown-fields"]; ok {
		if len(prop.Properties) > 0 {
			attributeType = "types.ObjectType"
			valueType = "types.Object"
			goType = "struct"
			return
		}
		attributeType = "utilities.DynamicType{}"
		valueType = "utilities.Dynamic"
		goType = "utilities.Dynamic"
		return
	}
	if prop.Type == "object" && prop.AdditionalProperties != nil && prop.AdditionalProperties.Value.Type == "string" {
		attributeType = "types.MapType{ElemType: types.StringType}"
		valueType = "types.Map"
		goType = "map[string]string"
		return
	}
	if prop.Type == "object" && prop.AdditionalProperties != nil && prop.AdditionalProperties.Value.Type == "object" {
		if prop.AdditionalProperties.Value.AdditionalProperties != nil && prop.AdditionalProperties.Value.AdditionalProperties.Value.Type == "string" {
			attributeType = "types.MapType{ElemType: types.MapType{ElemType: types.StringType}}"
			valueType = "types.Map"
			goType = "map[string]map[string]string"
			return
		}
	}
	if prop.Type == "object" && prop.AdditionalProperties != nil && prop.AdditionalProperties.Value.Type == "array" {
		if prop.AdditionalProperties.Value.Items.Value.Type == "string" {
			attributeType = "types.MapType{ElemType: types.ListType{ElemType: types.StringType}}"
			valueType = "types.Map"
			goType = "map[string][]string"
			return
		}
	}
	if prop.Type == "array" && prop.Items != nil && prop.Items.Value != nil && prop.Items.Value.Type == "object" {
		if _, ok := prop.Items.Value.Extensions["x-kubernetes-preserve-unknown-fields"]; ok {
			attributeType = "types.ListType{ElemType: types.MapType{ElemType: types.StringType}}"
			valueType = "types.List"
			goType = "[]map[string]string"
			return
		}
		attributeType = "types.ListType{ElemType: types.ObjectType}"
		valueType = "types.List"
		goType = "[]struct"
		return
	}
	switch prop.Type {
	case "boolean":
		attributeType = "types.BoolType"
		valueType = "types.Bool"
		goType = "bool"
		return
	case "string":
		attributeType = "types.StringType"
		valueType = "types.String"
		goType = "string"
		return
	case "integer":
		attributeType = "types.Int64Type"
		valueType = "types.Int64"
		goType = "int64"
		return
	case "number":
		if prop.Format == "float" || prop.Format == "double" {
			attributeType = "types.Float64Type"
			valueType = "types.Float64"
			goType = "float64"
			return
		}
		attributeType = "utilities.DynamicNumberType{}"
		valueType = "utilities.DynamicNumber"
		goType = "utilities.DynamicNumber"
		return
	case "array":
		attributeType = "types.ListType{ElemType: types.StringType}"
		valueType = "types.List"
		goType = "[]string"
		return
	case "object":
		if len(prop.Properties) > 0 {
			attributeType = "types.ObjectType"
			valueType = "types.Object"
			goType = "struct"
			return
		}
		attributeType = "types.MapType{ElemType: types.StringType}"
		valueType = "types.Map"
		goType = "map[string]string"
		return
	}
	attributeType = "UNKNOWN"
	valueType = "UNKNOWN"
	goType = "UNKNOWN"
	return
}

func CRDv1Validators(prop apiextensionsv1.JSONSchemaProps, uv *UsedValidators) []string {
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
		uv.Float64Validator = true
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

func OpenApiValidators(prop *openapi3.Schema, uv *UsedValidators) []string {
	validators := make([]string, 0)

	if prop.Type == "integer" && prop.Min != nil {
		validators = append(validators, fmt.Sprintf("int64validator.AtLeast(%v)", prop.Min))
		uv.Int64Validator = true
	}
	if prop.Type == "integer" && prop.Max != nil {
		validators = append(validators, fmt.Sprintf("int64validator.AtMost(%v)", prop.Max))
		uv.Int64Validator = true
	}
	//if prop.Type == "integer" && len(prop.Enum) > 0 {
	//	enums := intEnums(prop.Enum)
	//	validators = append(validators, fmt.Sprintf("int64validator.OneOf(%v)", concatEnums(enums)))
	//	uv.Int64Validator = true
	//}
	if prop.Type == "number" && prop.Min != nil {
		validators = append(validators, fmt.Sprintf("float64validator.AtLeast(%v)", prop.Min))
		uv.Float64Validator = true
	}
	if prop.Type == "number" && prop.Max != nil {
		validators = append(validators, fmt.Sprintf("float64validator.AtMost(%v)", prop.Max))
		uv.Float64Validator = true
	}
	//if prop.Type == "number" && len(prop.Enum) > 0 {
	//	enums := floatEnums(prop.Enum)
	//	validators = append(validators, fmt.Sprintf("float64validator.OneOf(%v)", concatEnums(enums)))
	//	uv.Float64Validator = true
	//}
	if prop.Type == "string" && prop.Format == "byte" {
		validators = append(validators, "validators.Base64Validator()")
	}
	if prop.Type == "string" && prop.Format == "date-time" {
		validators = append(validators, "validators.DateTime64Validator()")
	}
	if prop.Type == "string" && prop.MinLength != 0 {
		validators = append(validators, fmt.Sprintf("stringvalidator.LengthAtLeast(%v)", prop.MinLength))
		uv.StringValidator = true
	}
	if prop.Type == "string" && prop.MaxLength != nil {
		validators = append(validators, fmt.Sprintf("stringvalidator.LengthAtMost(%v)", *prop.MaxLength))
		uv.StringValidator = true
	}
	//if prop.Type == "string" && len(prop.Enum) > 0 {
	//	enums := stringEnums(prop.Enum)
	//	validators = append(validators, fmt.Sprintf("stringvalidator.OneOf(%s)", concatEnums(enums)))
	//	uv.StringValidator = true
	//}
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
