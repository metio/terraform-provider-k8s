//go:build generators

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */
package main

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/assert"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"testing"
)

func TestTranslateTypeWith(t *testing.T) {
	type testCase struct {
		translator    typeTranslator
		attributeType string
		valueType     string
		goType        string
	}

	testCases := map[string]testCase{
		"CRDv1/empty": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{},
			},
			attributeType: "UNKNOWN",
			valueType:     "UNKNOWN",
			goType:        "UNKNOWN",
		},
		"CRDv1/string": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "string",
				},
			},
			attributeType: "types.StringType",
			valueType:     "types.String",
			goType:        "string",
		},
		"CRDv1/boolean": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "boolean",
				},
			},
			attributeType: "types.BoolType",
			valueType:     "types.Bool",
			goType:        "bool",
		},
		"CRDv1/integer": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "integer",
				},
			},
			attributeType: "types.Int64Type",
			valueType:     "types.Int64",
			goType:        "int64",
		},
		"CRDv1/number": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "number",
				},
			},
			attributeType: "utilities.DynamicNumberType{}",
			valueType:     "utilities.DynamicNumber",
			goType:        "utilities.DynamicNumber",
		},
		"CRDv1/float-float": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type:   "number",
					Format: "float",
				},
			},
			attributeType: "types.Float64Type",
			valueType:     "types.Float64",
			goType:        "float64",
		},
		"CRDv1/float-double": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type:   "number",
					Format: "double",
				},
			},
			attributeType: "types.Float64Type",
			valueType:     "types.Float64",
			goType:        "float64",
		},
		"CRDv1/array": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "array",
				},
			},
			attributeType: "types.ListType{ElemType: types.StringType}",
			valueType:     "types.List",
			goType:        "[]string",
		},
		"CRDv1/object": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					Type: "object",
				},
			},
			attributeType: "types.MapType{ElemType: types.StringType}",
			valueType:     "types.Map",
			goType:        "map[string]string",
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
			attributeType: "types.ObjectType",
			valueType:     "types.Object",
			goType:        "struct",
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
			attributeType: "types.ListType{ElemType: types.ObjectType}",
			valueType:     "types.List",
			goType:        "[]struct",
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
			attributeType: "types.ListType{ElemType: types.MapType{ElemType: types.StringType}}",
			valueType:     "types.List",
			goType:        "[]map[string]string",
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
			attributeType: "types.ListType{ElemType: types.MapType{ElemType: types.StringType}}",
			valueType:     "types.List",
			goType:        "[]map[string]string",
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
			attributeType: "types.MapType{ElemType: types.StringType}",
			valueType:     "types.Map",
			goType:        "map[string]string",
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
			attributeType: "types.MapType{ElemType: types.ListType{ElemType: types.StringType}}",
			valueType:     "types.Map",
			goType:        "map[string][]string",
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
			attributeType: "types.ObjectType",
			valueType:     "types.Object",
			goType:        "struct",
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
			attributeType: "types.MapType{ElemType: types.MapType{ElemType: types.StringType}}",
			valueType:     "types.Map",
			goType:        "map[string]map[string]string",
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
			attributeType: "types.ObjectType",
			valueType:     "types.Object",
			goType:        "struct",
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
			attributeType: "utilities.DynamicType{}",
			valueType:     "utilities.Dynamic",
			goType:        "utilities.Dynamic",
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
			attributeType: "types.ObjectType",
			valueType:     "types.Object",
			goType:        "struct",
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
			attributeType: "types.MapType{ElemType: types.MapType{ElemType: types.ListType{ElemType: types.StringType}}}",
			valueType:     "types.Map",
			goType:        "map[string]map[string][]string",
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
			attributeType: "types.MapType{ElemType: types.StringType}",
			valueType:     "types.Map",
			goType:        "map[string]string",
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
			attributeType: "types.ListType{ElemType: types.StringType}",
			valueType:     "types.List",
			goType:        "[]string",
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
			attributeType: "types.BoolType",
			valueType:     "types.Bool",
			goType:        "bool",
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
			valueType:     "UNKNOWN",
			goType:        "UNKNOWN",
		},
		"CRDv1/unknown-fields": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					XPreserveUnknownFields: Ptr(true),
				},
			},
			attributeType: "utilities.DynamicType{}",
			valueType:     "utilities.Dynamic",
			goType:        "utilities.Dynamic",
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
			attributeType: "types.ObjectType",
			valueType:     "types.Object",
			goType:        "struct",
		},
		"CRDv1/int-or-string": {
			translator: &crd1TypeTranslator{
				property: &apiextensionsv1.JSONSchemaProps{
					XIntOrString: true,
				},
			},
			attributeType: "utilities.IntOrStringType{}",
			valueType:     "utilities.IntOrString",
			goType:        "utilities.IntOrString",
		},

		"OpenAPIv3/empty": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{},
			},
			attributeType: "UNKNOWN",
			valueType:     "UNKNOWN",
			goType:        "UNKNOWN",
		},
		"OpenAPIv3/string": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					Type: "string",
				},
			},
			attributeType: "types.StringType",
			valueType:     "types.String",
			goType:        "string",
		},
		"OpenAPIv3/boolean": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					Type: "boolean",
				},
			},
			attributeType: "types.BoolType",
			valueType:     "types.Bool",
			goType:        "bool",
		},
		"OpenAPIv3/integer": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					Type: "integer",
				},
			},
			attributeType: "types.Int64Type",
			valueType:     "types.Int64",
			goType:        "int64",
		},
		"OpenAPIv3/number": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					Type: "number",
				},
			},
			attributeType: "utilities.DynamicNumberType{}",
			valueType:     "utilities.DynamicNumber",
			goType:        "utilities.DynamicNumber",
		},
		"OpenAPIv3/float-float": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					Type:   "number",
					Format: "float",
				},
			},
			attributeType: "types.Float64Type",
			valueType:     "types.Float64",
			goType:        "float64",
		},
		"OpenAPIv3/float-double": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					Type:   "number",
					Format: "double",
				},
			},
			attributeType: "types.Float64Type",
			valueType:     "types.Float64",
			goType:        "float64",
		},
		"OpenAPIv3/array": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					Type: "array",
				},
			},
			attributeType: "types.ListType{ElemType: types.StringType}",
			valueType:     "types.List",
			goType:        "[]string",
		},
		"OpenAPIv3/object": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					Type: "object",
				},
			},
			attributeType: "types.MapType{ElemType: types.StringType}",
			valueType:     "types.Map",
			goType:        "map[string]string",
		},
		"OpenAPIv3/object-with-properties": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					Type: "object",
					Properties: openapi3.Schemas{
						"first": {},
					},
				},
			},
			attributeType: "types.ObjectType",
			valueType:     "types.Object",
			goType:        "struct",
		},
		"OpenAPIv3/array-of-objects": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					Type: "array",
					Items: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: "object",
						},
					},
				},
			},
			attributeType: "types.ListType{ElemType: types.ObjectType}",
			valueType:     "types.List",
			goType:        "[]struct",
		},
		"OpenAPIv3/array-of-objects-with-unknown-fields": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					Type: "array",
					Items: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: "object",
							ExtensionProps: openapi3.ExtensionProps{
								Extensions: map[string]interface{}{
									"x-kubernetes-preserve-unknown-fields": "true",
								},
							},
						},
					},
				},
			},
			attributeType: "types.ListType{ElemType: types.MapType{ElemType: types.StringType}}",
			valueType:     "types.List",
			goType:        "[]map[string]string",
		},
		"OpenAPIv3/array-of-objects-with-additional-string-properties": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					Type: "array",
					Items: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: "object",
							AdditionalProperties: &openapi3.SchemaRef{
								Value: &openapi3.Schema{
									Type: "string",
								},
							},
						},
					},
				},
			},
			attributeType: "types.ListType{ElemType: types.MapType{ElemType: types.StringType}}",
			valueType:     "types.List",
			goType:        "[]map[string]string",
		},
		"OpenAPIv3/object-with-additional-array-properties": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					Type: "object",
					AdditionalProperties: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: "array",
						},
					},
				},
			},
			attributeType: "types.MapType{ElemType: types.StringType}",
			valueType:     "types.Map",
			goType:        "map[string]string",
		},
		"OpenAPIv3/object-with-additional-array-properties-having-string-items": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					Type: "object",
					AdditionalProperties: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: "array",
							Items: &openapi3.SchemaRef{
								Value: &openapi3.Schema{
									Type: "string",
								},
							},
						},
					},
				},
			},
			attributeType: "types.MapType{ElemType: types.ListType{ElemType: types.StringType}}",
			valueType:     "types.Map",
			goType:        "map[string][]string",
		},
		"OpenAPIv3/object-with-additional-object-properties": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					Type: "object",
					AdditionalProperties: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: "object",
						},
					},
				},
			},
			attributeType: "types.ObjectType",
			valueType:     "types.Object",
			goType:        "struct",
		},
		"OpenAPIv3/object-with-additional-object-properties-having-additional-string-properties": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					Type: "object",
					AdditionalProperties: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: "object",
							AdditionalProperties: &openapi3.SchemaRef{
								Value: &openapi3.Schema{
									Type: "string",
								},
							},
						},
					},
				},
			},
			attributeType: "types.MapType{ElemType: types.MapType{ElemType: types.StringType}}",
			valueType:     "types.Map",
			goType:        "map[string]map[string]string",
		},
		"OpenAPIv3/object-with-additional-object-properties-having-additional-array-properties": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					Type: "object",
					AdditionalProperties: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: "object",
							AdditionalProperties: &openapi3.SchemaRef{
								Value: &openapi3.Schema{
									Type: "array",
								},
							},
						},
					},
				},
			},
			attributeType: "types.ObjectType",
			valueType:     "types.Object",
			goType:        "struct",
		},
		"OpenAPIv3/object-with-additional-object-properties-having-unknown-fields": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					Type: "object",
					AdditionalProperties: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: "object",
							ExtensionProps: openapi3.ExtensionProps{
								Extensions: map[string]interface{}{
									"x-kubernetes-preserve-unknown-fields": "true",
								},
							},
						},
					},
				},
			},
			attributeType: "utilities.DynamicType{}",
			valueType:     "utilities.Dynamic",
			goType:        "utilities.Dynamic",
		},
		"OpenAPIv3/object-with-additional-object-properties-having-properties": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					Type: "object",
					AdditionalProperties: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: "object",
							ExtensionProps: openapi3.ExtensionProps{
								Extensions: map[string]interface{}{
									"x-kubernetes-preserve-unknown-fields": "true",
								},
							},
							Properties: openapi3.Schemas{
								"firs": {},
							},
						},
					},
				},
			},
			attributeType: "types.ObjectType",
			valueType:     "types.Object",
			goType:        "struct",
		},
		"OpenAPIv3/object-with-additional-object-properties-having-additional-array-properties-with-string-items": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					Type: "object",
					AdditionalProperties: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: "object",
							AdditionalProperties: &openapi3.SchemaRef{
								Value: &openapi3.Schema{
									Type: "array",
									Items: &openapi3.SchemaRef{
										Value: &openapi3.Schema{
											Type: "string",
										},
									},
								},
							},
						},
					},
				},
			},
			attributeType: "types.MapType{ElemType: types.MapType{ElemType: types.ListType{ElemType: types.StringType}}}",
			valueType:     "types.Map",
			goType:        "map[string]map[string][]string",
		},
		"OpenAPIv3/object-with-additional-string-properties": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					Type: "object",
					AdditionalProperties: &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: "string",
						},
					},
				},
			},
			attributeType: "types.MapType{ElemType: types.StringType}",
			valueType:     "types.Map",
			goType:        "map[string]string",
		},
		"OpenAPIv3/one-of-array": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					OneOf: openapi3.SchemaRefs{
						&openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Type: "array",
							},
						},
					},
				},
			},
			attributeType: "types.ListType{ElemType: types.StringType}",
			valueType:     "types.List",
			goType:        "[]string",
		},
		"OpenAPIv3/one-of-boolean": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					OneOf: openapi3.SchemaRefs{
						&openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Type: "boolean",
							},
						},
					},
				},
			},
			attributeType: "types.BoolType",
			valueType:     "types.Bool",
			goType:        "bool",
		},
		"OpenAPIv3/one-of-string": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					OneOf: openapi3.SchemaRefs{
						&openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Type: "string",
							},
						},
					},
				},
			},
			attributeType: "UNKNOWN",
			valueType:     "UNKNOWN",
			goType:        "UNKNOWN",
		},
		"OpenAPIv3/unknown-fields": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					ExtensionProps: openapi3.ExtensionProps{
						Extensions: map[string]interface{}{
							"x-kubernetes-preserve-unknown-fields": "true",
						},
					},
				},
			},
			attributeType: "utilities.DynamicType{}",
			valueType:     "utilities.Dynamic",
			goType:        "utilities.Dynamic",
		},
		"OpenAPIv3/unknown-fields-with-properties": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					ExtensionProps: openapi3.ExtensionProps{
						Extensions: map[string]interface{}{
							"x-kubernetes-preserve-unknown-fields": "true",
						},
					},
					Properties: openapi3.Schemas{
						"firs": {},
					},
				},
			},
			attributeType: "types.ObjectType",
			valueType:     "types.Object",
			goType:        "struct",
		},
		"OpenAPIv3/int-or-string": {
			translator: &openapiv3TypeTranslator{
				property: &openapi3.Schema{
					ExtensionProps: openapi3.ExtensionProps{
						Extensions: map[string]interface{}{
							"x-kubernetes-int-or-string": "true",
						},
					},
				},
			},
			attributeType: "utilities.IntOrStringType{}",
			valueType:     "utilities.IntOrString",
			goType:        "utilities.IntOrString",
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			attributeType, valueType, goType := translateTypeWith(test.translator)

			assert.Equal(t, test.attributeType, attributeType, "attributeType")
			assert.Equal(t, test.valueType, valueType, "valueType")
			assert.Equal(t, test.goType, goType, "goType")
		})
	}
}

func Ptr[T any](v T) *T {
	return &v
}
