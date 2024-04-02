/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package customtypes

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/attr/xattr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"math/big"
	yaml "sigs.k8s.io/yaml/goyaml.v3"
	"strconv"
)

var (
	_ basetypes.DynamicTypable  = IntOrStringType{}
	_ xattr.TypeWithValidate    = IntOrStringType{}
	_ basetypes.DynamicValuable = IntOrStringValue{}
	_ yaml.Marshaler            = IntOrStringValue{}
)

type IntOrStringType struct {
	basetypes.DynamicType
	// ... potentially other fields ...
}

type IntOrStringValue struct {
	basetypes.DynamicValue
	// ... potentially other fields ...
}

func (t IntOrStringType) Equal(o attr.Type) bool {
	other, ok := o.(IntOrStringType)

	if !ok {
		return false
	}

	return t.DynamicType.Equal(other.DynamicType)
}

func (t IntOrStringType) String() string {
	return "IntOrStringType"
}

func (t IntOrStringType) ValueFromDynamic(ctx context.Context, in basetypes.DynamicValue) (basetypes.DynamicValuable, diag.Diagnostics) {
	value := IntOrStringValue{
		DynamicValue: in,
	}

	return value, nil
}

func (t IntOrStringType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.DynamicType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	dynamicValue, ok := attrValue.(basetypes.DynamicValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	dynamicValuable, diags := t.ValueFromDynamic(ctx, dynamicValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting DynamicValue to DynamicValuable: %v", diags)
	}

	return dynamicValuable, nil
}

func (t IntOrStringType) ValueType(ctx context.Context) attr.Value {
	return IntOrStringValue{}
}

func (t IntOrStringType) Validate(ctx context.Context, value tftypes.Value, path path.Path) diag.Diagnostics {
	if value.IsNull() || !value.IsKnown() {
		return nil
	}

	var diags diag.Diagnostics

	valueType := value.Type()

	if !valueType.Is(tftypes.String) && !valueType.Is(tftypes.Number) {
		diags.AddAttributeError(
			path,
			"Invalid Terraform Value",
			"IntOrString type only accepts values of type tftypes.String or tftypes.Number but got "+valueType.String()+". "+
				"Make sure your Terraform configuration matches this expectation.\n\n"+
				"Path: "+path.String(),
		)
		return diags
	}

	return diags
}

func (v IntOrStringValue) Equal(o attr.Value) bool {
	other, ok := o.(IntOrStringValue)

	if !ok {
		return false
	}

	return v.DynamicValue.Equal(other.DynamicValue)
}

func (v IntOrStringValue) Type(ctx context.Context) attr.Type {
	return IntOrStringType{}
}

func (v IntOrStringValue) MarshalYAML() (interface{}, error) {
	if v.IsUnknown() || v.IsNull() {
		return nil, nil
	}

	tfValue, err := v.UnderlyingValue().ToTerraformValue(context.Background())
	if err != nil {
		return nil, err
	}

	switch {
	case tfValue.Type().Is(tftypes.String):
		var yamlValue string
		err = tfValue.As(&yamlValue)
		if err != nil {
			return nil, err
		}
		i, err := strconv.Atoi(yamlValue)
		if err == nil {
			return i, nil
		}
		return yamlValue, nil
	case tfValue.Type().Is(tftypes.Number):
		var yamlValue big.Float
		err = tfValue.As(&yamlValue)
		if err != nil {
			return nil, err
		}
		if yamlValue.IsInt() {
			inv, acc := yamlValue.Int64()
			if acc != big.Exact {
				return nil, fmt.Errorf("%s inexact integer approximation when converting number value", v.String())
			}
			return inv, nil
		} else {
			return nil, fmt.Errorf("%s is not an integer", v.String())
		}
	default:
		return nil, nil
	}
}
