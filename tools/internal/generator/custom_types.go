/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package generator

type customTypeConfig struct {
	attributeType string
	valueType     string
	elementType   string
	goType        string
	customType    string
}

var customTypes = map[string]map[string]customTypeConfig{
	"kyverno_io_cluster_policy_v1": {
		"spec.rules.context.apiCall.data.value": customTypeConfig{
			attributeType: "schema.StringAttribute",
			valueType:     "types.String",
			elementType:   "",
			goType:        "custom_types.Normalized",
			customType:    "custom_types.NormalizedType{}",
		},
	},
}
