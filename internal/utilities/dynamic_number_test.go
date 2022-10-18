/*
 * SPDX-FileCopyrightText: The terraform-provider-k8scr Authors
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

func TestDynamicNumberTypeTerraformType(t *testing.T) {
	t.Parallel()
	result := DynamicNumberType{}.TerraformType(context.Background())
	if diff := cmp.Diff(result, tftypes.Number); diff != "" {
		t.Errorf("unexpected result (+expected, -got): %s", diff)
	}
}

func TestDynamicNumberTypeValueFromTerraform(t *testing.T) {
	t.Parallel()

	type testCase struct {
		receiver    DynamicNumberType
		input       tftypes.Value
		expected    attr.Value
		expectedErr string
	}
	tests := map[string]testCase{
		"int": {
			receiver: DynamicNumberType{},
			input:    tftypes.NewValue(tftypes.Number, 123),
			expected: DynamicNumber{
				Value:   tftypes.NewValue(tftypes.Number, 123),
				Null:    false,
				Unknown: false,
			},
		},
		"float": {
			receiver: DynamicNumberType{},
			input:    tftypes.NewValue(tftypes.Number, 123.456),
			expected: DynamicNumber{
				Value:   tftypes.NewValue(tftypes.Number, 123.456),
				Null:    false,
				Unknown: false,
			},
		},
		"nil": {
			receiver: DynamicNumberType{},
			input:    tftypes.NewValue(nil, nil),
			expected: DynamicNumber{
				Value:   tftypes.NewValue(nil, nil),
				Null:    true,
				Unknown: false,
			},
		},
		"unknown": {
			receiver: DynamicNumberType{},
			input:    tftypes.NewValue(tftypes.Number, tftypes.UnknownValue),
			expected: DynamicNumber{
				Value:   tftypes.NewValue(nil, nil),
				Null:    false,
				Unknown: true,
			},
		},
		"object": {
			receiver: DynamicNumberType{},
			input: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"a": tftypes.String,
				},
			}, nil),
			expectedErr: "expected tftypes.Number, got tftypes.Object[\"a\":tftypes.String]",
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

func TestDynamicNumberToTerraformValue(t *testing.T) {
	type testCase struct {
		receiver    DynamicNumber
		expected    tftypes.Value
		expectedErr string
	}
	tests := map[string]testCase{
		"int": {
			receiver: DynamicNumber{
				Value: tftypes.NewValue(tftypes.Number, 123),
			},
			expected: tftypes.NewValue(tftypes.Number, 123),
		},
		"float": {
			receiver: DynamicNumber{
				Value: tftypes.NewValue(tftypes.Number, 123.456),
			},
			expected: tftypes.NewValue(tftypes.Number, 123.456),
		},
		"unknown": {
			receiver: DynamicNumber{
				Unknown: true,
			},
			expected: tftypes.NewValue(tftypes.Number, tftypes.UnknownValue),
		},
		"null": {
			receiver: DynamicNumber{
				Null: true,
			},
			expected: tftypes.NewValue(tftypes.Number, nil),
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

func TestDynamicNumberYamlMarshaller(t *testing.T) {
	type testCase struct {
		number   interface{}
		expected string
	}
	tests := map[string]testCase{
		"int": {
			number:   123,
			expected: "123",
		},
		"float": {
			number:   123.456,
			expected: "123.456",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			dynamicNumber := DynamicNumber{Value: tftypes.NewValue(tftypes.Number, test.number)}
			marshal, err := yaml.Marshal(dynamicNumber)
			if err != nil {
				return
			}
			if diff := cmp.Diff(test.expected+"\n", string(marshal)); diff != "" {
				t.Errorf("unexpected result (-expected, +got): %s", diff)
			}
		})
	}
}
