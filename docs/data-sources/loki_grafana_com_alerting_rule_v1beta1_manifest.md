---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_loki_grafana_com_alerting_rule_v1beta1_manifest Data Source - terraform-provider-k8s"
subcategory: "loki.grafana.com"
description: |-
  AlertingRule is the Schema for the alertingrules API
---

# k8s_loki_grafana_com_alerting_rule_v1beta1_manifest (Data Source)

AlertingRule is the Schema for the alertingrules API

## Example Usage

```terraform
data "k8s_loki_grafana_com_alerting_rule_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `metadata` (Attributes) Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details. (see [below for nested schema](#nestedatt--metadata))

### Optional

- `spec` (Attributes) AlertingRuleSpec defines the desired state of AlertingRule (see [below for nested schema](#nestedatt--spec))

### Read-Only

- `yaml` (String) The generated manifest in YAML format.

<a id="nestedatt--metadata"></a>
### Nested Schema for `metadata`

Required:

- `name` (String) Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.
- `namespace` (String) Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.

Optional:

- `annotations` (Map of String) Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.
- `labels` (Map of String) Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.


<a id="nestedatt--spec"></a>
### Nested Schema for `spec`

Required:

- `tenant_id` (String) TenantID of tenant where the alerting rules are evaluated in.

Optional:

- `groups` (Attributes List) List of groups for alerting rules. (see [below for nested schema](#nestedatt--spec--groups))

<a id="nestedatt--spec--groups"></a>
### Nested Schema for `spec.groups`

Required:

- `name` (String) Name of the alerting rule group. Must be unique within all alerting rules.
- `rules` (Attributes List) Rules defines a list of alerting rules (see [below for nested schema](#nestedatt--spec--groups--rules))

Optional:

- `interval` (String) Interval defines the time interval between evaluation of the givenalerting rule.
- `limit` (Number) Limit defines the number of alerts an alerting rule can produce. 0 is no limit.

<a id="nestedatt--spec--groups--rules"></a>
### Nested Schema for `spec.groups.rules`

Required:

- `expr` (String) The LogQL expression to evaluate. Every evaluation cycle this isevaluated at the current time, and all resultant time series becomepending/firing alerts.

Optional:

- `alert` (String) The name of the alert. Must be a valid label value.
- `annotations` (Map of String) Annotations to add to each alert.
- `for` (String) Alerts are considered firing once they have been returned for this long.Alerts which have not yet fired for long enough are considered pending.
- `labels` (Map of String) Labels to add to each alert.