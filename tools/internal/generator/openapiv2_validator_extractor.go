/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"fmt"
	"github.com/pb33f/libopenapi/datamodel/high/base"
	"gopkg.in/yaml.v3"
	"strconv"
)

var _ validatorExtractor = (*openapiv2ValidatorExtractor)(nil)

type openapiv2ValidatorExtractor struct {
	property *base.Schema
	imports  *AdditionalImports
}

func (v *openapiv2ValidatorExtractor) integerWithMinimum() string {
	if v.property.Type[0] == "integer" && v.property.Minimum != nil {
		v.imports.Int64Validator = true
		minimum := *v.property.Minimum
		if v.property.ExclusiveMinimum.IsA() && v.property.ExclusiveMinimum.A {
			minimum = minimum + float64(1)
		}
		return fmt.Sprintf("int64validator.AtLeast(%v)", minimum)
	}
	return ""
}

func (v *openapiv2ValidatorExtractor) integerWithMaximum() string {
	if v.property.Type[0] == "integer" && v.property.Maximum != nil {
		v.imports.Int64Validator = true
		maximum := *v.property.Maximum
		if v.property.ExclusiveMaximum.IsA() && v.property.ExclusiveMaximum.A {
			maximum = maximum - float64(1)
		}
		return fmt.Sprintf("int64validator.AtMost(%v)", maximum)
	}
	return ""
}

func (v *openapiv2ValidatorExtractor) integerWithEnums() string {
	if v.property.Type[0] == "integer" && len(v.property.Enum) > 0 {
		v.imports.Int64Validator = true
		enums := openapiv2IntEnums(v.property.Enum)
		return fmt.Sprintf("int64validator.OneOf(%v)", concatEnums(enums))
	}
	return ""
}

func (v *openapiv2ValidatorExtractor) numberWithMinimum() string {
	if v.property.Type[0] == "number" && v.property.Minimum != nil {
		v.imports.Float64Validator = true
		minimum := *v.property.Minimum
		if v.property.ExclusiveMinimum.IsA() && v.property.ExclusiveMinimum.A {
			minimum = minimum + float64(1)
		}
		return fmt.Sprintf("float64validator.AtLeast(%v)", minimum)
	}
	return ""
}

func (v *openapiv2ValidatorExtractor) numberWithMaximum() string {
	if v.property.Type[0] == "number" && v.property.Maximum != nil {
		v.imports.Float64Validator = true
		maximum := *v.property.Maximum
		if v.property.ExclusiveMaximum.IsA() && v.property.ExclusiveMaximum.A {
			maximum = maximum - float64(1)
		}
		return fmt.Sprintf("float64validator.AtMost(%v)", maximum)
	}
	return ""
}

func (v *openapiv2ValidatorExtractor) numberWithEnums() string {
	if v.property.Type[0] == "number" && len(v.property.Enum) > 0 {
		v.imports.Float64Validator = true
		enums := openapiv2FloatEnums(v.property.Enum)
		return fmt.Sprintf("float64validator.OneOf(%v)", concatEnums(enums))
	}
	return ""
}

func (v *openapiv2ValidatorExtractor) stringWithByteFormat() string {
	if v.property.Type[0] == "string" && v.property.Format == "byte" {
		return "validators.Base64Validator()"
	}
	return ""
}

func (v *openapiv2ValidatorExtractor) stringWithDateTimeFormat() string {
	if v.property.Type[0] == "string" && v.property.Format == "date-time" {
		return "validators.DateTime64Validator()"
	}
	return ""
}

func (v *openapiv2ValidatorExtractor) stringWithMinimumLength() string {
	if v.property.Type[0] == "string" && v.property.MinLength != nil {
		v.imports.StringValidator = true
		return fmt.Sprintf("stringvalidator.LengthAtLeast(%v)", v.property.MinLength)
	}
	return ""
}

func (v *openapiv2ValidatorExtractor) stringWithMaximumLength() string {
	if v.property.Type[0] == "string" && v.property.MaxLength != nil {
		v.imports.StringValidator = true
		return fmt.Sprintf("stringvalidator.LengthAtMost(%v)", *v.property.MaxLength)
	}
	return ""
}

func (v *openapiv2ValidatorExtractor) stringWithEnums() string {
	if v.property.Type[0] == "string" && len(v.property.Enum) > 0 {
		v.imports.StringValidator = true
		enums := openapiv2StringEnums(v.property.Enum)
		return fmt.Sprintf("stringvalidator.OneOf(%s)", concatEnums(enums))
	}
	return ""
}

func (v *openapiv2ValidatorExtractor) stringWithPattern() string {
	if v.property.Type[0] == "string" && v.property.Pattern != "" {
		v.imports.Regexp = true
		v.imports.StringValidator = true
		return fmt.Sprintf(`stringvalidator.RegexMatches(regexp.MustCompile(%s), "")`, escapeRegexPattern(v.property.Pattern))
	}
	return ""
}

func openapiv2IntEnums(enums []*yaml.Node) []int64 {
	var values []int64

	for _, val := range enums {
		if number, err := strconv.ParseInt(val.Value, 10, 64); err != nil {
			values = append(values, number)
		}
	}

	return values
}

func openapiv2FloatEnums(enums []*yaml.Node) []float64 {
	var values []float64

	for _, val := range enums {
		if number, err := strconv.ParseFloat(val.Value, 10); err != nil {
			values = append(values, number)
		}
	}

	return values
}

func openapiv2StringEnums(enums []*yaml.Node) []string {
	var values []string

	for _, val := range enums {
		values = append(values, val.Value)
	}

	return values
}
