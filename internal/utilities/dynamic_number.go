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

type DynamicNumberType struct{}

var (
	_ attr.Type      = DynamicNumberType{}
	_ attr.Value     = DynamicNumber{}
	_ yaml.Marshaler = DynamicNumber{}
	_ yaml.IsZeroer  = DynamicNumber{}
)

func (d DynamicNumberType) ValueType(_ context.Context) attr.Value {
	return DynamicNumber{}
}

func (d DynamicNumberType) Equal(o attr.Type) bool {
	other, ok := o.(DynamicNumberType)
	if !ok {
		return false
	}
	return d == other
}

func (d DynamicNumberType) String() string {
	return "utilities.DynamicNumberType"
}

func (d DynamicNumberType) ApplyTerraform5AttributePathStep(_ tftypes.AttributePathStep) (interface{}, error) {
	return nil, tftypes.ErrInvalidStep
}

func (d DynamicNumberType) TerraformType(_ context.Context) tftypes.Type {
	return tftypes.Number
}

func (d DynamicNumberType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	dynamic := DynamicNumber{}

	if in.Type() == nil {
		dynamic.Null = true
		return dynamic, nil
	}
	if !in.Type().Is(tftypes.Number) {
		return nil, fmt.Errorf("expected %s, got %s", d.TerraformType(ctx), in.Type())
	}
	if in.IsNull() {
		dynamic.Null = true
		return dynamic, nil
	}
	if !in.IsKnown() {
		dynamic.Unknown = true
		return dynamic, nil
	}

	dynamic.Value = in

	return dynamic, nil
}

type DynamicNumber struct {
	Unknown bool
	Null    bool
	Value   tftypes.Value
}

func (d DynamicNumber) IsZero() bool {
	return d.Null || d.Unknown
}

func (d DynamicNumber) MarshalYAML() (interface{}, error) {
	return numberAsYamlValue(d.Value)
}

func numberAsYamlValue(value tftypes.Value) (interface{}, error) {
	if !value.IsKnown() {
		return nil, nil
	}

	switch {
	case value.Type().Is(tftypes.Number):
		var yamlValue big.Float
		err := value.As(&yamlValue)
		if err != nil {
			return nil, err
		}
		if yamlValue.IsInt() {
			inv, _ := yamlValue.Int64()
			return inv, nil
		} else {
			inv, _ := yamlValue.Float64()
			return inv, nil
		}
	default:
		return nil, nil
	}
}

func (d DynamicNumber) Type(_ context.Context) attr.Type {
	return DynamicNumberType{}
}

func (d DynamicNumber) ToTerraformValue(_ context.Context) (tftypes.Value, error) {
	if d.Unknown {
		return tftypes.NewValue(tftypes.Number, tftypes.UnknownValue), nil
	}
	if d.Null {
		return tftypes.NewValue(tftypes.Number, nil), nil
	}
	return d.Value, nil
}

func (d DynamicNumber) Equal(value attr.Value) bool {
	other, ok := value.(DynamicNumber)
	if !ok {
		return false
	}
	if d.Unknown != other.Unknown {
		return false
	}
	if d.Null != other.Null {
		return false
	}
	return d.Value.Equal(other.Value)
}

func (d DynamicNumber) IsNull() bool {
	return d.Null
}

func (d DynamicNumber) IsUnknown() bool {
	return d.Unknown
}

func (d DynamicNumber) String() string {
	if d.Unknown {
		return attr.UnknownValueString
	}

	if d.Null {
		return attr.NullValueString
	}

	return d.Value.String()
}
