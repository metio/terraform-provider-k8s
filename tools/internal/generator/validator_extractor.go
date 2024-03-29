/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

import (
	"fmt"
	"slices"
	"strings"
)

type validatorExtractor interface {
	integerWithMinimum() string
	integerWithMaximum() string
	integerWithEnums() string
	numberWithMinimum() string
	numberWithMaximum() string
	numberWithEnums() string
	stringWithByteFormat() string
	stringWithDateTimeFormat() string
	stringWithMinimumLength() string
	stringWithMaximumLength() string
	stringWithEnums() string
	stringWithPattern() string
}

func validatorsFor(validator validatorExtractor, terraformResourceName string, propPath string, imports *AdditionalImports) []string {
	validators := upstreamValidators(validator)
	validators = append(validators, customValidators(terraformResourceName, propPath, imports)...)
	return validators
}

func upstreamValidators(validator validatorExtractor) []string {
	validators := make([]string, 0)

	if val := validator.integerWithMinimum(); val != "" {
		validators = append(validators, val)
	}
	if val := validator.integerWithMaximum(); val != "" {
		validators = append(validators, val)
	}
	if val := validator.integerWithEnums(); val != "" {
		validators = append(validators, val)
	}
	if val := validator.numberWithMinimum(); val != "" {
		validators = append(validators, val)
	}
	if val := validator.numberWithMaximum(); val != "" {
		validators = append(validators, val)
	}
	if val := validator.numberWithEnums(); val != "" {
		validators = append(validators, val)
	}
	if val := validator.stringWithByteFormat(); val != "" {
		validators = append(validators, val)
	}
	if val := validator.stringWithDateTimeFormat(); val != "" {
		validators = append(validators, val)
	}
	if val := validator.stringWithMinimumLength(); val != "" {
		validators = append(validators, val)
	}
	if val := validator.stringWithMaximumLength(); val != "" {
		validators = append(validators, val)
	}
	if val := validator.stringWithEnums(); val != "" {
		validators = append(validators, val)
	}
	if val := validator.stringWithPattern(); val != "" {
		validators = append(validators, val)
	}

	return validators
}

func customValidators(terraformResourceName string, propPath string, imports *AdditionalImports) []string {
	validators := make([]string, 0)
	if cvs, ok := customValidations[terraformResourceName]; ok {
		if cv, ok := cvs[propPath]; ok {
			if usesValidatorPackage(cv, "int64validator") {
				imports.Int64Validator = true
			}
			if usesValidatorPackage(cv, "float64validator") {
				imports.Float64Validator = true
			}
			if usesValidatorPackage(cv, "stringvalidator") {
				imports.StringValidator = true
			}
			if usesValidatorPackage(cv, "boolvalidator") {
				imports.BoolValidator = true
			}
			validators = append(validators, cv...)
		}
	}
	return validators
}

func usesValidatorPackage(validators []string, pkg string) bool {
	for _, validator := range validators {
		if strings.Contains(validator, pkg) {
			return true
		}
	}
	return false
}

func mapAttributeTypeToValidatorsType(attributeType string) string {
	switch attributeType {
	case "schema.BoolAttribute":
		return "validator.Bool"
	case "schema.StringAttribute":
		return "validator.String"
	case "schema.Int64Attribute":
		return "validator.Int64"
	case "schema.Float64Attribute":
		return "validator.Float64"
	case "schema.NumberAttribute":
		return "validator.Number"
	case "schema.MapAttribute":
		return "validator.Map"
	case "schema.SingleNestedAttribute":
		return "validator.Object"
	case "schema.ListAttribute":
		return "validator.List"
	case "schema.ListNestedAttribute":
		return "validator.List"
	default:
		return "UNKNOWN"
	}
}

func mapAttributeTypeToValidatorsPackage(attributeType string) string {
	switch attributeType {
	case "schema.BoolAttribute":
		return "boolvalidator"
	case "schema.StringAttribute":
		return "stringvalidator"
	case "schema.Int64Attribute":
		return "int64validator"
	case "schema.Float64Attribute":
		return "float64validator"
	case "schema.NumberAttribute":
		return "numbervalidator"
	case "schema.MapAttribute":
		return "mapvalidator"
	case "schema.SingleNestedAttribute":
		return "objectvalidator"
	case "schema.ListAttribute":
		return "listvalidator"
	case "schema.ListNestedAttribute":
		return "listvalidator"
	default:
		return "UNKNOWN"
	}
}

func addValidatorImports(outer *Property, imports *AdditionalImports) {
	switch outer.ValidatorsPackage {
	case "boolvalidator":
		imports.BoolValidator = true
	case "listvalidator":
		imports.ListValidator = true
	case "objectvalidator":
		imports.ObjectValidator = true
	case "mapvalidator":
		imports.MapValidator = true
	}
}

func escapeRegexPattern(pattern string) string {
	splits := strings.Split(pattern, "`")
	splits = slices.DeleteFunc(splits, func(s string) bool {
		return s == ""
	})
	if strings.Contains(pattern, "`") {
		var sb strings.Builder
		if strings.HasPrefix(pattern, "`") {
			if len(splits) > 0 {
				sb.WriteString(fmt.Sprintf(`"%c"+`, '`'))
			} else {
				sb.WriteString(fmt.Sprintf(`"%c"`, '`'))
			}
		}
		for index, value := range splits {
			if index > 0 && splits[index-1] != "" {
				sb.WriteString(fmt.Sprintf(`+%c%s%c`, '`', value, '`'))
			} else {
				sb.WriteString(fmt.Sprintf(`%c%s%c`, '`', value, '`'))
			}
			if index < len(splits)-1 {
				sb.WriteString(fmt.Sprintf(`+"%c"`, '`'))
			}
		}
		if strings.HasSuffix(pattern, "`") {
			if len(splits) > 0 {
				sb.WriteString(fmt.Sprintf(`+"%c"`, '`'))
			}
		}
		return sb.String()
	} else {
		return fmt.Sprintf("%c%s%c", '`', pattern, '`')
	}
}
