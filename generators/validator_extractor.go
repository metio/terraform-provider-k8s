//go:build generators

/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package main

import "strings"

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
	case "schema.StringAttribute":
		return "validator.String"
	case "schema.Int64Attribute":
		return "validator.Int64"
	case "schema.MapAttribute":
		return "validator.Map"
	default:
		return "UNKNOWN"
	}
}
