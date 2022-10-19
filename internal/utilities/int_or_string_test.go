/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package utilities

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestIntOrStringTypeTerraformType(t *testing.T) {
	t.Parallel()
	result := IntOrStringType{}.TerraformType(context.Background())
	if diff := cmp.Diff(result, tftypes.DynamicPseudoType); diff != "" {
		t.Errorf("unexpected result (+expected, -got): %s", diff)
	}
}

func TestIntOrStringTypeValueFromTerraform(t *testing.T) {
	t.Parallel()

	type testCase struct {
		receiver    IntOrStringType
		input       tftypes.Value
		expected    attr.Value
		expectedErr string
	}
	tests := map[string]testCase{
		"string": {
			receiver: IntOrStringType{},
			input:    tftypes.NewValue(tftypes.String, "hello"),
			expected: IntOrString{
				Value:      tftypes.NewValue(tftypes.String, "hello"),
				StringType: true,
				NumberType: false,
				Null:       false,
				Unknown:    false,
			},
		},
		"number": {
			receiver: IntOrStringType{},
			input:    tftypes.NewValue(tftypes.Number, 123),
			expected: IntOrString{
				Value:      tftypes.NewValue(tftypes.Number, 123),
				StringType: false,
				NumberType: true,
				Null:       false,
				Unknown:    false,
			},
		},
		"nil": {
			receiver: IntOrStringType{},
			input:    tftypes.NewValue(nil, nil),
			expected: IntOrString{
				Value:      tftypes.NewValue(nil, nil),
				StringType: false,
				NumberType: false,
				Null:       true,
				Unknown:    false,
			},
		},
		"unknown": {
			receiver: IntOrStringType{},
			input:    tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
			expected: IntOrString{
				Value:      tftypes.NewValue(nil, nil),
				StringType: false,
				NumberType: false,
				Null:       false,
				Unknown:    true,
			},
		},
		"object": {
			receiver: IntOrStringType{},
			input: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"a": tftypes.String,
				},
			}, nil),
			expectedErr: "expected tftypes.Number or tftypes.String, got tftypes.Object[\"a\":tftypes.String]",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := test.receiver.ValueFromTerraform(context.Background(), test.input)
			if err != nil {
				if test.expectedErr == "" {
					t.Errorf("Unexpected error: %s", err.Error())
					return
				}
				if err.Error() != test.expectedErr {
					t.Errorf("Expected error to be %q, got %q", test.expectedErr, err.Error())
					return
				}
			}
			if test.expectedErr != "" && err == nil {
				t.Errorf("Expected err to be %q, got nil", test.expectedErr)
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

func TestIntOrStringToTerraformValue(t *testing.T) {
	type testCase struct {
		receiver    IntOrString
		expected    tftypes.Value
		expectedErr string
	}
	tests := map[string]testCase{
		"string": {
			receiver: IntOrString{
				Value: tftypes.NewValue(tftypes.String, "hello"),
			},
			expected: tftypes.NewValue(tftypes.String, "hello"),
		},
		"number": {
			receiver: IntOrString{
				Value: tftypes.NewValue(tftypes.Number, 123),
			},
			expected: tftypes.NewValue(tftypes.Number, 123),
		},
		"unknown": {
			receiver: IntOrString{
				Unknown: true,
			},
			expected: tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
		},
		"null": {
			receiver: IntOrString{
				Null: true,
			},
			expected: tftypes.NewValue(tftypes.DynamicPseudoType, nil),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, gotErr := test.receiver.ToTerraformValue(context.Background())

			if test.expectedErr == "" && gotErr != nil {
				t.Errorf("Unexpected error: %s", gotErr)
				return
			}

			if test.expectedErr != "" {
				if gotErr == nil {
					t.Errorf("Expected error to be %q, got none", test.expectedErr)
					return
				}

				if test.expectedErr != gotErr.Error() {
					t.Errorf("Expected error to be %q, got %q", test.expectedErr, gotErr.Error())
					return
				}
			}

			if diff := cmp.Diff(test.expected, got); diff != "" {
				t.Errorf("Unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestIntOrStringYamlMarshaller(t *testing.T) {
	type testCase struct {
		value    attr.Value
		expected string
	}
	tests := map[string]testCase{
		"int": {
			value:    IntOrString{Value: tftypes.NewValue(tftypes.Number, 123)},
			expected: "123",
		},
		"string": {
			value:    IntOrString{Value: tftypes.NewValue(tftypes.String, "hello")},
			expected: "hello",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			marshal, err := yaml.Marshal(test.value)
			if err != nil {
				return
			}
			if diff := cmp.Diff(test.expected+"\n", string(marshal)); diff != "" {
				t.Errorf("unexpected result (-expected, +got): %s", diff)
			}
		})
	}
}
