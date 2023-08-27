//go:build generators

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import (
	"github.com/getkin/kin-openapi/openapi3"
)

var _ typeTranslator = (*openapiv3TypeTranslator)(nil)

type openapiv3TypeTranslator struct {
	property *openapi3.Schema
}

func (t *openapiv3TypeTranslator) isIntOrString() bool {
	_, ok := t.property.Extensions["x-kubernetes-int-or-string"]
	return ok || t.property.Type == "string" && t.property.Format == "int-or-string"
}

func (t *openapiv3TypeTranslator) isBoolean() bool {
	return t.property.Type == "boolean"
}

func (t *openapiv3TypeTranslator) isString() bool {
	return t.property.Type == "string"
}

func (t *openapiv3TypeTranslator) isInteger() bool {
	return t.property.Type == "integer"
}

func (t *openapiv3TypeTranslator) isNumber() bool {
	return t.property.Type == "number"
}

func (t *openapiv3TypeTranslator) isFloat() bool {
	return t.property.Format == "float" || t.property.Format == "double"
}

func (t *openapiv3TypeTranslator) isArray() bool {
	return t.property.Type == "array"
}

func (t *openapiv3TypeTranslator) isObject() bool {
	return t.property.Type == "object"
}

func (t *openapiv3TypeTranslator) hasUnknownFields() bool {
	_, ok := t.property.Extensions["x-kubernetes-preserve-unknown-fields"]
	return ok
}

func (t *openapiv3TypeTranslator) hasProperties() bool {
	return len(t.property.Properties) > 0
}

func (t *openapiv3TypeTranslator) hasOneOf() bool {
	return len(t.property.OneOf) > 0
}

func (t *openapiv3TypeTranslator) isOneOfArray() bool {
	for _, oneOf := range t.property.OneOf {
		if oneOf.Value.Type == "array" {
			return true
		}
	}
	return false
}

func (t *openapiv3TypeTranslator) isOneOfBoolean() bool {
	for _, oneOf := range t.property.OneOf {
		if oneOf.Value.Type == "boolean" {
			return true
		}
	}
	return false
}

func (t *openapiv3TypeTranslator) isObjectWithAdditionalStringProperties() bool {
	return t.property.Type == "object" &&
		t.property.AdditionalProperties.Schema != nil &&
		t.property.AdditionalProperties.Schema.Value.Type == "string"
}

func (t *openapiv3TypeTranslator) isObjectWithAdditionalObjectProperties() bool {
	return t.property.Type == "object" &&
		t.property.AdditionalProperties.Schema != nil &&
		t.property.AdditionalProperties.Schema.Value.Type == "object"
}

func (t *openapiv3TypeTranslator) isObjectWithAdditionalArrayProperties() bool {
	return t.property.Type == "object" &&
		t.property.AdditionalProperties.Schema != nil &&
		t.property.AdditionalProperties.Schema.Value.Type == "array"
}

func (t *openapiv3TypeTranslator) isArrayWithObjectItems() bool {
	return t.property.Type == "array" &&
		t.property.Items != nil &&
		t.property.Items.Value != nil &&
		t.property.Items.Value.Type == "object"
}

func (t *openapiv3TypeTranslator) additionalPropertiesHaveStringItems() bool {
	return t.property.AdditionalProperties.Schema.Value.Items != nil &&
		t.property.AdditionalProperties.Schema.Value.Items.Value.Type == "string"
}

func (t *openapiv3TypeTranslator) additionalPropertiesHaveProperties() bool {
	return len(t.property.AdditionalProperties.Schema.Value.Properties) > 0
}

func (t *openapiv3TypeTranslator) additionalPropertiesHaveUnknownFields() bool {
	_, ok := t.property.AdditionalProperties.Schema.Value.Extensions["x-kubernetes-preserve-unknown-fields"]
	return ok
}

func (t *openapiv3TypeTranslator) additionalPropertiesHaveAdditionalStringProperties() bool {
	return t.property.AdditionalProperties.Schema.Value.AdditionalProperties.Schema != nil &&
		t.property.AdditionalProperties.Schema.Value.AdditionalProperties.Schema.Value.Type == "string"
}

func (t *openapiv3TypeTranslator) additionalPropertiesHaveAdditionalArrayProperties() bool {
	return t.property.AdditionalProperties.Schema.Value.AdditionalProperties.Schema != nil &&
		t.property.AdditionalProperties.Schema.Value.AdditionalProperties.Schema.Value.Type == "array"
}

func (t *openapiv3TypeTranslator) additionalPropertiesHaveAdditionalPropertiesWithStringItems() bool {
	return t.property.AdditionalProperties.Schema.Value.AdditionalProperties.Schema.Value.Items != nil &&
		t.property.AdditionalProperties.Schema.Value.AdditionalProperties.Schema.Value.Items.Value.Type == "string"
}

func (t *openapiv3TypeTranslator) itemsHaveUnknownFields() bool {
	_, ok := t.property.Items.Value.Extensions["x-kubernetes-preserve-unknown-fields"]
	return ok
}

func (t *openapiv3TypeTranslator) itemsHaveAdditionalStringProperties() bool {
	return t.property.Items.Value.AdditionalProperties.Schema != nil &&
		t.property.Items.Value.AdditionalProperties.Schema.Value.Type == "string"
}
