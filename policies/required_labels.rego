# SPDX-FileCopyrightText: The terraform-provider-k8s Authors
# SPDX-License-Identifier: 0BSD

# METADATA
# title: Required Labels
# description: This policy allows you to require certain labels are set on a resource.
package required_labels

import future.keywords

deny[msg] if {
    rule := input.resource[type][name]
    provided := {label | input.resource[type][name].metadata.labels[label]}
    required := {label | label := input.parameters.labels[_]}
    missing := required - provided
    count(missing) > 0

    msg := sprintf("%s.%s: Missing required labels: %v", [type, name, missing])
}
