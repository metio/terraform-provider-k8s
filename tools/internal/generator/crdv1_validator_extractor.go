/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"fmt"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"strconv"
)

type crdv1ValidatorExtractor struct {
	property *apiextensionsv1.JSONSchemaProps
	imports  *AdditionalImports
}

var _ validatorExtractor = (*crdv1ValidatorExtractor)(nil)

func (v *crdv1ValidatorExtractor) integerWithMinimum() string {
	if v.property.Type == "integer" && v.property.Minimum != nil {
		v.imports.Int64Validator = true
		return fmt.Sprintf("int64validator.AtLeast(%v)", crdv1MinValue(v.property))
	}
	return ""
}

func (v *crdv1ValidatorExtractor) integerWithMaximum() string {
	if v.property.Type == "integer" && v.property.Maximum != nil {
		v.imports.Int64Validator = true
		return fmt.Sprintf("int64validator.AtMost(%v)", crdv1MaxValue(v.property))
	}
	return ""
}

func (v *crdv1ValidatorExtractor) integerWithEnums() string {
	if v.property.Type == "integer" && len(v.property.Enum) > 0 {
		v.imports.Int64Validator = true
		enums := crdv1IntEnums(v.property.Enum)
		return fmt.Sprintf("int64validator.OneOf(%v)", concatEnums(enums))
	}
	return ""
}

func (v *crdv1ValidatorExtractor) numberWithMinimum() string {
	if v.property.Type == "number" && v.property.Minimum != nil {
		v.imports.Float64Validator = true
		return fmt.Sprintf("float64validator.AtLeast(%v)", crdv1MinValue(v.property))
	}
	return ""
}

func (v *crdv1ValidatorExtractor) numberWithMaximum() string {
	if v.property.Type == "number" && v.property.Maximum != nil {
		v.imports.Float64Validator = true
		return fmt.Sprintf("float64validator.AtMost(%v)", crdv1MaxValue(v.property))
	}
	return ""
}

func (v *crdv1ValidatorExtractor) numberWithEnums() string {
	if v.property.Type == "number" && len(v.property.Enum) > 0 {
		v.imports.Float64Validator = true
		enums := crdv1FloatEnums(v.property.Enum)
		return fmt.Sprintf("float64validator.OneOf(%v)", concatEnums(enums))
	}
	return ""
}

func (v *crdv1ValidatorExtractor) stringWithByteFormat() string {
	if v.property.Type == "string" && v.property.Format == "byte" {
		return "validators.Base64Validator()"
	}
	return ""
}

func (v *crdv1ValidatorExtractor) stringWithDateTimeFormat() string {
	if v.property.Type == "string" && v.property.Format == "date-time" {
		return "validators.DateTime64Validator()"
	}
	return ""
}

func (v *crdv1ValidatorExtractor) stringWithMinimumLength() string {
	if v.property.Type == "string" && v.property.MinLength != nil {
		v.imports.StringValidator = true
		return fmt.Sprintf("stringvalidator.LengthAtLeast(%v)", *v.property.MinLength)
	}
	return ""
}

func (v *crdv1ValidatorExtractor) stringWithMaximumLength() string {
	if v.property.Type == "string" && v.property.MaxLength != nil {
		v.imports.StringValidator = true
		return fmt.Sprintf("stringvalidator.LengthAtMost(%v)", *v.property.MaxLength)
	}
	return ""
}

func (v *crdv1ValidatorExtractor) stringWithEnums() string {
	if v.property.Type == "string" && len(v.property.Enum) > 0 {
		v.imports.StringValidator = true
		enums := crdv1StringEnums(v.property.Enum)
		return fmt.Sprintf("stringvalidator.OneOf(%s)", concatEnums(enums))
	}
	return ""
}

func (v *crdv1ValidatorExtractor) stringWithPattern() string {
	if v.property.Type == "string" && v.property.Pattern != "" {
		v.imports.Regexp = true
		v.imports.StringValidator = true
		return fmt.Sprintf(`stringvalidator.RegexMatches(regexp.MustCompile(%c%s%c), "")`, '`', v.property.Pattern, '`')
	}
	return ""
}

func crdv1StringEnums(enums []apiextensionsv1.JSON) []string {
	var values []string

	for _, val := range enums {
		if str := string(val.Raw); str != "" {
			values = append(values, str)
		}
	}

	return values
}

func crdv1IntEnums(enums []apiextensionsv1.JSON) []int64 {
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

func crdv1FloatEnums(enums []apiextensionsv1.JSON) []float64 {
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

func crdv1MaxValue(prop *apiextensionsv1.JSONSchemaProps) float64 {
	max := *prop.Maximum
	if prop.ExclusiveMaximum {
		max = max - 1
	}
	return max
}

func crdv1MinValue(prop *apiextensionsv1.JSONSchemaProps) float64 {
	min := *prop.Minimum
	if prop.ExclusiveMinimum {
		min = min + 1
	}
	return min
}
