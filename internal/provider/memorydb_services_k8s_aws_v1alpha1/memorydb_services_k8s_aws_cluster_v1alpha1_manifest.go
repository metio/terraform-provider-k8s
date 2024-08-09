/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package memorydb_services_k8s_aws_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &MemorydbServicesK8SAwsClusterV1Alpha1Manifest{}
)

func NewMemorydbServicesK8SAwsClusterV1Alpha1Manifest() datasource.DataSource {
	return &MemorydbServicesK8SAwsClusterV1Alpha1Manifest{}
}

type MemorydbServicesK8SAwsClusterV1Alpha1Manifest struct{}

type MemorydbServicesK8SAwsClusterV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AclName *string `tfsdk:"acl_name" json:"aclName,omitempty"`
		AclRef  *struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"acl_ref" json:"aclRef,omitempty"`
		AutoMinorVersionUpgrade *bool   `tfsdk:"auto_minor_version_upgrade" json:"autoMinorVersionUpgrade,omitempty"`
		Description             *string `tfsdk:"description" json:"description,omitempty"`
		EngineVersion           *string `tfsdk:"engine_version" json:"engineVersion,omitempty"`
		KmsKeyID                *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
		MaintenanceWindow       *string `tfsdk:"maintenance_window" json:"maintenanceWindow,omitempty"`
		Name                    *string `tfsdk:"name" json:"name,omitempty"`
		NodeType                *string `tfsdk:"node_type" json:"nodeType,omitempty"`
		NumReplicasPerShard     *int64  `tfsdk:"num_replicas_per_shard" json:"numReplicasPerShard,omitempty"`
		NumShards               *int64  `tfsdk:"num_shards" json:"numShards,omitempty"`
		ParameterGroupName      *string `tfsdk:"parameter_group_name" json:"parameterGroupName,omitempty"`
		ParameterGroupRef       *struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"parameter_group_ref" json:"parameterGroupRef,omitempty"`
		Port              *int64    `tfsdk:"port" json:"port,omitempty"`
		SecurityGroupIDs  *[]string `tfsdk:"security_group_i_ds" json:"securityGroupIDs,omitempty"`
		SecurityGroupRefs *[]struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"security_group_refs" json:"securityGroupRefs,omitempty"`
		SnapshotARNs *[]string `tfsdk:"snapshot_ar_ns" json:"snapshotARNs,omitempty"`
		SnapshotName *string   `tfsdk:"snapshot_name" json:"snapshotName,omitempty"`
		SnapshotRef  *struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"snapshot_ref" json:"snapshotRef,omitempty"`
		SnapshotRetentionLimit *int64  `tfsdk:"snapshot_retention_limit" json:"snapshotRetentionLimit,omitempty"`
		SnapshotWindow         *string `tfsdk:"snapshot_window" json:"snapshotWindow,omitempty"`
		SnsTopicARN            *string `tfsdk:"sns_topic_arn" json:"snsTopicARN,omitempty"`
		SnsTopicRef            *struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"sns_topic_ref" json:"snsTopicRef,omitempty"`
		SubnetGroupName *string `tfsdk:"subnet_group_name" json:"subnetGroupName,omitempty"`
		SubnetGroupRef  *struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"subnet_group_ref" json:"subnetGroupRef,omitempty"`
		Tags *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		TlsEnabled *bool `tfsdk:"tls_enabled" json:"tlsEnabled,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MemorydbServicesK8SAwsClusterV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_memorydb_services_k8s_aws_cluster_v1alpha1_manifest"
}

func (r *MemorydbServicesK8SAwsClusterV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Cluster is the Schema for the Clusters API",
		MarkdownDescription: "Cluster is the Schema for the Clusters API",
		Attributes: map[string]schema.Attribute{
			"yaml": schema.StringAttribute{
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"metadata": schema.SingleNestedAttribute{
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Optional:            false,
				Computed:            false,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.NameValidator(),
							stringvalidator.LengthAtLeast(1),
						},
					},

					"namespace": schema.StringAttribute{
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.NameValidator(),
							stringvalidator.LengthAtLeast(1),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Map{
							validators.LabelValidator(),
						},
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "ClusterSpec defines the desired state of Cluster.Contains all of the attributes of a specific cluster.",
				MarkdownDescription: "ClusterSpec defines the desired state of Cluster.Contains all of the attributes of a specific cluster.",
				Attributes: map[string]schema.Attribute{
					"acl_name": schema.StringAttribute{
						Description:         "The name of the Access Control List to associate with the cluster.",
						MarkdownDescription: "The name of the Access Control List to associate with the cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"acl_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"auto_minor_version_upgrade": schema.BoolAttribute{
						Description:         "When set to true, the cluster will automatically receive minor engine versionupgrades after launch.",
						MarkdownDescription: "When set to true, the cluster will automatically receive minor engine versionupgrades after launch.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"description": schema.StringAttribute{
						Description:         "An optional description of the cluster.",
						MarkdownDescription: "An optional description of the cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"engine_version": schema.StringAttribute{
						Description:         "The version number of the Redis engine to be used for the cluster.",
						MarkdownDescription: "The version number of the Redis engine to be used for the cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kms_key_id": schema.StringAttribute{
						Description:         "The ID of the KMS key used to encrypt the cluster.",
						MarkdownDescription: "The ID of the KMS key used to encrypt the cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"maintenance_window": schema.StringAttribute{
						Description:         "Specifies the weekly time range during which maintenance on the cluster isperformed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi(24H Clock UTC). The minimum maintenance window is a 60 minute period.",
						MarkdownDescription: "Specifies the weekly time range during which maintenance on the cluster isperformed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi(24H Clock UTC). The minimum maintenance window is a 60 minute period.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "The name of the cluster. This value must be unique as it also serves as thecluster identifier.",
						MarkdownDescription: "The name of the cluster. This value must be unique as it also serves as thecluster identifier.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"node_type": schema.StringAttribute{
						Description:         "The compute and memory capacity of the nodes in the cluster.",
						MarkdownDescription: "The compute and memory capacity of the nodes in the cluster.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"num_replicas_per_shard": schema.Int64Attribute{
						Description:         "The number of replicas to apply to each shard. The default value is 1. Themaximum is 5.",
						MarkdownDescription: "The number of replicas to apply to each shard. The default value is 1. Themaximum is 5.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"num_shards": schema.Int64Attribute{
						Description:         "The number of shards the cluster will contain. The default value is 1.",
						MarkdownDescription: "The number of shards the cluster will contain. The default value is 1.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"parameter_group_name": schema.StringAttribute{
						Description:         "The name of the parameter group associated with the cluster.",
						MarkdownDescription: "The name of the parameter group associated with the cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"parameter_group_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"port": schema.Int64Attribute{
						Description:         "The port number on which each of the nodes accepts connections.",
						MarkdownDescription: "The port number on which each of the nodes accepts connections.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"security_group_i_ds": schema.ListAttribute{
						Description:         "A list of security group names to associate with this cluster.",
						MarkdownDescription: "A list of security group names to associate with this cluster.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"security_group_refs": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"from": schema.SingleNestedAttribute{
									Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
									MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"snapshot_ar_ns": schema.ListAttribute{
						Description:         "A list of Amazon Resource Names (ARN) that uniquely identify the RDB snapshotfiles stored in Amazon S3. The snapshot files are used to populate the newcluster. The Amazon S3 object name in the ARN cannot contain any commas.",
						MarkdownDescription: "A list of Amazon Resource Names (ARN) that uniquely identify the RDB snapshotfiles stored in Amazon S3. The snapshot files are used to populate the newcluster. The Amazon S3 object name in the ARN cannot contain any commas.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"snapshot_name": schema.StringAttribute{
						Description:         "The name of a snapshot from which to restore data into the new cluster. Thesnapshot status changes to restoring while the new cluster is being created.",
						MarkdownDescription: "The name of a snapshot from which to restore data into the new cluster. Thesnapshot status changes to restoring while the new cluster is being created.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"snapshot_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"snapshot_retention_limit": schema.Int64Attribute{
						Description:         "The number of days for which MemoryDB retains automatic snapshots beforedeleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshotthat was taken today is retained for 5 days before being deleted.",
						MarkdownDescription: "The number of days for which MemoryDB retains automatic snapshots beforedeleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshotthat was taken today is retained for 5 days before being deleted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"snapshot_window": schema.StringAttribute{
						Description:         "The daily time range (in UTC) during which MemoryDB begins taking a dailysnapshot of your shard.Example: 05:00-09:00If you do not specify this parameter, MemoryDB automatically chooses an appropriatetime range.",
						MarkdownDescription: "The daily time range (in UTC) during which MemoryDB begins taking a dailysnapshot of your shard.Example: 05:00-09:00If you do not specify this parameter, MemoryDB automatically chooses an appropriatetime range.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"sns_topic_arn": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) of the Amazon Simple Notification Service(SNS) topic to which notifications are sent.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of the Amazon Simple Notification Service(SNS) topic to which notifications are sent.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"sns_topic_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"subnet_group_name": schema.StringAttribute{
						Description:         "The name of the subnet group to be used for the cluster.",
						MarkdownDescription: "The name of the subnet group to be used for the cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"subnet_group_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "A list of tags to be added to this resource. Tags are comma-separated key,valuepairs (e.g. Key=myKey, Value=myKeyValue. You can include multiple tags asshown following: Key=myKey, Value=myKeyValue Key=mySecondKey, Value=mySecondKeyValue.",
						MarkdownDescription: "A list of tags to be added to this resource. Tags are comma-separated key,valuepairs (e.g. Key=myKey, Value=myKeyValue. You can include multiple tags asshown following: Key=myKey, Value=myKeyValue Key=mySecondKey, Value=mySecondKeyValue.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tls_enabled": schema.BoolAttribute{
						Description:         "A flag to enable in-transit encryption on the cluster.",
						MarkdownDescription: "A flag to enable in-transit encryption on the cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *MemorydbServicesK8SAwsClusterV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_memorydb_services_k8s_aws_cluster_v1alpha1_manifest")

	var model MemorydbServicesK8SAwsClusterV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("memorydb.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Cluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
