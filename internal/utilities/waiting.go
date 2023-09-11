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
		if valueString, typed := value.(types.String); typed {
			timeout, err := time.ParseDuration(valueString.ValueString())
			if err == nil {
				if timeout > time.Second*0 {
					return timeout
				}
				return time.Hour * 168
			}
		}
	}
	return time.Second * 30
}

func DeterminePollInterval(attributes map[string]attr.Value) time.Duration {
	if value, exists := attributes["poll_interval"]; exists {
		if valueString, typed := value.(types.String); typed {
			timeout, err := time.ParseDuration(valueString.ValueString())
			if err == nil {
				return timeout
			}
		}
	}
	return time.Second * 5
}
