---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_iam_services_k8s_aws_group_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "iam.services.k8s.aws"
description: |-
  Group is the Schema for the Groups API
---

# k8s_iam_services_k8s_aws_group_v1alpha1_manifest (Data Source)

Group is the Schema for the Groups API

## Example Usage

```terraform
data "k8s_iam_services_k8s_aws_group_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) GroupSpec defines the desired state of Group.Contains information about an IAM group entity.This data type is used as a response element in the following operations:   * CreateGroup   * GetGroup   * ListGroups (see [below for nested schema](#nestedatt--spec))

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

- `name` (String) The name of the group to create. Do not include the path in this value.IAM user, group, role, and policy names must be unique within the account.Names are not distinguished by case. For example, you cannot create resourcesnamed both 'MyResource' and 'myresource'.

Optional:

- `inline_policies` (Map of String)
- `path` (String) The path to the group. For more information about paths, see IAM identifiers(https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html)in the IAM User Guide.This parameter is optional. If it is not included, it defaults to a slash(/).This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))a string of characters consisting of either a forward slash (/) by itselfor a string that must begin and end with forward slashes. In addition, itcan contain any ASCII character from the ! (u0021) through the DEL character(u007F), including most punctuation characters, digits, and upper and lowercasedletters.
- `policies` (List of String)
- `policy_refs` (Attributes List) (see [below for nested schema](#nestedatt--spec--policy_refs))

<a id="nestedatt--spec--policy_refs"></a>
### Nested Schema for `spec.policy_refs`

Optional:

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--policy_refs--from))

<a id="nestedatt--spec--policy_refs--from"></a>
### Nested Schema for `spec.policy_refs.from`

Optional:

- `name` (String)