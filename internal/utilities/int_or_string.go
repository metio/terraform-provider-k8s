/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package utilities

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"gopkg.in/yaml.v3"
	"math/big"
)

type IntOrStringType struct{}

var (
	_ attr.Type      = IntOrStringType{}
	_ attr.Value     = IntOrString{}
	_ yaml.Marshaler = IntOrString{}
	_ yaml.IsZeroer  = IntOrString{}
)

func (d IntOrStringType) ValueType(_ context.Context) attr.Value {
	return IntOrString{}
}

func (d IntOrStringType) Equal(o attr.Type) bool {
	other, ok := o.(IntOrStringType)
	if !ok {
		return false
	}
	return d == other
}

func (d IntOrStringType) String() string {
	return "utilities.IntOrStringType"
}

func (d IntOrStringType) ApplyTerraform5AttributePathStep(_ tftypes.AttributePathStep) (interface{}, error) {
	return nil, tftypes.ErrInvalidStep
}

func (d IntOrStringType) TerraformType(_ context.Context) tftypes.Type {
	return tftypes.DynamicPseudoType
}

func (d IntOrStringType) ValueFromTerraform(_ context.Context, in tftypes.Value) (attr.Value, error) {
	dynamic := IntOrString{}

	if in.Type() == nil {
		dynamic.Null = true
		return dynamic, nil
	}
	if !(in.Type().Is(tftypes.Number) || in.Type().Is(tftypes.String) || in.Type().Is(tftypes.DynamicPseudoType)) {
		return nil, fmt.Errorf("expected %s or %s, got %s", tftypes.Number, tftypes.String, in.Type())
	}
	if in.IsNull() {
		dynamic.Null = true
		dynamic.StringType = in.Type().Is(tftypes.String)
		dynamic.NumberType = in.Type().Is(tftypes.Number)
		return dynamic, nil
	}
	if !in.IsKnown() {
		dynamic.Unknown = true
		dynamic.StringType = in.Type().Is(tftypes.String)
		dynamic.NumberType = in.Type().Is(tftypes.Number)
		return dynamic, nil
	}

	var stringValue string
	err := in.As(&stringValue)
	if err == nil {
		dynamic.StringType = true
	}
	var numberValue big.Float
	err = in.As(&numberValue)
	if err == nil {
		dynamic.NumberType = true
	}

	dynamic.Value = in

	return dynamic, nil
}

type IntOrString struct {
	Unknown    bool
	Null       bool
	StringType bool
	NumberType bool
	Value      tftypes.Value
}

func (d IntOrString) IsZero() bool {
	return d.Null || d.Unknown
}

func (d IntOrString) MarshalYAML() (interface{}, error) {
	return intOrStringAsYaml(d.Value)
}

func intOrStringAsYaml(value tftypes.Value) (interface{}, error) {
	if !value.IsKnown() {
		return nil, nil
	}

	switch {
	case value.Type().Is(tftypes.String):
		var yamlValue string
		err := value.As(&yamlValue)
		if err != nil {
			return nil, err
		}
		return yamlValue, nil
	case value.Type().Is(tftypes.Number):
		var yamlValue big.Float
		err := value.As(&yamlValue)
		if err != nil {
			return nil, err
		}
		if yamlValue.IsInt() {
			inv, acc := yamlValue.Int64()
			if acc != big.Exact {
				return nil, fmt.Errorf("%s inexact integer approximation when converting number value", value.String())
			}
			return inv, nil
		} else {
			return nil, fmt.Errorf("%s is not an integer", value.String())
		}
	default:
		return nil, nil
	}
}

func (d IntOrString) Type(_ context.Context) attr.Type {
	return IntOrStringType{}
}

func (d IntOrString) ToTerraformValue(_ context.Context) (tftypes.Value, error) {
	if d.Unknown {
		if d.StringType {
			return tftypes.NewValue(tftypes.String, tftypes.UnknownValue), nil
		}
		if d.NumberType {
			return tftypes.NewValue(tftypes.Number, tftypes.UnknownValue), nil
		}
		return tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue), nil
	}
	if d.Null {
		if d.StringType {
			return tftypes.NewValue(tftypes.String, nil), nil
		}
		if d.NumberType {
			return tftypes.NewValue(tftypes.Number, nil), nil
		}
		return tftypes.NewValue(tftypes.DynamicPseudoType, nil), nil
	}
	return d.Value, nil
}

func (d IntOrString) Equal(value attr.Value) bool {
	other, ok := value.(IntOrString)
	if !ok {
		return false
	}
	if d.Unknown != other.Unknown {
		return false
	}
	if d.Null != other.Null {
		return false
	}
	if d.StringType != other.StringType {
		return false
	}
	if d.NumberType != other.NumberType {
		return false
	}
	return d.Value.Equal(other.Value)
}

func (d IntOrString) IsNull() bool {
	return d.Null
}

func (d IntOrString) IsUnknown() bool {
	return d.Unknown
}

func (d IntOrString) String() string {
	if d.Unknown {
		return attr.UnknownValueString
	}

	if d.Null {
		return attr.NullValueString
	}

	return d.Value.String()
}
