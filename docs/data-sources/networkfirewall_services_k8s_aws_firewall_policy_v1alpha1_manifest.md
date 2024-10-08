---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_networkfirewall_services_k8s_aws_firewall_policy_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "networkfirewall.services.k8s.aws"
description: |-
  FirewallPolicy is the Schema for the FirewallPolicies API
---

# k8s_networkfirewall_services_k8s_aws_firewall_policy_v1alpha1_manifest (Data Source)

FirewallPolicy is the Schema for the FirewallPolicies API

## Example Usage

```terraform
data "k8s_networkfirewall_services_k8s_aws_firewall_policy_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) FirewallPolicySpec defines the desired state of FirewallPolicy. The firewall policy defines the behavior of a firewall using a collection of stateless and stateful rule groups and other settings. You can use one firewall policy for multiple firewalls. This, along with FirewallPolicyResponse, define the policy. You can retrieve all objects for a firewall policy by calling DescribeFirewallPolicy. (see [below for nested schema](#nestedatt--spec))

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

- `firewall_policy` (Attributes) The rule groups and policy actions to use in the firewall policy. (see [below for nested schema](#nestedatt--spec--firewall_policy))
- `firewall_policy_name` (String) The descriptive name of the firewall policy. You can't change the name of a firewall policy after you create it.

Optional:

- `description` (String) A description of the firewall policy.
- `encryption_configuration` (Attributes) A complex type that contains settings for encryption of your firewall policy resources. (see [below for nested schema](#nestedatt--spec--encryption_configuration))
- `tags` (Attributes List) The key:value pairs to associate with the resource. (see [below for nested schema](#nestedatt--spec--tags))

<a id="nestedatt--spec--firewall_policy"></a>
### Nested Schema for `spec.firewall_policy`

Optional:

- `policy_variables` (Attributes) Contains variables that you can use to override default Suricata settings in your firewall policy. (see [below for nested schema](#nestedatt--spec--firewall_policy--policy_variables))
- `stateful_default_actions` (List of String)
- `stateful_engine_options` (Attributes) Configuration settings for the handling of the stateful rule groups in a firewall policy. (see [below for nested schema](#nestedatt--spec--firewall_policy--stateful_engine_options))
- `stateful_rule_group_references` (Attributes List) (see [below for nested schema](#nestedatt--spec--firewall_policy--stateful_rule_group_references))
- `stateless_custom_actions` (Attributes List) (see [below for nested schema](#nestedatt--spec--firewall_policy--stateless_custom_actions))
- `stateless_default_actions` (List of String)
- `stateless_fragment_default_actions` (List of String)
- `stateless_rule_group_references` (Attributes List) (see [below for nested schema](#nestedatt--spec--firewall_policy--stateless_rule_group_references))
- `tls_inspection_configuration_arn` (String)

<a id="nestedatt--spec--firewall_policy--policy_variables"></a>
### Nested Schema for `spec.firewall_policy.policy_variables`

Optional:

- `rule_variables` (Attributes) (see [below for nested schema](#nestedatt--spec--firewall_policy--policy_variables--rule_variables))

<a id="nestedatt--spec--firewall_policy--policy_variables--rule_variables"></a>
### Nested Schema for `spec.firewall_policy.policy_variables.rule_variables`

Optional:

- `definition` (List of String)



<a id="nestedatt--spec--firewall_policy--stateful_engine_options"></a>
### Nested Schema for `spec.firewall_policy.stateful_engine_options`

Optional:

- `rule_order` (String)
- `stream_exception_policy` (String)


<a id="nestedatt--spec--firewall_policy--stateful_rule_group_references"></a>
### Nested Schema for `spec.firewall_policy.stateful_rule_group_references`

Optional:

- `override` (Attributes) The setting that allows the policy owner to change the behavior of the rule group within a policy. (see [below for nested schema](#nestedatt--spec--firewall_policy--stateful_rule_group_references--override))
- `priority` (Number)
- `resource_arn` (String)

<a id="nestedatt--spec--firewall_policy--stateful_rule_group_references--override"></a>
### Nested Schema for `spec.firewall_policy.stateful_rule_group_references.override`

Optional:

- `action` (String)



<a id="nestedatt--spec--firewall_policy--stateless_custom_actions"></a>
### Nested Schema for `spec.firewall_policy.stateless_custom_actions`

Optional:

- `action_definition` (Attributes) A custom action to use in stateless rule actions settings. This is used in CustomAction. (see [below for nested schema](#nestedatt--spec--firewall_policy--stateless_custom_actions--action_definition))
- `action_name` (String)

<a id="nestedatt--spec--firewall_policy--stateless_custom_actions--action_definition"></a>
### Nested Schema for `spec.firewall_policy.stateless_custom_actions.action_definition`

Optional:

- `publish_metric_action` (Attributes) Stateless inspection criteria that publishes the specified metrics to Amazon CloudWatch for the matching packet. This setting defines a CloudWatch dimension value to be published. (see [below for nested schema](#nestedatt--spec--firewall_policy--stateless_custom_actions--action_definition--publish_metric_action))

<a id="nestedatt--spec--firewall_policy--stateless_custom_actions--action_definition--publish_metric_action"></a>
### Nested Schema for `spec.firewall_policy.stateless_custom_actions.action_definition.publish_metric_action`

Optional:

- `dimensions` (Attributes List) (see [below for nested schema](#nestedatt--spec--firewall_policy--stateless_custom_actions--action_definition--publish_metric_action--dimensions))

<a id="nestedatt--spec--firewall_policy--stateless_custom_actions--action_definition--publish_metric_action--dimensions"></a>
### Nested Schema for `spec.firewall_policy.stateless_custom_actions.action_definition.publish_metric_action.dimensions`

Optional:

- `value` (String)





<a id="nestedatt--spec--firewall_policy--stateless_rule_group_references"></a>
### Nested Schema for `spec.firewall_policy.stateless_rule_group_references`

Optional:

- `priority` (Number)
- `resource_arn` (String)



<a id="nestedatt--spec--encryption_configuration"></a>
### Nested Schema for `spec.encryption_configuration`

Optional:

- `key_id` (String)
- `type_` (String)


<a id="nestedatt--spec--tags"></a>
### Nested Schema for `spec.tags`

Optional:

- `key` (String)
- `value` (String)
