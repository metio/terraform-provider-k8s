/*
 * SPDX-FileCopyrightText: The terraform-provider-k8s Authors
 * SPDX-License-Identifier: 0BSD
 */

package utilities

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func DetermineTimeout(attributes map[string]attr.Value) time.Duration {
	if value, exists := attributes["timeout"]; exists {
		if valueInt, typed := value.(types.Int64); typed {
			return time.Second * time.Duration(valueInt.ValueInt64())
		}
	}
	return time.Second * 30
}

func DeterminePollInterval(attributes map[string]attr.Value) time.Duration {
	if value, exists := attributes["poll_interval"]; exists {
		if valueInt, typed := value.(types.Int64); typed {
			return time.Second * time.Duration(valueInt.ValueInt64())
		}
	}
	return time.Second * 5
}
