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
	"math/big"
	"testing"
)

func TestDynamicTypeTerraformType(t *testing.T) {
	t.Parallel()
	result := DynamicType{}.TerraformType(context.Background())
	if diff := cmp.Diff(result, tftypes.DynamicPseudoType); diff != "" {
		t.Errorf("unexpected result (+expected, -got): %s", diff)
	}
}

func TestDynamicTypeValueFromTerraform(t *testing.T) {
	t.Parallel()

	type testCase struct {
		receiver    DynamicType
		input       tftypes.Value
		expected    attr.Value
		expectedErr string
	}
	tests := map[string]testCase{
		"simple-object": {
			receiver: DynamicType{},
			input: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"a": tftypes.String,
					"b": tftypes.Bool,
					"c": tftypes.Number,
				},
			}, map[string]tftypes.Value{
				"a": tftypes.NewValue(tftypes.String, "red"),
				"b": tftypes.NewValue(tftypes.Bool, true),
				"c": tftypes.NewValue(tftypes.Number, 123),
			}),
			expected: Dynamic{
				Value: tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"a": tftypes.String,
						"b": tftypes.Bool,
						"c": tftypes.Number,
					},
				}, map[string]tftypes.Value{
					"a": tftypes.NewValue(tftypes.String, "red"),
					"b": tftypes.NewValue(tftypes.Bool, true),
					"c": tftypes.NewValue(tftypes.Number, 123),
				}),
			},
		},
		"single-string": {
			receiver: DynamicType{},
			input:    tftypes.NewValue(tftypes.String, "hello"),
			expected: Dynamic{Value: tftypes.NewValue(tftypes.String, "hello")},
		},
		"nil-type": {
			receiver: DynamicType{},
			input:    tftypes.NewValue(nil, nil),
			expected: Dynamic{Null: true},
		},
		"unknown": {
			receiver: DynamicType{},
			input: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"a": tftypes.String,
				},
			}, tftypes.UnknownValue),
			expected: Dynamic{Unknown: true},
		},
		"null": {
			receiver: DynamicType{},
			input: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"a": tftypes.String,
				},
			}, nil),
			expected: Dynamic{Null: true},
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

func TestDynamicTypeToTerraformValue(t *testing.T) {
	type testCase struct {
		receiver    Dynamic
		expected    tftypes.Value
		expectedErr string
	}
	tests := map[string]testCase{
		"value": {
			receiver: Dynamic{
				Value: tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"a": tftypes.List{ElementType: tftypes.String},
						"b": tftypes.String,
						"c": tftypes.Bool,
						"d": tftypes.Number,
						"e": tftypes.Object{AttributeTypes: map[string]tftypes.Type{"name": tftypes.String}},
						"f": tftypes.Set{ElementType: tftypes.String},
					},
				}, map[string]tftypes.Value{
					"a": tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
						tftypes.NewValue(tftypes.String, "hello"),
						tftypes.NewValue(tftypes.String, "world"),
					}),
					"b": tftypes.NewValue(tftypes.String, "woohoo"),
					"c": tftypes.NewValue(tftypes.Bool, true),
					"d": tftypes.NewValue(tftypes.Number, big.NewFloat(1234)),
					"e": tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"name": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"name": tftypes.NewValue(tftypes.String, "testing123"),
					}),
					"f": tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, []tftypes.Value{
						tftypes.NewValue(tftypes.String, "hello"),
						tftypes.NewValue(tftypes.String, "world"),
					}),
				}),
			},
			expected: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"a": tftypes.List{ElementType: tftypes.String},
					"b": tftypes.String,
					"c": tftypes.Bool,
					"d": tftypes.Number,
					"e": tftypes.Object{AttributeTypes: map[string]tftypes.Type{"name": tftypes.String}},
					"f": tftypes.Set{ElementType: tftypes.String},
				},
			}, map[string]tftypes.Value{
				"a": tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
					tftypes.NewValue(tftypes.String, "hello"),
					tftypes.NewValue(tftypes.String, "world"),
				}),
				"b": tftypes.NewValue(tftypes.String, "woohoo"),
				"c": tftypes.NewValue(tftypes.Bool, true),
				"d": tftypes.NewValue(tftypes.Number, big.NewFloat(1234)),
				"e": tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"name": tftypes.String,
					},
				}, map[string]tftypes.Value{
					"name": tftypes.NewValue(tftypes.String, "testing123"),
				}),
				"f": tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, []tftypes.Value{
					tftypes.NewValue(tftypes.String, "hello"),
					tftypes.NewValue(tftypes.String, "world"),
				}),
			}),
		},
		"unknown": {
			receiver: Dynamic{
				Unknown: true,
			},
			expected: tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
		},
		"null": {
			receiver: Dynamic{
				Null: true,
			},
			expected: tftypes.NewValue(tftypes.DynamicPseudoType, nil),
		},
		"partial-unknown": {
			receiver: Dynamic{
				Value: tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"a": tftypes.List{ElementType: tftypes.String},
						"b": tftypes.String,
						"c": tftypes.Bool,
						"d": tftypes.Number,
						"e": tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"name": tftypes.String,
							},
						},
						"f": tftypes.Set{ElementType: tftypes.String},
					},
				}, map[string]tftypes.Value{
					"a": tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
						tftypes.NewValue(tftypes.String, "hello"),
						tftypes.NewValue(tftypes.String, "world"),
					}),
					"b": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
					"c": tftypes.NewValue(tftypes.Bool, true),
					"d": tftypes.NewValue(tftypes.Number, big.NewFloat(1234)),
					"e": tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"name": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"name": tftypes.NewValue(tftypes.String, "testing123"),
					}),
					"f": tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, []tftypes.Value{
						tftypes.NewValue(tftypes.String, "hello"),
						tftypes.NewValue(tftypes.String, "world"),
					}),
				}),
			},
			expected: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"a": tftypes.List{ElementType: tftypes.String},
					"b": tftypes.String,
					"c": tftypes.Bool,
					"d": tftypes.Number,
					"e": tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"name": tftypes.String,
						},
					},
					"f": tftypes.Set{ElementType: tftypes.String},
				},
			}, map[string]tftypes.Value{
				"a": tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
					tftypes.NewValue(tftypes.String, "hello"),
					tftypes.NewValue(tftypes.String, "world"),
				}),
				"b": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
				"c": tftypes.NewValue(tftypes.Bool, true),
				"d": tftypes.NewValue(tftypes.Number, big.NewFloat(1234)),
				"e": tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"name": tftypes.String,
					},
				}, map[string]tftypes.Value{
					"name": tftypes.NewValue(tftypes.String, "testing123"),
				}),
				"f": tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, []tftypes.Value{
					tftypes.NewValue(tftypes.String, "hello"),
					tftypes.NewValue(tftypes.String, "world"),
				}),
			}),
		},
		"partial-null": {
			receiver: Dynamic{
				Value: tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"a": tftypes.List{ElementType: tftypes.String},
						"b": tftypes.String,
						"c": tftypes.Bool,
						"d": tftypes.Number,
						"e": tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"name": tftypes.String,
							},
						},
						"f": tftypes.Set{ElementType: tftypes.String},
					},
				}, map[string]tftypes.Value{
					"a": tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
						tftypes.NewValue(tftypes.String, "hello"),
						tftypes.NewValue(tftypes.String, "world"),
					}),
					"b": tftypes.NewValue(tftypes.String, nil),
					"c": tftypes.NewValue(tftypes.Bool, true),
					"d": tftypes.NewValue(tftypes.Number, big.NewFloat(1234)),
					"e": tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"name": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"name": tftypes.NewValue(tftypes.String, "testing123"),
					}),
					"f": tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, []tftypes.Value{
						tftypes.NewValue(tftypes.String, "hello"),
						tftypes.NewValue(tftypes.String, "world"),
					}),
				}),
			},
			expected: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"a": tftypes.List{ElementType: tftypes.String},
					"b": tftypes.String,
					"c": tftypes.Bool,
					"d": tftypes.Number,
					"e": tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"name": tftypes.String,
						},
					},
					"f": tftypes.Set{ElementType: tftypes.String},
				},
			}, map[string]tftypes.Value{
				"a": tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
					tftypes.NewValue(tftypes.String, "hello"),
					tftypes.NewValue(tftypes.String, "world"),
				}),
				"b": tftypes.NewValue(tftypes.String, nil),
				"c": tftypes.NewValue(tftypes.Bool, true),
				"d": tftypes.NewValue(tftypes.Number, big.NewFloat(1234)),
				"e": tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"name": tftypes.String,
					},
				}, map[string]tftypes.Value{
					"name": tftypes.NewValue(tftypes.String, "testing123"),
				}),
				"f": tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, []tftypes.Value{
					tftypes.NewValue(tftypes.String, "hello"),
					tftypes.NewValue(tftypes.String, "world"),
				}),
			}),
		},
		"deep-partial-unknown": {
			receiver: Dynamic{
				Value: tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"a": tftypes.List{ElementType: tftypes.String},
						"b": tftypes.String,
						"c": tftypes.Bool,
						"d": tftypes.Number,
						"e": tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"name": tftypes.String,
							},
						},
						"f": tftypes.Set{ElementType: tftypes.String},
					},
				}, map[string]tftypes.Value{
					"a": tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
						tftypes.NewValue(tftypes.String, "hello"),
						tftypes.NewValue(tftypes.String, "world"),
					}),
					"b": tftypes.NewValue(tftypes.String, "woohoo"),
					"c": tftypes.NewValue(tftypes.Bool, true),
					"d": tftypes.NewValue(tftypes.Number, big.NewFloat(1234)),
					"e": tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"name": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"name": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
					}),
					"f": tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, []tftypes.Value{
						tftypes.NewValue(tftypes.String, "hello"),
						tftypes.NewValue(tftypes.String, "world"),
					}),
				}),
			},
			expected: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"a": tftypes.List{ElementType: tftypes.String},
					"b": tftypes.String,
					"c": tftypes.Bool,
					"d": tftypes.Number,
					"e": tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"name": tftypes.String,
						},
					},
					"f": tftypes.Set{ElementType: tftypes.String},
				},
			}, map[string]tftypes.Value{
				"a": tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
					tftypes.NewValue(tftypes.String, "hello"),
					tftypes.NewValue(tftypes.String, "world"),
				}),
				"b": tftypes.NewValue(tftypes.String, "woohoo"),
				"c": tftypes.NewValue(tftypes.Bool, true),
				"d": tftypes.NewValue(tftypes.Number, big.NewFloat(1234)),
				"e": tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"name": tftypes.String,
					},
				}, map[string]tftypes.Value{
					"name": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
				}),
				"f": tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, []tftypes.Value{
					tftypes.NewValue(tftypes.String, "hello"),
					tftypes.NewValue(tftypes.String, "world"),
				}),
			}),
		},
		"deep-partial-null": {
			receiver: Dynamic{
				Value: tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"a": tftypes.List{ElementType: tftypes.String},
						"b": tftypes.String,
						"c": tftypes.Bool,
						"d": tftypes.Number,
						"e": tftypes.Object{
							AttributeTypes: map[string]tftypes.Type{
								"name": tftypes.String,
							},
						},
						"f": tftypes.Set{ElementType: tftypes.String},
					},
				}, map[string]tftypes.Value{
					"a": tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
						tftypes.NewValue(tftypes.String, "hello"),
						tftypes.NewValue(tftypes.String, "world"),
					}),
					"b": tftypes.NewValue(tftypes.String, "woohoo"),
					"c": tftypes.NewValue(tftypes.Bool, true),
					"d": tftypes.NewValue(tftypes.Number, big.NewFloat(1234)),
					"e": tftypes.NewValue(tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"name": tftypes.String,
						},
					}, map[string]tftypes.Value{
						"name": tftypes.NewValue(tftypes.String, nil),
					}),
					"f": tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, []tftypes.Value{
						tftypes.NewValue(tftypes.String, "hello"),
						tftypes.NewValue(tftypes.String, "world"),
					}),
				}),
			},
			expected: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"a": tftypes.List{ElementType: tftypes.String},
					"b": tftypes.String,
					"c": tftypes.Bool,
					"d": tftypes.Number,
					"e": tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"name": tftypes.String,
						},
					},
					"f": tftypes.Set{ElementType: tftypes.String},
				},
			}, map[string]tftypes.Value{
				"a": tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
					tftypes.NewValue(tftypes.String, "hello"),
					tftypes.NewValue(tftypes.String, "world"),
				}),
				"b": tftypes.NewValue(tftypes.String, "woohoo"),
				"c": tftypes.NewValue(tftypes.Bool, true),
				"d": tftypes.NewValue(tftypes.Number, big.NewFloat(1234)),
				"e": tftypes.NewValue(tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"name": tftypes.String,
					},
				}, map[string]tftypes.Value{
					"name": tftypes.NewValue(tftypes.String, nil),
				}),
				"f": tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, []tftypes.Value{
					tftypes.NewValue(tftypes.String, "hello"),
					tftypes.NewValue(tftypes.String, "world"),
				}),
			}),
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

func TestDynamicTypeYamlMarshaller_Object(t *testing.T) {
	value := Dynamic{
		Value: tftypes.NewValue(tftypes.Object{
			AttributeTypes: map[string]tftypes.Type{
				"a": tftypes.List{ElementType: tftypes.String},
				"b": tftypes.String,
				"c": tftypes.Bool,
				"d": tftypes.Number,
				"e": tftypes.Object{AttributeTypes: map[string]tftypes.Type{"name": tftypes.String}},
				"f": tftypes.Set{ElementType: tftypes.String},
			},
		}, map[string]tftypes.Value{
			"a": tftypes.NewValue(tftypes.List{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "hello"),
				tftypes.NewValue(tftypes.String, "world"),
			}),
			"b": tftypes.NewValue(tftypes.String, "woohoo"),
			"c": tftypes.NewValue(tftypes.Bool, true),
			"d": tftypes.NewValue(tftypes.Number, big.NewFloat(1234)),
			"e": tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"name": tftypes.String,
				},
			}, map[string]tftypes.Value{
				"name": tftypes.NewValue(tftypes.String, "testing123"),
			}),
			"f": tftypes.NewValue(tftypes.Set{ElementType: tftypes.String}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "hello"),
				tftypes.NewValue(tftypes.String, "world"),
			}),
		}),
	}

	marshal, err := yaml.Marshal(value)
	if err != nil {
		return
	}
	if diff := cmp.Diff("a:\n    - hello\n    - world\nb: woohoo\nc: true\nd: 1234\ne:\n    name: testing123\nf:\n    - hello\n    - world\n", string(marshal)); diff != "" {
		t.Errorf("unexpected result (-expected, +got): %s", diff)
	}
}

func TestDynamicTypeYamlMarshaller_String(t *testing.T) {
	value := Dynamic{
		Value: tftypes.NewValue(tftypes.String, "world"),
	}

	marshal, err := yaml.Marshal(value)
	if err != nil {
		return
	}
	if diff := cmp.Diff("world\n", string(marshal)); diff != "" {
		t.Errorf("unexpected result (-expected, +got): %s", diff)
	}
}
