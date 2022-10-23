# SPDX-FileCopyrightText: The terraform-provider-k8s Authors
# SPDX-License-Identifier: 0BSD

package required_labels

import future.keywords

test_label_exists if {
    count(deny) == 0 with input as {
        "resource": {"k8s_config_map_v1": {"sample": {"metadata": {"labels": {"some": "value"}}}}},
        "parameters": {"labels": ["some"]},
    }
}

test_label_missing if {
    count(deny) == 1 with input as {
        "resource": {"k8s_config_map_v1": {"sample": {"metadata": {"labels": {"some": "value"}}}}},
        "parameters": {"labels": ["one", "two"]},
    }
}
