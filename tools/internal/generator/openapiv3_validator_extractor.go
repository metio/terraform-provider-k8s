/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
)

var _ validatorExtractor = (*openapiv3ValidatorExtractor)(nil)

type openapiv3ValidatorExtractor struct {
	property *openapi3.Schema
	imports  *AdditionalImports
}

func (v *openapiv3ValidatorExtractor) integerWithMinimum() string {
	if v.property.Type == "integer" && v.property.Min != nil {
		v.imports.Int64Validator = true
		min := *v.property.Min
		if v.property.ExclusiveMin {
			min = min + float64(1)
		}
		return fmt.Sprintf("int64validator.AtLeast(%v)", min)
	}
	return ""
}

func (v *openapiv3ValidatorExtractor) integerWithMaximum() string {
	if v.property.Type == "integer" && v.property.Max != nil {
		v.imports.Int64Validator = true
		max := *v.property.Max
		if v.property.ExclusiveMax {
			max = max - float64(1)
		}
		return fmt.Sprintf("int64validator.AtMost(%v)", max)
	}
	return ""
}

func (v *openapiv3ValidatorExtractor) integerWithEnums() string {
	if v.property.Type == "integer" && len(v.property.Enum) > 0 {
		v.imports.Int64Validator = true
		enums := openapiIntEnums(v.property.Enum)
		return fmt.Sprintf("int64validator.OneOf(%v)", concatEnums(enums))
	}
	return ""
}

func (v *openapiv3ValidatorExtractor) numberWithMinimum() string {
	if v.property.Type == "number" && v.property.Min != nil {
		v.imports.Float64Validator = true
		min := *v.property.Min
		if v.property.ExclusiveMin {
			min = min + float64(1)
		}
		return fmt.Sprintf("float64validator.AtLeast(%v)", min)
	}
	return ""
}

func (v *openapiv3ValidatorExtractor) numberWithMaximum() string {
	if v.property.Type == "number" && v.property.Max != nil {
		v.imports.Float64Validator = true
		max := *v.property.Max
		if v.property.ExclusiveMax {
			max = max - float64(1)
		}
		return fmt.Sprintf("float64validator.AtMost(%v)", max)
	}
	return ""
}

func (v *openapiv3ValidatorExtractor) numberWithEnums() string {
	if v.property.Type == "number" && len(v.property.Enum) > 0 {
		v.imports.Float64Validator = true
		enums := openapiFloatEnums(v.property.Enum)
		return fmt.Sprintf("float64validator.OneOf(%v)", concatEnums(enums))
	}
	return ""
}

func (v *openapiv3ValidatorExtractor) stringWithByteFormat() string {
	if v.property.Type == "string" && v.property.Format == "byte" {
		return "validators.Base64Validator()"
	}
	return ""
}

func (v *openapiv3ValidatorExtractor) stringWithDateTimeFormat() string {
	if v.property.Type == "string" && v.property.Format == "date-time" {
		return "validators.DateTime64Validator()"
	}
	return ""
}

func (v *openapiv3ValidatorExtractor) stringWithMinimumLength() string {
	if v.property.Type == "string" && v.property.MinLength != 0 {
		v.imports.StringValidator = true
		return fmt.Sprintf("stringvalidator.LengthAtLeast(%v)", v.property.MinLength)
	}
	return ""
}

func (v *openapiv3ValidatorExtractor) stringWithMaximumLength() string {
	if v.property.Type == "string" && v.property.MaxLength != nil {
		v.imports.StringValidator = true
		return fmt.Sprintf("stringvalidator.LengthAtMost(%v)", *v.property.MaxLength)
	}
	return ""
}

func (v *openapiv3ValidatorExtractor) stringWithEnums() string {
	if v.property.Type == "string" && len(v.property.Enum) > 0 {
		v.imports.StringValidator = true
		enums := openapiStringEnums(v.property.Enum)
		return fmt.Sprintf("stringvalidator.OneOf(%s)", concatEnums(enums))
	}
	return ""
}

func (v *openapiv3ValidatorExtractor) stringWithPattern() string {
	if v.property.Type == "string" && v.property.Pattern != "" {
		v.imports.Regexp = true
		v.imports.StringValidator = true
		return fmt.Sprintf(`stringvalidator.RegexMatches(regexp.MustCompile(%c%s%c), "")`, '`', v.property.Pattern, '`')
	}
	return ""
}

func openapiIntEnums(enums []interface{}) []int64 {
	var values []int64

	for _, val := range enums {
		if number, ok := val.(int64); ok {
			values = append(values, number)
		}
	}

	return values
}

func openapiFloatEnums(enums []interface{}) []float64 {
	var values []float64

	for _, val := range enums {
		if number, ok := val.(float64); ok {
			values = append(values, number)
		}
	}

	return values
}

func openapiStringEnums(enums []interface{}) []string {
	var values []string

	for _, val := range enums {
		if str, ok := val.(string); ok {
			values = append(values, str)
		}
	}

	return values
}
