---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_rds_services_k8s_aws_db_cluster_parameter_group_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "rds.services.k8s.aws"
description: |-
  DBClusterParameterGroup is the Schema for the DBClusterParameterGroups API
---

# k8s_rds_services_k8s_aws_db_cluster_parameter_group_v1alpha1_manifest (Data Source)

DBClusterParameterGroup is the Schema for the DBClusterParameterGroups API

## Example Usage

```terraform
data "k8s_rds_services_k8s_aws_db_cluster_parameter_group_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) DBClusterParameterGroupSpec defines the desired state of DBClusterParameterGroup.Contains the details of an Amazon RDS DB cluster parameter group.This data type is used as a response element in the DescribeDBClusterParameterGroupsaction. (see [below for nested schema](#nestedatt--spec))

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

- `description` (String) The description for the DB cluster parameter group.
- `family` (String) The DB cluster parameter group family name. A DB cluster parameter groupcan be associated with one and only one DB cluster parameter group family,and can be applied only to a DB cluster running a database engine and engineversion compatible with that DB cluster parameter group family.Aurora MySQLExample: aurora5.6, aurora-mysql5.7, aurora-mysql8.0Aurora PostgreSQLExample: aurora-postgresql9.6RDS for MySQLExample: mysql8.0RDS for PostgreSQLExample: postgres12To list all of the available parameter group families for a DB engine, usethe following command:aws rds describe-db-engine-versions --query 'DBEngineVersions[].DBParameterGroupFamily'--engine <engine>For example, to list all of the available parameter group families for theAurora PostgreSQL DB engine, use the following command:aws rds describe-db-engine-versions --query 'DBEngineVersions[].DBParameterGroupFamily'--engine aurora-postgresqlThe output contains duplicates.The following are the valid DB engine values:   * aurora (for MySQL 5.6-compatible Aurora)   * aurora-mysql (for MySQL 5.7-compatible and MySQL 8.0-compatible Aurora)   * aurora-postgresql   * mysql   * postgres
- `name` (String) The name of the DB cluster parameter group.Constraints:   * Must not match the name of an existing DB cluster parameter group.This value is stored as a lowercase string.

Optional:

- `parameter_overrides` (Map of String)
- `parameters` (Attributes List) A list of parameters in the DB cluster parameter group to modify.Valid Values (for the application method): immediate | pending-rebootYou can use the immediate value with dynamic parameters only. You can usethe pending-reboot value for both dynamic and static parameters.When the application method is immediate, changes to dynamic parameters areapplied immediately to the DB clusters associated with the parameter group.When the application method is pending-reboot, changes to dynamic and staticparameters are applied after a reboot without failover to the DB clustersassociated with the parameter group. (see [below for nested schema](#nestedatt--spec--parameters))
- `tags` (Attributes List) Tags to assign to the DB cluster parameter group. (see [below for nested schema](#nestedatt--spec--tags))

<a id="nestedatt--spec--parameters"></a>
### Nested Schema for `spec.parameters`

Optional:

- `allowed_values` (String)
- `apply_method` (String)
- `apply_type` (String)
- `data_type` (String)
- `description` (String)
- `is_modifiable` (Boolean)
- `minimum_engine_version` (String)
- `parameter_name` (String)
- `parameter_value` (String)
- `source` (String)
- `supported_engine_modes` (List of String)


<a id="nestedatt--spec--tags"></a>
### Nested Schema for `spec.tags`

Optional:

- `key` (String)
- `value` (String)