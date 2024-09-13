/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"github.com/pb33f/libopenapi/datamodel/high/base"
)

var _ typeTranslator = (*openapiv2TypeTranslator)(nil)

type openapiv2TypeTranslator struct {
	property *base.Schema
}

func (t *openapiv2TypeTranslator) hasNoType() bool {
	return t.property.Type == nil || len(t.property.Type) == 0
}

func (t *openapiv2TypeTranslator) isIntOrString() bool {
	if t.property.Extensions != nil {
		_, ok := t.property.Extensions.Get("x-kubernetes-int-or-string")
		return ok || len(t.property.Type) > 0 && t.property.Type[0] == "string" && t.property.Format == "int-or-string"
	}
	return false
}

func (t *openapiv2TypeTranslator) isBoolean() bool {
	return len(t.property.Type) > 0 && t.property.Type[0] == "boolean"
}

func (t *openapiv2TypeTranslator) isString() bool {
	return len(t.property.Type) > 0 && t.property.Type[0] == "string"
}

func (t *openapiv2TypeTranslator) isInteger() bool {
	return len(t.property.Type) > 0 && t.property.Type[0] == "integer"
}

func (t *openapiv2TypeTranslator) isNumber() bool {
	return len(t.property.Type) > 0 && t.property.Type[0] == "number"
}

func (t *openapiv2TypeTranslator) isFloat() bool {
	return t.property.Format == "float" || t.property.Format == "double"
}

func (t *openapiv2TypeTranslator) isArray() bool {
	return len(t.property.Type) > 0 && t.property.Type[0] == "array"
}

func (t *openapiv2TypeTranslator) isObject() bool {
	return len(t.property.Type) > 0 && t.property.Type[0] == "object"
}

func (t *openapiv2TypeTranslator) hasUnknownFields() bool {
	if t.property.Extensions != nil {
		_, ok := t.property.Extensions.Get("x-kubernetes-preserve-unknown-fields")
		return ok
	}
	return false
}

func (t *openapiv2TypeTranslator) hasProperties() bool {
	return t.property.Properties != nil && t.property.Properties.Len() > 0
}

func (t *openapiv2TypeTranslator) hasOneOf() bool {
	return len(t.property.OneOf) > 0
}

func (t *openapiv2TypeTranslator) isOneOfArray() bool {
	for _, oneOf := range t.property.OneOf {
		if oneOf.Schema().Type[0] == "array" {
			return true
		}
	}
	return false
}

func (t *openapiv2TypeTranslator) isOneOfBoolean() bool {
	for _, oneOf := range t.property.OneOf {
		if oneOf.Schema().Type[0] == "boolean" {
			return true
		}
	}
	return false
}

func (t *openapiv2TypeTranslator) isObjectWithAdditionalStringProperties() bool {
	return len(t.property.Type) > 0 && t.property.Type[0] == "object" &&
		t.property.AdditionalProperties != nil &&
		t.property.AdditionalProperties.IsA() &&
		t.property.AdditionalProperties.A.Schema().Type[0] == "string"
}

func (t *openapiv2TypeTranslator) isObjectWithAdditionalObjectProperties() bool {
	return len(t.property.Type) > 0 && t.property.Type[0] == "object" &&
		t.property.AdditionalProperties != nil &&
		t.property.AdditionalProperties.IsA() &&
		t.property.AdditionalProperties.A.Schema().Type[0] == "object"
}

func (t *openapiv2TypeTranslator) isObjectWithAdditionalArrayProperties() bool {
	return len(t.property.Type) > 0 && t.property.Type[0] == "object" &&
		t.property.AdditionalProperties != nil &&
		t.property.AdditionalProperties.IsA() &&
		t.property.AdditionalProperties.A.Schema().Type[0] == "array"
}

func (t *openapiv2TypeTranslator) isArrayWithObjectItems() bool {
	return len(t.property.Type) > 0 && t.property.Type[0] == "array" &&
		t.property.Items != nil &&
		t.property.Items.IsA() &&
		t.property.Items.A.Schema().Type[0] == "object"
}

func (t *openapiv2TypeTranslator) additionalPropertiesHaveStringItems() bool {
	return t.property.AdditionalProperties.A.Schema().Items != nil &&
		t.property.AdditionalProperties.A.Schema().Items.IsA() &&
		t.property.AdditionalProperties.A.Schema().Items.A.Schema().Type[0] == "string"
}

func (t *openapiv2TypeTranslator) additionalPropertiesHaveProperties() bool {
	return t.property.AdditionalProperties != nil &&
		t.property.AdditionalProperties.IsA() &&
		t.property.AdditionalProperties.A != nil &&
		t.property.AdditionalProperties.A.Schema() != nil &&
		t.property.AdditionalProperties.A.Schema().Properties != nil &&
		t.property.AdditionalProperties.A.Schema().Properties.Len() > 0
}

func (t *openapiv2TypeTranslator) additionalPropertiesHaveUnknownFields() bool {
	if t.property.AdditionalProperties != nil &&
		t.property.AdditionalProperties.IsA() &&
		t.property.AdditionalProperties.A != nil &&
		t.property.AdditionalProperties.A.Schema() != nil &&
		t.property.AdditionalProperties.A.Schema().Extensions != nil {
		_, ok := t.property.AdditionalProperties.A.Schema().Extensions.Get("x-kubernetes-preserve-unknown-fields")
		return ok
	}
	return false
}

func (t *openapiv2TypeTranslator) additionalPropertiesHaveAdditionalStringProperties() bool {
	return t.property.AdditionalProperties.A.Schema().AdditionalProperties != nil &&
		t.property.AdditionalProperties.A.Schema().AdditionalProperties.IsA() &&
		t.property.AdditionalProperties.A.Schema().AdditionalProperties.A.Schema().Type[0] == "string"
}

func (t *openapiv2TypeTranslator) additionalPropertiesHaveAdditionalArrayProperties() bool {
	return t.property.AdditionalProperties.A.Schema().AdditionalProperties != nil &&
		t.property.AdditionalProperties.A.Schema().AdditionalProperties.IsA() &&
		t.property.AdditionalProperties.A.Schema().AdditionalProperties.A.Schema().Type[0] == "array"
}

func (t *openapiv2TypeTranslator) additionalPropertiesHaveAdditionalPropertiesWithStringItems() bool {
	return t.property.AdditionalProperties != nil &&
		t.property.AdditionalProperties.IsA() &&
		t.property.AdditionalProperties.A != nil &&
		t.property.AdditionalProperties.A.Schema() != nil &&
		t.property.AdditionalProperties.A.Schema().AdditionalProperties != nil &&
		t.property.AdditionalProperties.A.Schema().AdditionalProperties.IsA() &&
		t.property.AdditionalProperties.A.Schema().AdditionalProperties.A != nil &&
		t.property.AdditionalProperties.A.Schema().AdditionalProperties.A.Schema() != nil &&
		t.property.AdditionalProperties.A.Schema().AdditionalProperties.A.Schema().Items != nil &&
		t.property.AdditionalProperties.A.Schema().AdditionalProperties.A.Schema().Items.IsA() &&
		t.property.AdditionalProperties.A.Schema().AdditionalProperties.A.Schema().Items.A != nil &&
		t.property.AdditionalProperties.A.Schema().AdditionalProperties.A.Schema().Items.A.Schema() != nil &&
		t.property.AdditionalProperties.A.Schema().AdditionalProperties.A.Schema().Items.A.Schema().Type[0] == "string"
}

func (t *openapiv2TypeTranslator) itemsHaveUnknownFields() bool {
	if t.property.Items != nil && t.property.Items.IsA() {
		if t.property.Items.A.Schema().Extensions != nil {
			_, ok := t.property.Items.A.Schema().Extensions.Get("x-kubernetes-preserve-unknown-fields")
			return ok
		}
	}
	return false
}

func (t *openapiv2TypeTranslator) itemsHaveAdditionalStringProperties() bool {
	return t.property.Items.A.Schema().AdditionalProperties != nil &&
		t.property.Items.A.Schema().AdditionalProperties.IsA() &&
		t.property.Items.A.Schema().AdditionalProperties.A.Schema().Type[0] == "string"
}
