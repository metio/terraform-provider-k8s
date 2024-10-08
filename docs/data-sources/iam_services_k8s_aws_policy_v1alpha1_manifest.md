---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_iam_services_k8s_aws_policy_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "iam.services.k8s.aws"
description: |-
  Policy is the Schema for the Policies API
---

# k8s_iam_services_k8s_aws_policy_v1alpha1_manifest (Data Source)

Policy is the Schema for the Policies API

## Example Usage

```terraform
data "k8s_iam_services_k8s_aws_policy_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) PolicySpec defines the desired state of Policy. Contains information about a managed policy. This data type is used as a response element in the CreatePolicy, GetPolicy, and ListPolicies operations. For more information about managed policies, refer to Managed policies and inline policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html) in the IAM User Guide. (see [below for nested schema](#nestedatt--spec))

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

- `name` (String) The friendly name of the policy. IAM user, group, role, and policy names must be unique within the account. Names are not distinguished by case. For example, you cannot create resources named both 'MyResource' and 'myresource'.
- `policy_document` (String) The JSON policy document that you want to use as the content for the new policy. You must provide policies in JSON format in IAM. However, for CloudFormation templates formatted in YAML, you can provide the policy in JSON or YAML format. CloudFormation always converts a YAML policy to JSON format before submitting it to IAM. The maximum length of the policy document that you can pass in this operation, including whitespace, is listed below. To view the maximum character counts of a managed policy with no whitespaces, see IAM and STS character quotas (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html#reference_iam-quotas-entity-length). To learn more about JSON policy grammar, see Grammar of the IAM JSON policy language (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_grammar.html) in the IAM User Guide. The regex pattern (http://wikipedia.org/wiki/regex) used to validate this parameter is a string of characters consisting of the following: * Any printable ASCII character ranging from the space character (u0020) through the end of the ASCII character range * The printable characters in the Basic Latin and Latin-1 Supplement character set (through u00FF) * The special characters tab (u0009), line feed (u000A), and carriage return (u000D)

Optional:

- `description` (String) A friendly description of the policy. Typically used to store information about the permissions defined in the policy. For example, 'Grants access to production DynamoDB tables.' The policy description is immutable. After a value is assigned, it cannot be changed.
- `path` (String) The path for the policy. For more information about paths, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the IAM User Guide. This parameter is optional. If it is not included, it defaults to a slash (/). This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of characters consisting of either a forward slash (/) by itself or a string that must begin and end with forward slashes. In addition, it can contain any ASCII character from the ! (u0021) through the DEL character (u007F), including most punctuation characters, digits, and upper and lowercased letters. You cannot use an asterisk (*) in the path name.
- `tags` (Attributes List) A list of tags that you want to attach to the new IAM customer managed policy. Each tag consists of a key name and an associated value. For more information about tagging, see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the IAM User Guide. If any one of the tags is invalid or if you exceed the allowed maximum number of tags, then the entire request fails and the resource is not created. (see [below for nested schema](#nestedatt--spec--tags))

<a id="nestedatt--spec--tags"></a>
### Nested Schema for `spec.tags`

Optional:

- `key` (String)
- `value` (String)
