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

type DynamicType struct{}

var (
	_ attr.Type      = DynamicType{}
	_ attr.Value     = Dynamic{}
	_ yaml.Marshaler = Dynamic{}
	_ yaml.IsZeroer  = Dynamic{}
)

func (d DynamicType) ValueType(_ context.Context) attr.Value {
	return Dynamic{}
}

func (d DynamicType) Equal(o attr.Type) bool {
	other, ok := o.(DynamicType)
	if !ok {
		return false
	}
	return d == other
}

func (d DynamicType) String() string {
	return "utilities.DynamicType"
}

func (d DynamicType) ApplyTerraform5AttributePathStep(_ tftypes.AttributePathStep) (interface{}, error) {
	return nil, tftypes.ErrInvalidStep
}

func (d DynamicType) TerraformType(_ context.Context) tftypes.Type {
	return tftypes.DynamicPseudoType
}

func (d DynamicType) ValueFromTerraform(_ context.Context, in tftypes.Value) (attr.Value, error) {
	dynamic := Dynamic{}

	if in.Type() == nil {
		dynamic.Null = true
		return dynamic, nil
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

type Dynamic struct {
	Unknown bool
	Null    bool
	Value   tftypes.Value
}

func (d Dynamic) IsZero() bool {
	return d.Null || d.Unknown
}

func (d Dynamic) MarshalYAML() (interface{}, error) {
	return dynamicAsYamlValue(d.Value)
}

func dynamicAsYamlValue(value tftypes.Value) (interface{}, error) {
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
	case value.Type().Is(tftypes.Bool):
		var yamlValue bool
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
			inv, acc := yamlValue.Float64()
			if acc != big.Exact {
				return nil, fmt.Errorf("%s inexact float approximation when converting number value", value.String())
			}
			return inv, nil
		}
	case value.Type().Is(tftypes.List{}) || value.Type().Is(tftypes.Tuple{}) || value.Type().Is(tftypes.Set{}):
		var internalValues []tftypes.Value
		var yamlValues []interface{}
		err := value.As(&internalValues)
		if err != nil {
			return nil, fmt.Errorf("[%s] cannot extract contents of attribute: %s", value.String(), err)
		}
		for _, v := range internalValues {
			yamlValue, err := dynamicAsYamlValue(v)
			if err != nil {
				return nil, fmt.Errorf("[%s] cannot convert list element: %s", v.String(), err)
			}
			yamlValues = append(yamlValues, yamlValue)
		}
		return yamlValues, nil
	case value.Type().Is(tftypes.Map{}) || value.Type().Is(tftypes.Object{}):
		internalValues := make(map[string]tftypes.Value)
		yamlValues := make(map[string]interface{})
		err := value.As(&internalValues)
		if err != nil {
			return nil, fmt.Errorf("[%s] cannot extract contents of attribute: %s", value.String(), err)
		}
		for k, v := range internalValues {
			yamlValue, err := dynamicAsYamlValue(v)
			if err != nil {
				return nil, fmt.Errorf("[%s] cannot convert list element: %s", v.String(), err)
			}
			yamlValues[k] = yamlValue
		}
		return yamlValues, nil
	default:
		return nil, nil
	}
}

func (d Dynamic) Type(_ context.Context) attr.Type {
	return DynamicType{}
}

func (d Dynamic) ToTerraformValue(_ context.Context) (tftypes.Value, error) {
	if d.Unknown {
		return tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue), nil
	}
	if d.Null {
		return tftypes.NewValue(tftypes.DynamicPseudoType, nil), nil
	}
	return d.Value, nil
}

func (d Dynamic) Equal(value attr.Value) bool {
	other, ok := value.(Dynamic)
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

func (d Dynamic) IsNull() bool {
	return d.Null
}

func (d Dynamic) IsUnknown() bool {
	return d.Unknown
}

func (d Dynamic) String() string {
	if d.Unknown {
		return attr.UnknownValueString
	}

	if d.Null {
		return attr.NullValueString
	}

	return d.Value.String()
}
