---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_rds_services_k8s_aws_db_subnet_group_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "rds.services.k8s.aws"
description: |-
  DBSubnetGroup is the Schema for the DBSubnetGroups API
---

# k8s_rds_services_k8s_aws_db_subnet_group_v1alpha1_manifest (Data Source)

DBSubnetGroup is the Schema for the DBSubnetGroups API

## Example Usage

```terraform
data "k8s_rds_services_k8s_aws_db_subnet_group_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) DBSubnetGroupSpec defines the desired state of DBSubnetGroup. Contains the details of an Amazon RDS DB subnet group. This data type is used as a response element in the DescribeDBSubnetGroups action. (see [below for nested schema](#nestedatt--spec))

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

- `description` (String) The description for the DB subnet group.
- `name` (String) The name for the DB subnet group. This value is stored as a lowercase string. Constraints: * Must contain no more than 255 letters, numbers, periods, underscores, spaces, or hyphens. * Must not be default. * First character must be a letter. Example: mydbsubnetgroup

Optional:

- `subnet_i_ds` (List of String) The EC2 Subnet IDs for the DB subnet group.
- `subnet_refs` (Attributes List) (see [below for nested schema](#nestedatt--spec--subnet_refs))
- `tags` (Attributes List) Tags to assign to the DB subnet group. (see [below for nested schema](#nestedatt--spec--tags))

<a id="nestedatt--spec--subnet_refs"></a>
### Nested Schema for `spec.subnet_refs`

Optional:

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--subnet_refs--from))

<a id="nestedatt--spec--subnet_refs--from"></a>
### Nested Schema for `spec.subnet_refs.from`

Optional:

- `name` (String)
- `namespace` (String)



<a id="nestedatt--spec--tags"></a>
### Nested Schema for `spec.tags`

Optional:

- `key` (String)
- `value` (String)
