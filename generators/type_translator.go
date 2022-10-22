//go:build generators

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

type typeTranslator interface {
	isIntOrString() bool
	isBoolean() bool
	isString() bool
	isInteger() bool
	isNumber() bool
	isFloat() bool
	isArray() bool
	isObject() bool
	hasUnknownFields() bool
	hasProperties() bool
	hasOneOf() bool
	isOneOfArray() bool
	isOneOfBoolean() bool
	isObjectWithAdditionalStringProperties() bool
	isObjectWithAdditionalObjectProperties() bool
	isObjectWithAdditionalArrayProperties() bool
	isArrayWithObjectItems() bool
	additionalPropertiesHaveStringItems() bool
	additionalPropertiesHaveProperties() bool
	additionalPropertiesHaveUnknownFields() bool
	additionalPropertiesHaveAdditionalStringProperties() bool
	additionalPropertiesHaveAdditionalArrayProperties() bool
	additionalPropertiesHaveAdditionalPropertiesWithStringItems() bool
	itemsHaveUnknownFields() bool
	itemsHaveAdditionalStringProperties() bool
}

func translateTypeWith(translator typeTranslator) (attributeType string, valueType string, goType string) {
	if translator.isIntOrString() {
		attributeType = "utilities.IntOrStringType{}"
		valueType = "utilities.IntOrString"
		goType = "utilities.IntOrString"
		return
	}
	if translator.hasUnknownFields() {
		if translator.hasProperties() {
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
	if translator.hasOneOf() {
		if translator.isOneOfArray() {
			attributeType = "types.ListType{ElemType: types.StringType}"
			valueType = "types.List"
			goType = "[]string"
			return
		}
		if translator.isOneOfBoolean() {
			attributeType = "types.BoolType"
			valueType = "types.Bool"
			goType = "bool"
			return
		}
	}
	if translator.isObjectWithAdditionalStringProperties() {
		attributeType = "types.MapType{ElemType: types.StringType}"
		valueType = "types.Map"
		goType = "map[string]string"
		return
	}
	if translator.isObjectWithAdditionalObjectProperties() {
		if translator.additionalPropertiesHaveAdditionalStringProperties() {
			attributeType = "types.MapType{ElemType: types.MapType{ElemType: types.StringType}}"
			valueType = "types.Map"
			goType = "map[string]map[string]string"
			return
		}
		if translator.additionalPropertiesHaveAdditionalArrayProperties() {
			if translator.additionalPropertiesHaveAdditionalPropertiesWithStringItems() {
				attributeType = "types.MapType{ElemType: types.MapType{ElemType: types.ListType{ElemType: types.StringType}}}"
				valueType = "types.Map"
				goType = "map[string]map[string][]string"
				return
			}
		}
		if translator.additionalPropertiesHaveUnknownFields() {
			if !translator.additionalPropertiesHaveProperties() {
				attributeType = "utilities.DynamicType{}"
				valueType = "utilities.Dynamic"
				goType = "utilities.Dynamic"
				return
			}
		}
		attributeType = "types.ObjectType"
		valueType = "types.Object"
		goType = "struct"
		return
	}
	if translator.isObjectWithAdditionalArrayProperties() {
		if translator.additionalPropertiesHaveStringItems() {
			attributeType = "types.MapType{ElemType: types.ListType{ElemType: types.StringType}}"
			valueType = "types.Map"
			goType = "map[string][]string"
			return
		}
	}
	if translator.isArrayWithObjectItems() {
		if translator.itemsHaveUnknownFields() {
			attributeType = "types.ListType{ElemType: types.MapType{ElemType: types.StringType}}"
			valueType = "types.List"
			goType = "[]map[string]string"
			return
		}
		if translator.itemsHaveAdditionalStringProperties() {
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
	if translator.isObject() {
		if translator.hasProperties() {
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
	if translator.isArray() {
		attributeType = "types.ListType{ElemType: types.StringType}"
		valueType = "types.List"
		goType = "[]string"
		return
	}
	if translator.isNumber() {
		if translator.isFloat() {
			attributeType = "types.Float64Type"
			valueType = "types.Float64"
			goType = "float64"
			return
		}
		attributeType = "utilities.DynamicNumberType{}"
		valueType = "utilities.DynamicNumber"
		goType = "utilities.DynamicNumber"
		return
	}
	if translator.isInteger() {
		attributeType = "types.Int64Type"
		valueType = "types.Int64"
		goType = "int64"
		return
	}
	if translator.isString() {
		attributeType = "types.StringType"
		valueType = "types.String"
		goType = "string"
		return
	}
	if translator.isBoolean() {
		attributeType = "types.BoolType"
		valueType = "types.Bool"
		goType = "bool"
		return
	}

	attributeType = "UNKNOWN"
	valueType = "UNKNOWN"
	goType = "UNKNOWN"
	return
}
