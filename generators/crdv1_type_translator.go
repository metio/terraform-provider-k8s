//go:build generators

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

var _ typeTranslator = (*crd1TypeTranslator)(nil)

type crd1TypeTranslator struct {
	property *apiextensionsv1.JSONSchemaProps
}

func (t *crd1TypeTranslator) isIntOrString() bool {
	return t.property.XIntOrString || t.property.Type == "string" && t.property.Format == "int-or-string"
}

func (t *crd1TypeTranslator) isBoolean() bool {
	return t.property.Type == "boolean"
}

func (t *crd1TypeTranslator) isString() bool {
	return t.property.Type == "string"
}

func (t *crd1TypeTranslator) isInteger() bool {
	return t.property.Type == "integer"
}

func (t *crd1TypeTranslator) isNumber() bool {
	return t.property.Type == "number"
}

func (t *crd1TypeTranslator) isFloat() bool {
	return t.property.Format == "float" || t.property.Format == "double"
}

func (t *crd1TypeTranslator) isArray() bool {
	return t.property.Type == "array"
}

func (t *crd1TypeTranslator) isObject() bool {
	return t.property.Type == "object"
}

func (t *crd1TypeTranslator) hasUnknownFields() bool {
	return t.property.XPreserveUnknownFields != nil && *t.property.XPreserveUnknownFields
}

func (t *crd1TypeTranslator) hasProperties() bool {
	return len(t.property.Properties) > 0
}

func (t *crd1TypeTranslator) hasOneOf() bool {
	return len(t.property.OneOf) > 0
}

func (t *crd1TypeTranslator) isOneOfArray() bool {
	for _, oneOf := range t.property.OneOf {
		if oneOf.Type == "array" {
			return true
		}
	}
	return false
}

func (t *crd1TypeTranslator) isOneOfBoolean() bool {
	for _, oneOf := range t.property.OneOf {
		if oneOf.Type == "boolean" {
			return true
		}
	}
	return false
}

func (t *crd1TypeTranslator) isObjectWithAdditionalStringProperties() bool {
	return t.property.Type == "object" &&
		t.property.AdditionalProperties != nil &&
		t.property.AdditionalProperties.Schema.Type == "string"
}

func (t *crd1TypeTranslator) isObjectWithAdditionalObjectProperties() bool {
	return t.property.Type == "object" &&
		t.property.AdditionalProperties != nil &&
		t.property.AdditionalProperties.Schema.Type == "object"
}

func (t *crd1TypeTranslator) isObjectWithAdditionalArrayProperties() bool {
	return t.property.Type == "object" &&
		t.property.AdditionalProperties != nil &&
		t.property.AdditionalProperties.Schema.Type == "array"
}

func (t *crd1TypeTranslator) isArrayWithObjectItems() bool {
	return t.property.Type == "array" &&
		t.property.Items != nil &&
		t.property.Items.Schema.Type == "object"
}

func (t *crd1TypeTranslator) additionalPropertiesHaveStringItems() bool {
	return t.property.AdditionalProperties.Schema.Items != nil &&
		t.property.AdditionalProperties.Schema.Items.Schema.Type == "string"
}

func (t *crd1TypeTranslator) additionalPropertiesHaveProperties() bool {
	return len(t.property.AdditionalProperties.Schema.Properties) > 0
}

func (t *crd1TypeTranslator) additionalPropertiesHaveUnknownFields() bool {
	return t.property.AdditionalProperties.Schema.XPreserveUnknownFields != nil && *t.property.AdditionalProperties.Schema.XPreserveUnknownFields
}

func (t *crd1TypeTranslator) additionalPropertiesHaveAdditionalStringProperties() bool {
	return t.property.AdditionalProperties.Schema.AdditionalProperties != nil &&
		t.property.AdditionalProperties.Schema.AdditionalProperties.Schema.Type == "string"
}

func (t *crd1TypeTranslator) additionalPropertiesHaveAdditionalArrayProperties() bool {
	return t.property.AdditionalProperties.Schema.AdditionalProperties != nil &&
		t.property.AdditionalProperties.Schema.AdditionalProperties.Schema.Type == "array"
}

func (t *crd1TypeTranslator) additionalPropertiesHaveAdditionalPropertiesWithStringItems() bool {
	return t.property.AdditionalProperties.Schema.AdditionalProperties.Schema.Items != nil &&
		t.property.AdditionalProperties.Schema.AdditionalProperties.Schema.Items.Schema.Type == "string"
}

func (t *crd1TypeTranslator) itemsHaveUnknownFields() bool {
	return t.property.Items.Schema.XPreserveUnknownFields != nil && *t.property.Items.Schema.XPreserveUnknownFields
}

func (t *crd1TypeTranslator) itemsHaveAdditionalStringProperties() bool {
	return t.property.Items.Schema.AdditionalProperties != nil &&
		t.property.Items.Schema.AdditionalProperties.Schema.Type == "string"
}
