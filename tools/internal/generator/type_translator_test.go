/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */
package generator

import (
	"github.com/pb33f/libopenapi/datamodel/high/base"
	"github.com/pb33f/libopenapi/orderedmap"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"testing"
)

func TestTranslateTypeWith(t *testing.T) {
	type testCase struct {
		translator            typeTranslator
		attributeType         string
		elementType           string
		valueType             string
		goType                string
		customType            string
		terraformResourceName string
		propPath              string
	}

	testCases := map[string]testCase{
		"CRDv1/empty": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{},
			},
			attributeType: "UNKNOWN",
			elementType:   "UNKNOWN",
			valueType:     "UNKNOWN",
			goType:        "UNKNOWN",
			customType:    "UNKNOWN",
		},
		"CRDv1/string": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "string",
				},
			},
			attributeType: "schema.StringAttribute",
			elementType:   "",
			valueType:     "types.String",
			goType:        "string",
			customType:    "",
		},
		"CRDv1/boolean": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "boolean",
				},
			},
			attributeType: "schema.BoolAttribute",
			elementType:   "",
			valueType:     "types.Bool",
			goType:        "bool",
			customType:    "",
		},
		"CRDv1/integer": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "integer",
				},
			},
			attributeType: "schema.Int64Attribute",
			elementType:   "",
			valueType:     "types.Int64",
			goType:        "int64",
			customType:    "",
		},
		"CRDv1/number": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "number",
				},
			},
			attributeType: "schema.Float64Attribute",
			elementType:   "",
			valueType:     "types.Float64",
			goType:        "float64",
			customType:    "",
		},
		"CRDv1/float-float": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type:   "number",
					Format: "float",
				},
			},
			attributeType: "schema.Float64Attribute",
			elementType:   "",
			valueType:     "types.Float64",
			goType:        "float64",
			customType:    "",
		},
		"CRDv1/float-double": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type:   "number",
					Format: "double",
				},
			},
			attributeType: "schema.Float64Attribute",
			elementType:   "",
			valueType:     "types.Float64",
			goType:        "float64",
			customType:    "",
		},
		"CRDv1/array": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "array",
				},
			},
			attributeType: "schema.ListAttribute",
			elementType:   "types.StringType",
			valueType:     "types.List",
			goType:        "[]string",
			customType:    "",
		},
		"CRDv1/object": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "object",
				},
			},
			attributeType: "schema.MapAttribute",
			elementType:   "types.StringType",
			valueType:     "types.Map",
			goType:        "map[string]string",
			customType:    "",
		},
		"CRDv1/object-with-properties": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "object",
					Properties: map[string]apiextensionsv1.JSONSchemaProps{
						"first": {},
					},
				},
			},
			attributeType: "schema.SingleNestedAttribute",
			elementType:   "",
			valueType:     "types.Object",
			goType:        "struct",
			customType:    "",
		},
		"CRDv1/array-of-objects": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "array",
					Items: &apiextensionsv1.JSONSchemaPropsOrArray{
						Schema: &apiextensionsv1.JSONSchemaProps{
							Type: "object",
						},
					},
				},
			},
			attributeType: "schema.ListNestedAttribute",
			elementType:   "",
			valueType:     "types.List",
			goType:        "[]struct",
			customType:    "",
		},
		"CRDv1/array-of-objects-with-unknown-fields": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "array",
					Items: &apiextensionsv1.JSONSchemaPropsOrArray{
						Schema: &apiextensionsv1.JSONSchemaProps{
							Type:                   "object",
							XPreserveUnknownFields: Ptr(true),
						},
					},
				},
			},
			attributeType: "schema.ListAttribute",
			elementType:   "types.MapType{ElemType: types.StringType}",
			valueType:     "types.List",
			goType:        "[]map[string]string",
			customType:    "",
		},
		"CRDv1/array-of-objects-with-additional-string-properties": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "array",
					Items: &apiextensionsv1.JSONSchemaPropsOrArray{
						Schema: &apiextensionsv1.JSONSchemaProps{
							Type: "object",
							AdditionalProperties: &apiextensionsv1.JSONSchemaPropsOrBool{
								Schema: &apiextensionsv1.JSONSchemaProps{
									Type: "string",
								},
							},
						},
					},
				},
			},
			attributeType: "schema.ListAttribute",
			elementType:   "types.MapType{ElemType: types.StringType}",
			valueType:     "types.List",
			goType:        "[]map[string]string",
			customType:    "",
		},
		"CRDv1/object-with-additional-array-properties": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "object",
					AdditionalProperties: &apiextensionsv1.JSONSchemaPropsOrBool{
						Schema: &apiextensionsv1.JSONSchemaProps{
							Type: "array",
						},
					},
				},
			},
			attributeType: "schema.MapAttribute",
			elementType:   "types.StringType",
			valueType:     "types.Map",
			goType:        "map[string]string",
			customType:    "",
		},
		"CRDv1/object-with-additional-array-properties-having-string-items": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "object",
					AdditionalProperties: &apiextensionsv1.JSONSchemaPropsOrBool{
						Schema: &apiextensionsv1.JSONSchemaProps{
							Type: "array",
							Items: &apiextensionsv1.JSONSchemaPropsOrArray{
								Schema: &apiextensionsv1.JSONSchemaProps{
									Type: "string",
								},
							},
						},
					},
				},
			},
			attributeType: "schema.MapAttribute",
			elementType:   "types.ListType{ElemType: types.StringType}",
			valueType:     "types.Map",
			goType:        "map[string][]string",
			customType:    "",
		},
		"CRDv1/object-with-additional-object-properties": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "object",
					AdditionalProperties: &apiextensionsv1.JSONSchemaPropsOrBool{
						Schema: &apiextensionsv1.JSONSchemaProps{
							Type: "object",
						},
					},
				},
			},
			attributeType: "schema.SingleNestedAttribute",
			elementType:   "",
			valueType:     "types.Object",
			goType:        "struct",
			customType:    "",
		},
		"CRDv1/object-with-additional-object-properties-having-additional-string-properties": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "object",
					AdditionalProperties: &apiextensionsv1.JSONSchemaPropsOrBool{
						Schema: &apiextensionsv1.JSONSchemaProps{
							Type: "object",
							AdditionalProperties: &apiextensionsv1.JSONSchemaPropsOrBool{
								Schema: &apiextensionsv1.JSONSchemaProps{
									Type: "string",
								},
							},
						},
					},
				},
			},
			attributeType: "schema.MapAttribute",
			elementType:   "types.MapType{ElemType: types.StringType}",
			valueType:     "types.Map",
			goType:        "map[string]map[string]string",
			customType:    "",
		},
		"CRDv1/object-with-additional-object-properties-having-additional-array-properties": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "object",
					AdditionalProperties: &apiextensionsv1.JSONSchemaPropsOrBool{
						Schema: &apiextensionsv1.JSONSchemaProps{
							Type: "object",
							AdditionalProperties: &apiextensionsv1.JSONSchemaPropsOrBool{
								Schema: &apiextensionsv1.JSONSchemaProps{
									Type: "array",
								},
							},
						},
					},
				},
			},
			attributeType: "schema.SingleNestedAttribute",
			elementType:   "",
			valueType:     "types.Object",
			goType:        "struct",
			customType:    "",
		},
		"CRDv1/object-with-additional-object-properties-having-unknown-fields": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "object",
					AdditionalProperties: &apiextensionsv1.JSONSchemaPropsOrBool{
						Schema: &apiextensionsv1.JSONSchemaProps{
							Type:                   "object",
							XPreserveUnknownFields: Ptr(true),
						},
					},
				},
			},
			attributeType: "schema.MapAttribute",
			elementType:   "types.StringType",
			valueType:     "types.Map",
			goType:        "map[string]string",
			customType:    "",
		},
		"CRDv1/object-with-additional-object-properties-having-properties": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "object",
					AdditionalProperties: &apiextensionsv1.JSONSchemaPropsOrBool{
						Schema: &apiextensionsv1.JSONSchemaProps{
							Type:                   "object",
							XPreserveUnknownFields: Ptr(true),
							Properties: map[string]apiextensionsv1.JSONSchemaProps{
								"firs": {},
							},
						},
					},
				},
			},
			attributeType: "schema.SingleNestedAttribute",
			elementType:   "",
			valueType:     "types.Object",
			goType:        "struct",
			customType:    "",
		},
		"CRDv1/object-with-additional-object-properties-having-additional-array-properties-with-string-items": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "object",
					AdditionalProperties: &apiextensionsv1.JSONSchemaPropsOrBool{
						Schema: &apiextensionsv1.JSONSchemaProps{
							Type: "object",
							AdditionalProperties: &apiextensionsv1.JSONSchemaPropsOrBool{
								Schema: &apiextensionsv1.JSONSchemaProps{
									Type: "array",
									Items: &apiextensionsv1.JSONSchemaPropsOrArray{
										Schema: &apiextensionsv1.JSONSchemaProps{
											Type: "string",
										},
									},
								},
							},
						},
					},
				},
			},
			attributeType: "schema.MapAttribute",
			elementType:   "types.MapType{ElemType: types.ListType{ElemType: types.StringType}}",
			valueType:     "types.Map",
			goType:        "map[string]map[string][]string",
			customType:    "",
		},
		"CRDv1/object-with-additional-string-properties": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "object",
					AdditionalProperties: &apiextensionsv1.JSONSchemaPropsOrBool{
						Schema: &apiextensionsv1.JSONSchemaProps{
							Type: "string",
						},
					},
				},
			},
			attributeType: "schema.MapAttribute",
			elementType:   "types.StringType",
			valueType:     "types.Map",
			goType:        "map[string]string",
			customType:    "",
		},
		"CRDv1/one-of-array": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					OneOf: []apiextensionsv1.JSONSchemaProps{
						{
							Type: "array",
						},
					},
				},
			},
			attributeType: "schema.ListAttribute",
			elementType:   "types.StringType",
			valueType:     "types.List",
			goType:        "[]string",
			customType:    "",
		},
		"CRDv1/one-of-boolean": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					OneOf: []apiextensionsv1.JSONSchemaProps{
						{
							Type: "boolean",
						},
					},
				},
			},
			attributeType: "schema.BoolAttribute",
			elementType:   "",
			valueType:     "types.Bool",
			goType:        "bool",
			customType:    "",
		},
		"CRDv1/one-of-string": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					OneOf: []apiextensionsv1.JSONSchemaProps{
						{
							Type: "string",
						},
					},
				},
			},
			attributeType: "UNKNOWN",
			elementType:   "UNKNOWN",
			valueType:     "UNKNOWN",
			goType:        "UNKNOWN",
			customType:    "UNKNOWN",
		},
		"CRDv1/unknown-fields": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					XPreserveUnknownFields: Ptr(true),
				},
			},
			attributeType: "schema.MapAttribute",
			elementType:   "types.StringType",
			valueType:     "types.Map",
			goType:        "map[string]string",
			customType:    "",
		},
		"CRDv1/unknown-fields-with-properties": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					XPreserveUnknownFields: Ptr(true),
					Properties: map[string]apiextensionsv1.JSONSchemaProps{
						"firs": {},
					},
				},
			},
			attributeType: "schema.SingleNestedAttribute",
			elementType:   "",
			valueType:     "types.Object",
			goType:        "struct",
			customType:    "",
		},
		"CRDv1/int-or-string": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					XIntOrString: true,
				},
			},
			attributeType: "schema.StringAttribute",
			elementType:   "",
			valueType:     "types.String",
			goType:        "string",
			customType:    "",
		},
		"CRDv1/string-or-int": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type:   "string",
					Format: "int-or-string",
				},
			},
			attributeType: "schema.StringAttribute",
			elementType:   "",
			valueType:     "types.String",
			goType:        "string",
			customType:    "",
		},

		"OpenAPIv2/empty": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{},
			},
			attributeType: "UNKNOWN",
			elementType:   "UNKNOWN",
			valueType:     "UNKNOWN",
			goType:        "UNKNOWN",
			customType:    "UNKNOWN",
		},
		"OpenAPIv2/string": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Type: []string{"string"},
				},
			},
			attributeType: "schema.StringAttribute",
			elementType:   "",
			valueType:     "types.String",
			goType:        "string",
			customType:    "",
		},
		"OpenAPIv2/boolean": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Type: []string{"boolean"},
				},
			},
			attributeType: "schema.BoolAttribute",
			elementType:   "",
			valueType:     "types.Bool",
			goType:        "bool",
			customType:    "",
		},
		"OpenAPIv2/integer": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Type: []string{"integer"},
				},
			},
			attributeType: "schema.Int64Attribute",
			elementType:   "",
			valueType:     "types.Int64",
			goType:        "int64",
			customType:    "",
		},
		"OpenAPIv2/number": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Type: []string{"number"},
				},
			},
			attributeType: "schema.Float64Attribute",
			elementType:   "",
			valueType:     "types.Float64",
			goType:        "float64",
			customType:    "",
		},
		"OpenAPIv2/float-float": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Type:   []string{"number"},
					Format: "float",
				},
			},
			attributeType: "schema.Float64Attribute",
			elementType:   "",
			valueType:     "types.Float64",
			goType:        "float64",
			customType:    "",
		},
		"OpenAPIv2/float-double": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Type:   []string{"number"},
					Format: "double",
				},
			},
			attributeType: "schema.Float64Attribute",
			elementType:   "",
			valueType:     "types.Float64",
			goType:        "float64",
			customType:    "",
		},
		"OpenAPIv2/array": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Type: []string{"array"},
				},
			},
			attributeType: "schema.ListAttribute",
			elementType:   "types.StringType",
			valueType:     "types.List",
			goType:        "[]string",
			customType:    "",
		},
		"OpenAPIv2/object": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Type: []string{"object"},
				},
			},
			attributeType: "schema.MapAttribute",
			elementType:   "types.StringType",
			valueType:     "types.Map",
			goType:        "map[string]string",
			customType:    "",
		},
		"OpenAPIv2/object-with-properties": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Type:       []string{"object"},
					Properties: orderedmap.FromPairs(orderedmap.NewPair("first", &base.SchemaProxy{})),
				},
			},
			attributeType: "schema.SingleNestedAttribute",
			elementType:   "",
			valueType:     "types.Object",
			goType:        "struct",
			customType:    "",
		},
		"OpenAPIv2/array-of-objects": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Type: []string{"array"},
					Items: &base.DynamicValue[*base.SchemaProxy, bool]{
						N: 0,
						A: base.CreateSchemaProxy(&base.Schema{Type: []string{"object"}}),
						B: false,
					},
				},
			},
			attributeType: "schema.ListNestedAttribute",
			elementType:   "",
			valueType:     "types.List",
			goType:        "[]struct",
			customType:    "",
		},
		"OpenAPIv2/array-of-objects-with-unknown-fields": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Type: []string{"array"},
					Items: &base.DynamicValue[*base.SchemaProxy, bool]{
						N: 0,
						A: base.CreateSchemaProxy(&base.Schema{
							Type:       []string{"object"},
							Extensions: orderedmap.FromPairs(orderedmap.NewPair("x-kubernetes-preserve-unknown-fields", &yaml.Node{Value: "true"})),
						}),
						B: false,
					},
				},
			},
			attributeType: "schema.ListAttribute",
			elementType:   "types.MapType{ElemType: types.StringType}",
			valueType:     "types.List",
			goType:        "[]map[string]string",
			customType:    "",
		},
		"OpenAPIv2/array-of-objects-with-additional-string-properties": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Type: []string{"array"},
					Items: &base.DynamicValue[*base.SchemaProxy, bool]{
						N: 0,
						A: base.CreateSchemaProxy(&base.Schema{
							Type: []string{"object"},
							AdditionalProperties: &base.DynamicValue[*base.SchemaProxy, bool]{
								N: 0,
								A: base.CreateSchemaProxy(&base.Schema{
									Type: []string{"string"},
								}),
								B: false,
							},
						}),
						B: false,
					},
				},
			},
			attributeType: "schema.ListAttribute",
			elementType:   "types.MapType{ElemType: types.StringType}",
			valueType:     "types.List",
			goType:        "[]map[string]string",
			customType:    "",
		},
		"OpenAPIv2/object-with-additional-array-properties": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Type: []string{"object"},
					AdditionalProperties: &base.DynamicValue[*base.SchemaProxy, bool]{
						N: 0,
						A: base.CreateSchemaProxy(&base.Schema{
							Type: []string{"array"},
						}),
						B: false,
					},
				},
			},
			attributeType: "schema.MapAttribute",
			elementType:   "types.StringType",
			valueType:     "types.Map",
			goType:        "map[string]string",
			customType:    "",
		},
		"OpenAPIv2/object-with-additional-array-properties-having-string-items": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Type: []string{"object"},
					AdditionalProperties: &base.DynamicValue[*base.SchemaProxy, bool]{
						N: 0,
						A: base.CreateSchemaProxy(&base.Schema{
							Type: []string{"array"},
							Items: &base.DynamicValue[*base.SchemaProxy, bool]{
								N: 0,
								A: base.CreateSchemaProxy(&base.Schema{
									Type: []string{"string"},
								}),
								B: false,
							},
						}),
						B: false,
					},
				},
			},
			attributeType: "schema.MapAttribute",
			elementType:   "types.ListType{ElemType: types.StringType}",
			valueType:     "types.Map",
			goType:        "map[string][]string",
			customType:    "",
		},
		"OpenAPIv2/object-with-additional-object-properties": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Type: []string{"object"},
					AdditionalProperties: &base.DynamicValue[*base.SchemaProxy, bool]{
						N: 0,
						A: base.CreateSchemaProxy(&base.Schema{
							Type: []string{"object"},
						}),
						B: false,
					},
				},
			},
			attributeType: "schema.SingleNestedAttribute",
			elementType:   "",
			valueType:     "types.Object",
			goType:        "struct",
			customType:    "",
		},
		"OpenAPIv2/object-with-additional-object-properties-having-additional-string-properties": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Type: []string{"object"},
					AdditionalProperties: &base.DynamicValue[*base.SchemaProxy, bool]{
						N: 0,
						A: base.CreateSchemaProxy(&base.Schema{
							Type: []string{"object"},
							AdditionalProperties: &base.DynamicValue[*base.SchemaProxy, bool]{
								N: 0,
								A: base.CreateSchemaProxy(&base.Schema{
									Type: []string{"string"},
								}),
								B: false,
							},
						}),
						B: false,
					},
				},
			},
			attributeType: "schema.MapAttribute",
			elementType:   "types.MapType{ElemType: types.StringType}",
			valueType:     "types.Map",
			goType:        "map[string]map[string]string",
			customType:    "",
		},
		"OpenAPIv2/object-with-additional-object-properties-having-additional-array-properties": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Type: []string{"object"},
					AdditionalProperties: &base.DynamicValue[*base.SchemaProxy, bool]{
						N: 0,
						A: base.CreateSchemaProxy(&base.Schema{
							Type: []string{"object"},
							AdditionalProperties: &base.DynamicValue[*base.SchemaProxy, bool]{
								N: 0,
								A: base.CreateSchemaProxy(&base.Schema{
									Type: []string{"array"},
								}),
								B: false,
							},
						}),
						B: false,
					},
				},
			},
			attributeType: "schema.SingleNestedAttribute",
			elementType:   "",
			valueType:     "types.Object",
			goType:        "struct",
			customType:    "",
		},
		"OpenAPIv2/object-with-additional-object-properties-having-unknown-fields": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Type: []string{"object"},
					AdditionalProperties: &base.DynamicValue[*base.SchemaProxy, bool]{
						N: 0,
						A: base.CreateSchemaProxy(&base.Schema{
							Type:       []string{"object"},
							Extensions: orderedmap.FromPairs(orderedmap.NewPair("x-kubernetes-preserve-unknown-fields", &yaml.Node{Value: "true"})),
						}),
						B: false,
					},
				},
			},
			attributeType: "schema.MapAttribute",
			elementType:   "types.StringType",
			valueType:     "types.Map",
			goType:        "map[string]string",
			customType:    "",
		},
		"OpenAPIv2/object-with-additional-object-properties-having-properties": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Type: []string{"object"},
					AdditionalProperties: &base.DynamicValue[*base.SchemaProxy, bool]{
						N: 0,
						A: base.CreateSchemaProxy(&base.Schema{
							Type:       []string{"object"},
							Extensions: orderedmap.FromPairs(orderedmap.NewPair("x-kubernetes-preserve-unknown-fields", &yaml.Node{Value: "true"})),
							Properties: orderedmap.FromPairs(orderedmap.NewPair("first", base.CreateSchemaProxy(&base.Schema{}))),
						}),
						B: false,
					},
				},
			},
			attributeType: "schema.SingleNestedAttribute",
			elementType:   "",
			valueType:     "types.Object",
			goType:        "struct",
			customType:    "",
		},
		"OpenAPIv2/object-with-additional-object-properties-having-additional-array-properties-with-string-items": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Type: []string{"object"},
					AdditionalProperties: &base.DynamicValue[*base.SchemaProxy, bool]{
						N: 0,
						A: base.CreateSchemaProxy(&base.Schema{
							Type: []string{"object"},
							AdditionalProperties: &base.DynamicValue[*base.SchemaProxy, bool]{
								N: 0,
								A: base.CreateSchemaProxy(&base.Schema{
									Type: []string{"array"},
									Items: &base.DynamicValue[*base.SchemaProxy, bool]{
										N: 0,
										A: base.CreateSchemaProxy(&base.Schema{
											Type: []string{"string"},
										}),
										B: false,
									},
								}),
								B: false,
							},
						}),
						B: false,
					},
				},
			},
			attributeType: "schema.MapAttribute",
			elementType:   "types.MapType{ElemType: types.ListType{ElemType: types.StringType}}",
			valueType:     "types.Map",
			goType:        "map[string]map[string][]string",
			customType:    "",
		},
		"OpenAPIv2/object-with-additional-string-properties": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Type: []string{"object"},
					AdditionalProperties: &base.DynamicValue[*base.SchemaProxy, bool]{
						N: 0,
						A: base.CreateSchemaProxy(&base.Schema{
							Type: []string{"string"},
						}),
						B: false,
					},
				},
			},
			attributeType: "schema.MapAttribute",
			elementType:   "types.StringType",
			valueType:     "types.Map",
			goType:        "map[string]string",
			customType:    "",
		},
		"OpenAPIv2/one-of-array": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					OneOf: []*base.SchemaProxy{
						base.CreateSchemaProxy(&base.Schema{
							Type: []string{"array"},
						}),
					},
				},
			},
			attributeType: "schema.ListAttribute",
			elementType:   "types.StringType",
			valueType:     "types.List",
			goType:        "[]string",
			customType:    "",
		},
		"OpenAPIv2/one-of-boolean": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					OneOf: []*base.SchemaProxy{
						base.CreateSchemaProxy(&base.Schema{
							Type: []string{"boolean"},
						}),
					},
				},
			},
			attributeType: "schema.BoolAttribute",
			elementType:   "",
			valueType:     "types.Bool",
			goType:        "bool",
			customType:    "",
		},
		"OpenAPIv2/one-of-string": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					OneOf: []*base.SchemaProxy{
						base.CreateSchemaProxy(&base.Schema{
							Type: []string{"string"},
						}),
					},
				},
			},
			attributeType: "UNKNOWN",
			elementType:   "UNKNOWN",
			valueType:     "UNKNOWN",
			goType:        "UNKNOWN",
			customType:    "UNKNOWN",
		},
		"OpenAPIv2/unknown-fields": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Extensions: orderedmap.FromPairs(orderedmap.NewPair("x-kubernetes-preserve-unknown-fields", &yaml.Node{Value: "true"})),
				},
			},
			attributeType: "schema.MapAttribute",
			elementType:   "types.StringType",
			valueType:     "types.Map",
			goType:        "map[string]string",
			customType:    "",
		},
		"OpenAPIv2/unknown-fields-with-properties": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Extensions: orderedmap.FromPairs(orderedmap.NewPair("x-kubernetes-preserve-unknown-fields", &yaml.Node{Value: "true"})),
					Properties: orderedmap.FromPairs(orderedmap.NewPair("first", base.CreateSchemaProxy(&base.Schema{}))),
				},
			},
			attributeType: "schema.SingleNestedAttribute",
			elementType:   "",
			valueType:     "types.Object",
			goType:        "struct",
			customType:    "",
		},
		"OpenAPIv2/int-or-string": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Extensions: orderedmap.FromPairs(orderedmap.NewPair("x-kubernetes-int-or-string", &yaml.Node{Value: "true"})),
				},
			},
			attributeType: "schema.StringAttribute",
			elementType:   "",
			valueType:     "types.String",
			goType:        "string",
			customType:    "",
		},
		"OpenAPIv2/string-or-int": {
			translator: &openapiv2TypeTranslator{
				property: &base.Schema{
					Type:   []string{"string"},
					Format: "int-or-string",
				},
			},
			attributeType: "schema.StringAttribute",
			elementType:   "",
			valueType:     "types.String",
			goType:        "string",
			customType:    "",
		},
		"custom-types/kyverno_io_cluster_policy_v1/spec.rules.context.apiCall.data.value": {
			translator:            &openapiv2TypeTranslator{},
			attributeType:         "schema.StringAttribute",
			elementType:           "",
			valueType:             "types.String",
			goType:                "custom_types.Normalized",
			customType:            "custom_types.NormalizedType{}",
			terraformResourceName: "kyverno_io_cluster_policy_v1",
			propPath:              "spec.rules.context.apiCall.data.value",
		},
	}

	for name, test := range testCases {
		test := test
		t.Run(name, func(t *testing.T) {
			attributeType, valueType, elementType, goType, customType := translateTypeWith(test.translator, test.terraformResourceName, test.propPath)

			assert.Equal(t, test.attributeType, attributeType, "attributeType")
			assert.Equal(t, test.valueType, valueType, "valueType")
			assert.Equal(t, test.elementType, elementType, "elementType")
			assert.Equal(t, test.goType, goType, "goType")
			assert.Equal(t, test.customType, customType, "customType")
		})
	}
}

func Ptr[T any](v T) *T {
	return &v
}
