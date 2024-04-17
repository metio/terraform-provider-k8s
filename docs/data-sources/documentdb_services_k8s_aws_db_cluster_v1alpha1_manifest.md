---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "k8s_documentdb_services_k8s_aws_db_cluster_v1alpha1_manifest Data Source - terraform-provider-k8s"
subcategory: "documentdb.services.k8s.aws"
description: |-
  DBCluster is the Schema for the DBClusters API
---

# k8s_documentdb_services_k8s_aws_db_cluster_v1alpha1_manifest (Data Source)

DBCluster is the Schema for the DBClusters API

## Example Usage

```terraform
data "k8s_documentdb_services_k8s_aws_db_cluster_v1alpha1_manifest" "example" {
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

- `spec` (Attributes) DBClusterSpec defines the desired state of DBCluster.Detailed information about a cluster. (see [below for nested schema](#nestedatt--spec))

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

- `db_cluster_identifier` (String) The cluster identifier. This parameter is stored as a lowercase string.Constraints:   * Must contain from 1 to 63 letters, numbers, or hyphens.   * The first character must be a letter.   * Cannot end with a hyphen or contain two consecutive hyphens.Example: my-cluster
- `engine` (String) The name of the database engine to be used for this cluster.Valid values: docdb

Optional:

- `availability_zones` (List of String) A list of Amazon EC2 Availability Zones that instances in the cluster canbe created in.
- `backup_retention_period` (Number) The number of days for which automated backups are retained. You must specifya minimum value of 1.Default: 1Constraints:   * Must be a value from 1 to 35.
- `db_cluster_parameter_group_name` (String) The name of the cluster parameter group to associate with this cluster.
- `db_subnet_group_name` (String) A subnet group to associate with this cluster.Constraints: Must match the name of an existing DBSubnetGroup. Must not bedefault.Example: mySubnetgroup
- `db_subnet_group_ref` (Attributes) AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api (see [below for nested schema](#nestedatt--spec--db_subnet_group_ref))
- `deletion_protection` (Boolean) Specifies whether this cluster can be deleted. If DeletionProtection is enabled,the cluster cannot be deleted unless it is modified and DeletionProtectionis disabled. DeletionProtection protects clusters from being accidentallydeleted.
- `destination_region` (String) DestinationRegion is used for presigning the request to a given region.
- `enable_cloudwatch_logs_exports` (List of String) A list of log types that need to be enabled for exporting to Amazon CloudWatchLogs. You can enable audit logs or profiler logs. For more information, seeAuditing Amazon DocumentDB Events (https://docs.aws.amazon.com/documentdb/latest/developerguide/event-auditing.html)and Profiling Amazon DocumentDB Operations (https://docs.aws.amazon.com/documentdb/latest/developerguide/profiling.html).
- `engine_version` (String) The version number of the database engine to use. The --engine-version willdefault to the latest major engine version. For production workloads, werecommend explicitly declaring this parameter with the intended major engineversion.
- `global_cluster_identifier` (String) The cluster identifier of the new global cluster.
- `kms_key_id` (String) The KMS key identifier for an encrypted cluster.The KMS key identifier is the Amazon Resource Name (ARN) for the KMS encryptionkey. If you are creating a cluster using the same Amazon Web Services accountthat owns the KMS encryption key that is used to encrypt the new cluster,you can use the KMS key alias instead of the ARN for the KMS encryption key.If an encryption key is not specified in KmsKeyId:   * If the StorageEncrypted parameter is true, Amazon DocumentDB uses your   default encryption key.KMS creates the default encryption key for your Amazon Web Services account.Your Amazon Web Services account has a different default encryption key foreach Amazon Web Services Regions.
- `kms_key_ref` (Attributes) AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api (see [below for nested schema](#nestedatt--spec--kms_key_ref))
- `master_user_password` (Attributes) The password for the master database user. This password can contain anyprintable ASCII character except forward slash (/), double quote ('), orthe 'at' symbol (@).Constraints: Must contain from 8 to 100 characters. (see [below for nested schema](#nestedatt--spec--master_user_password))
- `master_username` (String) The name of the master user for the cluster.Constraints:   * Must be from 1 to 63 letters or numbers.   * The first character must be a letter.   * Cannot be a reserved word for the chosen database engine.
- `port` (Number) The port number on which the instances in the cluster accept connections.
- `pre_signed_url` (String) Not currently supported.
- `preferred_backup_window` (String) The daily time range during which automated backups are created if automatedbackups are enabled using the BackupRetentionPeriod parameter.The default is a 30-minute window selected at random from an 8-hour blockof time for each Amazon Web Services Region.Constraints:   * Must be in the format hh24:mi-hh24:mi.   * Must be in Universal Coordinated Time (UTC).   * Must not conflict with the preferred maintenance window.   * Must be at least 30 minutes.
- `preferred_maintenance_window` (String) The weekly time range during which system maintenance can occur, in UniversalCoordinated Time (UTC).Format: ddd:hh24:mi-ddd:hh24:miThe default is a 30-minute window selected at random from an 8-hour blockof time for each Amazon Web Services Region, occurring on a random day ofthe week.Valid days: Mon, Tue, Wed, Thu, Fri, Sat, SunConstraints: Minimum 30-minute window.
- `snapshot_identifier` (String) The identifier for the snapshot or cluster snapshot to restore from.You can use either the name or the Amazon Resource Name (ARN) to specifya cluster snapshot. However, you can use only the ARN to specify a snapshot.Constraints:   * Must match the identifier of an existing snapshot.
- `source_region` (String) SourceRegion is the source region where the resource exists. This is notsent over the wire and is only used for presigning. This value should alwayshave the same region as the source ARN.
- `storage_encrypted` (Boolean) Specifies whether the cluster is encrypted.
- `storage_type` (String) The storage type to associate with the DB cluster.For information on storage types for Amazon DocumentDB clusters, see Clusterstorage configurations in the Amazon DocumentDB Developer Guide.Valid values for storage type - standard | iopt1Default value is standardWhen you create a DocumentDB DB cluster with the storage type set to iopt1,the storage type is returned in the response. The storage type isn't returnedwhen you set it to standard.
- `tags` (Attributes List) The tags to be assigned to the cluster. (see [below for nested schema](#nestedatt--spec--tags))
- `vpc_security_group_i_ds` (List of String) A list of EC2 VPC security groups to associate with this cluster.
- `vpc_security_group_refs` (Attributes List) (see [below for nested schema](#nestedatt--spec--vpc_security_group_refs))

<a id="nestedatt--spec--db_subnet_group_ref"></a>
### Nested Schema for `spec.db_subnet_group_ref`

Optional:

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--db_subnet_group_ref--from))

<a id="nestedatt--spec--db_subnet_group_ref--from"></a>
### Nested Schema for `spec.db_subnet_group_ref.from`

Optional:

- `name` (String)



<a id="nestedatt--spec--kms_key_ref"></a>
### Nested Schema for `spec.kms_key_ref`

Optional:

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--kms_key_ref--from))

<a id="nestedatt--spec--kms_key_ref--from"></a>
### Nested Schema for `spec.kms_key_ref.from`

Optional:

- `name` (String)



<a id="nestedatt--spec--master_user_password"></a>
### Nested Schema for `spec.master_user_password`

Required:

- `key` (String) Key is the key within the secret

Optional:

- `name` (String) name is unique within a namespace to reference a secret resource.
- `namespace` (String) namespace defines the space within which the secret name must be unique.


<a id="nestedatt--spec--tags"></a>
### Nested Schema for `spec.tags`

Optional:

- `key` (String)
- `value` (String)


<a id="nestedatt--spec--vpc_security_group_refs"></a>
### Nested Schema for `spec.vpc_security_group_refs`

Optional:

- `from` (Attributes) AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name) (see [below for nested schema](#nestedatt--spec--vpc_security_group_refs--from))

<a id="nestedatt--spec--vpc_security_group_refs--from"></a>
### Nested Schema for `spec.vpc_security_group_refs.from`

Optional:

- `name` (String)