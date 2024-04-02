/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package customtypes

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"math/big"
	yaml "sigs.k8s.io/yaml/goyaml.v3"
	"testing"
)

func TestIntOrStringTypeTerraformType(t *testing.T) {
	result := IntOrStringType{}.TerraformType(context.Background())
	if diff := cmp.Diff(result, tftypes.DynamicPseudoType); diff != "" {
		t.Errorf("unexpected result (+expected, -got): %s", diff)
	}
}

func TestIntOrStringTypeValueFromTerraform(t *testing.T) {
	type testCase struct {
		receiver IntOrStringType
		input    tftypes.Value
		expected attr.Value
	}
	tests := map[string]testCase{
		"string": {
			receiver: IntOrStringType{},
			input:    tftypes.NewValue(tftypes.String, "hello"),
			expected: IntOrStringValue{
				DynamicValue: basetypes.NewDynamicValue(basetypes.NewStringValue("hello")),
			},
		},
		"number": {
			receiver: IntOrStringType{},
			input:    tftypes.NewValue(tftypes.Number, 123),
			expected: IntOrStringValue{
				DynamicValue: basetypes.NewDynamicValue(basetypes.NewNumberValue(big.NewFloat(123))),
			},
		},
		"nil": {
			receiver: IntOrStringType{},
			input:    tftypes.NewValue(nil, nil),
			expected: IntOrStringValue{
				DynamicValue: basetypes.NewDynamicNull(),
			},
		},
		"unknown": {
			receiver: IntOrStringType{},
			input:    tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
			expected: IntOrStringValue{
				DynamicValue: basetypes.NewDynamicUnknown(),
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := test.receiver.ValueFromTerraform(context.Background(), test.input)
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
				return
			}
			if diff := cmp.Diff(test.expected, got); diff != "" {
				t.Errorf("unexpected result (-expected, +got): %s", diff)
			}
			if test.expected != nil && test.expected.IsNull() != test.input.IsNull() {
				t.Errorf("Expected null-ness match: expected %t, got %t", test.expected.IsNull(), test.input.IsNull())
			}
			if test.expected != nil && test.expected.IsUnknown() != !test.input.IsKnown() {
				t.Errorf("Expected unknown-ness match: expected %t, got %t", test.expected.IsUnknown(), !test.input.IsKnown())
			}
		})
	}
}

func TestIntOrStringValueToTerraformValue(t *testing.T) {
	type testCase struct {
		receiver IntOrStringValue
		expected tftypes.Value
	}
	tests := map[string]testCase{
		"string": {
			receiver: IntOrStringValue{
				DynamicValue: basetypes.NewDynamicValue(basetypes.NewStringValue("hello")),
			},
			expected: tftypes.NewValue(tftypes.String, "hello"),
		},
		"number": {
			receiver: IntOrStringValue{
				DynamicValue: basetypes.NewDynamicValue(basetypes.NewNumberValue(big.NewFloat(123))),
			},
			expected: tftypes.NewValue(tftypes.Number, 123),
		},
		"unknown": {
			receiver: IntOrStringValue{
				DynamicValue: basetypes.NewDynamicUnknown(),
			},
			expected: tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
		},
		"null": {
			receiver: IntOrStringValue{
				DynamicValue: basetypes.NewDynamicNull(),
			},
			expected: tftypes.NewValue(tftypes.DynamicPseudoType, nil),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := test.receiver.ToTerraformValue(context.Background())
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
				return
			}
			if diff := cmp.Diff(test.expected, got); diff != "" {
				t.Errorf("Unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestIntOrStringValueYamlMarshaller(t *testing.T) {
	type testCase struct {
		value    attr.Value
		expected string
	}
	tests := map[string]testCase{
		"int": {
			value: IntOrStringValue{
				DynamicValue: basetypes.NewDynamicValue(basetypes.NewNumberValue(big.NewFloat(123))),
			},
			expected: "123",
		},
		"string": {
			value: IntOrStringValue{
				DynamicValue: basetypes.NewDynamicValue(basetypes.NewStringValue("hello")),
			},
			expected: "hello",
		},
		"null": {
			value: IntOrStringValue{
				DynamicValue: basetypes.NewDynamicNull(),
			},
			expected: "null",
		},
		"unknown": {
			value: IntOrStringValue{
				DynamicValue: basetypes.NewDynamicUnknown(),
			},
			expected: "null",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			marshal, err := yaml.Marshal(test.value)
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
				return
			}
			if diff := cmp.Diff(test.expected+"\n", string(marshal)); diff != "" {
				t.Errorf("unexpected result (-expected, +got): %s", diff)
			}
		})
	}
}
