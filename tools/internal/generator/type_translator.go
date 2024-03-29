/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

type typeTranslator interface {
	hasNoType() bool
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

func translateTypeWith(translator typeTranslator) (attributeType string, valueType string, elementType string, goType string) {
	if translator.isBoolean() {
		//attributeType = "types.BoolType"
		attributeType = "schema.BoolAttribute"
		elementType = ""
		valueType = "types.Bool"
		goType = "bool"
		return
	}
	if translator.isInteger() {
		//attributeType = "types.Int64Type"
		attributeType = "schema.Int64Attribute"
		elementType = ""
		valueType = "types.Int64"
		goType = "int64"
		return
	}
	if translator.isString() {
		//attributeType = "types.StringType"
		attributeType = "schema.StringAttribute"
		elementType = ""
		valueType = "types.String"
		goType = "string"
		return
	}
	if translator.isIntOrString() {
		//attributeType = "types.StringType"
		attributeType = "schema.StringAttribute"
		elementType = ""
		valueType = "types.String"
		goType = "string"
		return
	}
	if translator.isNumber() {
		if translator.isFloat() {
			//attributeType = "types.Float64Type"
			attributeType = "schema.Float64Attribute"
			elementType = ""
			valueType = "types.Float64"
			goType = "float64"
			return
		}
		//attributeType = "types.NumberType"
		//attributeType = "schema.NumberAttribute" // TODO: add support for big.Float
		attributeType = "schema.Float64Attribute"
		elementType = ""
		valueType = "types.Float64"
		goType = "float64"
		return
	}
	if translator.hasUnknownFields() {
		if translator.hasProperties() {
			//attributeType = "types.ObjectType"
			attributeType = "schema.SingleNestedAttribute"
			elementType = ""
			valueType = "types.Object"
			goType = "struct"
			return
		}
		//attributeType = "types.MapType{ElemType: types.StringType}"
		attributeType = "schema.MapAttribute"
		elementType = "types.StringType"
		valueType = "types.Map"
		goType = "map[string]string"
		return
	}
	if translator.hasNoType() && translator.hasProperties() {
		attributeType = "schema.SingleNestedAttribute"
		elementType = ""
		valueType = "types.Object"
		goType = "struct"
		return
	}
	if translator.hasOneOf() {
		if translator.isOneOfArray() {
			//attributeType = "types.ListType{ElemType: types.StringType}"
			attributeType = "schema.ListAttribute"
			elementType = "types.StringType"
			valueType = "types.List"
			goType = "[]string"
			return
		}
		if translator.isOneOfBoolean() {
			//attributeType = "types.BoolType"
			attributeType = "schema.BoolAttribute"
			elementType = ""
			valueType = "types.Bool"
			goType = "bool"
			return
		}
	}
	if translator.isObjectWithAdditionalStringProperties() {
		//attributeType = "types.MapType{ElemType: types.StringType}"
		attributeType = "schema.MapAttribute"
		elementType = "types.StringType"
		valueType = "types.Map"
		goType = "map[string]string"
		return
	}
	if translator.isObjectWithAdditionalObjectProperties() {
		if translator.additionalPropertiesHaveAdditionalStringProperties() {
			//attributeType = "types.MapType{ElemType: types.MapType{ElemType: types.StringType}}"
			attributeType = "schema.MapAttribute"
			elementType = "types.MapType{ElemType: types.StringType}"
			valueType = "types.Map"
			goType = "map[string]map[string]string"
			return
		}
		if translator.additionalPropertiesHaveAdditionalArrayProperties() {
			if translator.additionalPropertiesHaveAdditionalPropertiesWithStringItems() {
				//attributeType = "types.MapType{ElemType: types.MapType{ElemType: types.ListType{ElemType: types.StringType}}}"
				attributeType = "schema.MapAttribute"
				elementType = "types.MapType{ElemType: types.ListType{ElemType: types.StringType}}"
				valueType = "types.Map"
				goType = "map[string]map[string][]string"
				return
			}
		}
		if translator.additionalPropertiesHaveUnknownFields() {
			if !translator.additionalPropertiesHaveProperties() {
				//attributeType = "types.MapType{ElemType: types.StringType}"
				attributeType = "schema.MapAttribute"
				elementType = "types.StringType"
				valueType = "types.Map"
				goType = "map[string]string"
				return
			}
		}
		//attributeType = "types.ObjectType"
		attributeType = "schema.SingleNestedAttribute"
		elementType = ""
		valueType = "types.Object"
		goType = "struct"
		return
	}
	if translator.isObjectWithAdditionalArrayProperties() {
		if translator.additionalPropertiesHaveStringItems() {
			//attributeType = "types.MapType{ElemType: types.ListType{ElemType: types.StringType}}"
			attributeType = "schema.MapAttribute"
			elementType = "types.ListType{ElemType: types.StringType}"
			valueType = "types.Map"
			goType = "map[string][]string"
			return
		}
	}
	if translator.isArray() {
		if translator.isArrayWithObjectItems() {
			if translator.itemsHaveUnknownFields() {
				//attributeType = "types.ListType{ElemType: types.MapType{ElemType: types.StringType}}"
				attributeType = "schema.ListAttribute"
				elementType = "types.MapType{ElemType: types.StringType}"
				valueType = "types.List"
				goType = "[]map[string]string"
				return
			}
			if translator.itemsHaveAdditionalStringProperties() {
				//attributeType = "types.ListType{ElemType: types.MapType{ElemType: types.StringType}}"
				attributeType = "schema.ListAttribute"
				elementType = "types.MapType{ElemType: types.StringType}"
				valueType = "types.List"
				goType = "[]map[string]string"
				return
			}
			//attributeType = "types.ListType{ElemType: types.ObjectType}"
			attributeType = "schema.ListNestedAttribute"
			elementType = ""
			valueType = "types.List"
			goType = "[]struct"
			return
		}
		//attributeType = "types.ListType{ElemType: types.StringType}"
		attributeType = "schema.ListAttribute"
		elementType = "types.StringType"
		valueType = "types.List"
		goType = "[]string"
		return
	}

	if translator.isObject() {
		if translator.hasProperties() {
			//attributeType = "types.ObjectType"
			attributeType = "schema.SingleNestedAttribute"
			elementType = ""
			valueType = "types.Object"
			goType = "struct"
			return
		}
		//attributeType = "types.MapType{ElemType: types.StringType}"
		attributeType = "schema.MapAttribute"
		elementType = "types.StringType"
		valueType = "types.Map"
		goType = "map[string]string"
		return
	}

	attributeType = "UNKNOWN"
	elementType = "UNKNOWN"
	valueType = "UNKNOWN"
	goType = "UNKNOWN"
	return
}
