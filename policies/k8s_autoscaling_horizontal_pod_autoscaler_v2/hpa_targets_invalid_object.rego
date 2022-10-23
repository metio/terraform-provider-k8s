# SPDX-FileCopyrightText: The terraform-provider-k8s Authors
# SPDX-License-Identifier: 0BSD

package autoscaling_horizontal_pod_autoscaler_v2

import future.keywords

deny[msg] if {
    resource := input.resource.k8s_autoscaling_horizontal_pod_autoscaler_v2[name]
    metric := resource.spec.metrics[index]
    not checkIsValidObjectMetric(metric)

    msg := sprintf("k8s_autoscaling_horizontal_pod_autoscaler_v2.%s: Targets Invalid Object", [name])
}

checkIsValidObjectMetric(metric) {
    metric.type == "Object"
    metric.object != null
    metric.object.metric != null
    metric.object.target != null
    metric.object.describedObject.name != null
    metric.object.describedObject.apiVersion != null
    metric.object.describedObject.kind != null
}
