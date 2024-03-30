/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kafka_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &KafkaServicesK8SAwsClusterV1Alpha1Manifest{}
)

func NewKafkaServicesK8SAwsClusterV1Alpha1Manifest() datasource.DataSource {
	return &KafkaServicesK8SAwsClusterV1Alpha1Manifest{}
}

type KafkaServicesK8SAwsClusterV1Alpha1Manifest struct{}

type KafkaServicesK8SAwsClusterV1Alpha1ManifestData struct {
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
		BrokerNodeGroupInfo *struct {
			BrokerAZDistribution *string   `tfsdk:"broker_az_distribution" json:"brokerAZDistribution,omitempty"`
			ClientSubnets        *[]string `tfsdk:"client_subnets" json:"clientSubnets,omitempty"`
			ConnectivityInfo     *struct {
				PublicAccess *struct {
					Type_ *string `tfsdk:"type_" json:"type_,omitempty"`
				} `tfsdk:"public_access" json:"publicAccess,omitempty"`
			} `tfsdk:"connectivity_info" json:"connectivityInfo,omitempty"`
			InstanceType   *string   `tfsdk:"instance_type" json:"instanceType,omitempty"`
			SecurityGroups *[]string `tfsdk:"security_groups" json:"securityGroups,omitempty"`
			StorageInfo    *struct {
				EbsStorageInfo *struct {
					ProvisionedThroughput *struct {
						Enabled          *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
						VolumeThroughput *int64 `tfsdk:"volume_throughput" json:"volumeThroughput,omitempty"`
					} `tfsdk:"provisioned_throughput" json:"provisionedThroughput,omitempty"`
					VolumeSize *int64 `tfsdk:"volume_size" json:"volumeSize,omitempty"`
				} `tfsdk:"ebs_storage_info" json:"ebsStorageInfo,omitempty"`
			} `tfsdk:"storage_info" json:"storageInfo,omitempty"`
		} `tfsdk:"broker_node_group_info" json:"brokerNodeGroupInfo,omitempty"`
		ClientAuthentication *struct {
			Sasl *struct {
				Iam *struct {
					Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
				} `tfsdk:"iam" json:"iam,omitempty"`
				Scram *struct {
					Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
				} `tfsdk:"scram" json:"scram,omitempty"`
			} `tfsdk:"sasl" json:"sasl,omitempty"`
			Tls *struct {
				CertificateAuthorityARNList *[]string `tfsdk:"certificate_authority_arn_list" json:"certificateAuthorityARNList,omitempty"`
				Enabled                     *bool     `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			Unauthenticated *struct {
				Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"unauthenticated" json:"unauthenticated,omitempty"`
		} `tfsdk:"client_authentication" json:"clientAuthentication,omitempty"`
		ConfigurationInfo *struct {
			Arn      *string `tfsdk:"arn" json:"arn,omitempty"`
			Revision *int64  `tfsdk:"revision" json:"revision,omitempty"`
		} `tfsdk:"configuration_info" json:"configurationInfo,omitempty"`
		EncryptionInfo *struct {
			EncryptionAtRest *struct {
				DataVolumeKMSKeyID *string `tfsdk:"data_volume_kms_key_id" json:"dataVolumeKMSKeyID,omitempty"`
			} `tfsdk:"encryption_at_rest" json:"encryptionAtRest,omitempty"`
			EncryptionInTransit *struct {
				ClientBroker *string `tfsdk:"client_broker" json:"clientBroker,omitempty"`
				InCluster    *bool   `tfsdk:"in_cluster" json:"inCluster,omitempty"`
			} `tfsdk:"encryption_in_transit" json:"encryptionInTransit,omitempty"`
		} `tfsdk:"encryption_info" json:"encryptionInfo,omitempty"`
		EnhancedMonitoring *string `tfsdk:"enhanced_monitoring" json:"enhancedMonitoring,omitempty"`
		KafkaVersion       *string `tfsdk:"kafka_version" json:"kafkaVersion,omitempty"`
		LoggingInfo        *struct {
			BrokerLogs *struct {
				CloudWatchLogs *struct {
					Enabled  *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					LogGroup *string `tfsdk:"log_group" json:"logGroup,omitempty"`
				} `tfsdk:"cloud_watch_logs" json:"cloudWatchLogs,omitempty"`
				Firehose *struct {
					DeliveryStream *string `tfsdk:"delivery_stream" json:"deliveryStream,omitempty"`
					Enabled        *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				} `tfsdk:"firehose" json:"firehose,omitempty"`
				S3 *struct {
					Bucket  *string `tfsdk:"bucket" json:"bucket,omitempty"`
					Enabled *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Prefix  *string `tfsdk:"prefix" json:"prefix,omitempty"`
				} `tfsdk:"s3" json:"s3,omitempty"`
			} `tfsdk:"broker_logs" json:"brokerLogs,omitempty"`
		} `tfsdk:"logging_info" json:"loggingInfo,omitempty"`
		Name                *string `tfsdk:"name" json:"name,omitempty"`
		NumberOfBrokerNodes *int64  `tfsdk:"number_of_broker_nodes" json:"numberOfBrokerNodes,omitempty"`
		OpenMonitoring      *struct {
			Prometheus *struct {
				JmxExporter *struct {
					EnabledInBroker *bool `tfsdk:"enabled_in_broker" json:"enabledInBroker,omitempty"`
				} `tfsdk:"jmx_exporter" json:"jmxExporter,omitempty"`
				NodeExporter *struct {
					EnabledInBroker *bool `tfsdk:"enabled_in_broker" json:"enabledInBroker,omitempty"`
				} `tfsdk:"node_exporter" json:"nodeExporter,omitempty"`
			} `tfsdk:"prometheus" json:"prometheus,omitempty"`
		} `tfsdk:"open_monitoring" json:"openMonitoring,omitempty"`
		StorageMode *string            `tfsdk:"storage_mode" json:"storageMode,omitempty"`
		Tags        *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KafkaServicesK8SAwsClusterV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kafka_services_k8s_aws_cluster_v1alpha1_manifest"
}

func (r *KafkaServicesK8SAwsClusterV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "ClusterSpec defines the desired state of Cluster.Returns information about a cluster of either the provisioned or the serverlesstype.",
				MarkdownDescription: "ClusterSpec defines the desired state of Cluster.Returns information about a cluster of either the provisioned or the serverlesstype.",
				Attributes: map[string]schema.Attribute{
					"broker_node_group_info": schema.SingleNestedAttribute{
						Description:         "Information about the brokers.",
						MarkdownDescription: "Information about the brokers.",
						Attributes: map[string]schema.Attribute{
							"broker_az_distribution": schema.StringAttribute{
								Description:         "The distribution of broker nodes across Availability Zones. By default, brokernodes are distributed among the Availability Zones of your Region. Currently,the only supported value is DEFAULT. You can either specify this value explicitlyor leave it out.",
								MarkdownDescription: "The distribution of broker nodes across Availability Zones. By default, brokernodes are distributed among the Availability Zones of your Region. Currently,the only supported value is DEFAULT. You can either specify this value explicitlyor leave it out.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_subnets": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"connectivity_info": schema.SingleNestedAttribute{
								Description:         "Information about the broker access configuration.",
								MarkdownDescription: "Information about the broker access configuration.",
								Attributes: map[string]schema.Attribute{
									"public_access": schema.SingleNestedAttribute{
										Description:         "Broker public access control.",
										MarkdownDescription: "Broker public access control.",
										Attributes: map[string]schema.Attribute{
											"type_": schema.StringAttribute{
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

							"instance_type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"security_groups": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"storage_info": schema.SingleNestedAttribute{
								Description:         "Contains information about storage volumes attached to MSK broker nodes.",
								MarkdownDescription: "Contains information about storage volumes attached to MSK broker nodes.",
								Attributes: map[string]schema.Attribute{
									"ebs_storage_info": schema.SingleNestedAttribute{
										Description:         "Contains information about the EBS storage volumes attached to Apache Kafkabroker nodes.",
										MarkdownDescription: "Contains information about the EBS storage volumes attached to Apache Kafkabroker nodes.",
										Attributes: map[string]schema.Attribute{
											"provisioned_throughput": schema.SingleNestedAttribute{
												Description:         "Contains information about provisioned throughput for EBS storage volumesattached to kafka broker nodes.",
												MarkdownDescription: "Contains information about provisioned throughput for EBS storage volumesattached to kafka broker nodes.",
												Attributes: map[string]schema.Attribute{
													"enabled": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"volume_throughput": schema.Int64Attribute{
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

											"volume_size": schema.Int64Attribute{
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
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"client_authentication": schema.SingleNestedAttribute{
						Description:         "Includes all client authentication related information.",
						MarkdownDescription: "Includes all client authentication related information.",
						Attributes: map[string]schema.Attribute{
							"sasl": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"iam": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
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

									"scram": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
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

							"tls": schema.SingleNestedAttribute{
								Description:         "Details for client authentication using TLS.",
								MarkdownDescription: "Details for client authentication using TLS.",
								Attributes: map[string]schema.Attribute{
									"certificate_authority_arn_list": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
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

							"unauthenticated": schema.SingleNestedAttribute{
								Description:         "Contains information about unauthenticated traffic to the cluster.",
								MarkdownDescription: "Contains information about unauthenticated traffic to the cluster.",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
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

					"configuration_info": schema.SingleNestedAttribute{
						Description:         "Represents the configuration that you want MSK to use for the cluster.",
						MarkdownDescription: "Represents the configuration that you want MSK to use for the cluster.",
						Attributes: map[string]schema.Attribute{
							"arn": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"revision": schema.Int64Attribute{
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

					"encryption_info": schema.SingleNestedAttribute{
						Description:         "Includes all encryption-related information.",
						MarkdownDescription: "Includes all encryption-related information.",
						Attributes: map[string]schema.Attribute{
							"encryption_at_rest": schema.SingleNestedAttribute{
								Description:         "The data-volume encryption details.",
								MarkdownDescription: "The data-volume encryption details.",
								Attributes: map[string]schema.Attribute{
									"data_volume_kms_key_id": schema.StringAttribute{
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

							"encryption_in_transit": schema.SingleNestedAttribute{
								Description:         "The settings for encrypting data in transit.",
								MarkdownDescription: "The settings for encrypting data in transit.",
								Attributes: map[string]schema.Attribute{
									"client_broker": schema.StringAttribute{
										Description:         "Client-broker encryption in transit setting.",
										MarkdownDescription: "Client-broker encryption in transit setting.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"in_cluster": schema.BoolAttribute{
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

					"enhanced_monitoring": schema.StringAttribute{
						Description:         "Specifies the level of monitoring for the MSK cluster. The possible valuesare DEFAULT, PER_BROKER, PER_TOPIC_PER_BROKER, and PER_TOPIC_PER_PARTITION.",
						MarkdownDescription: "Specifies the level of monitoring for the MSK cluster. The possible valuesare DEFAULT, PER_BROKER, PER_TOPIC_PER_BROKER, and PER_TOPIC_PER_PARTITION.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kafka_version": schema.StringAttribute{
						Description:         "The version of Apache Kafka.",
						MarkdownDescription: "The version of Apache Kafka.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"logging_info": schema.SingleNestedAttribute{
						Description:         "LoggingInfo details.",
						MarkdownDescription: "LoggingInfo details.",
						Attributes: map[string]schema.Attribute{
							"broker_logs": schema.SingleNestedAttribute{
								Description:         "The broker logs configuration for this MSK cluster.",
								MarkdownDescription: "The broker logs configuration for this MSK cluster.",
								Attributes: map[string]schema.Attribute{
									"cloud_watch_logs": schema.SingleNestedAttribute{
										Description:         "Details of the CloudWatch Logs destination for broker logs.",
										MarkdownDescription: "Details of the CloudWatch Logs destination for broker logs.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"log_group": schema.StringAttribute{
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

									"firehose": schema.SingleNestedAttribute{
										Description:         "Firehose details for BrokerLogs.",
										MarkdownDescription: "Firehose details for BrokerLogs.",
										Attributes: map[string]schema.Attribute{
											"delivery_stream": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
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

									"s3": schema.SingleNestedAttribute{
										Description:         "The details of the Amazon S3 destination for broker logs.",
										MarkdownDescription: "The details of the Amazon S3 destination for broker logs.",
										Attributes: map[string]schema.Attribute{
											"bucket": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"prefix": schema.StringAttribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"name": schema.StringAttribute{
						Description:         "The name of the cluster.",
						MarkdownDescription: "The name of the cluster.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"number_of_broker_nodes": schema.Int64Attribute{
						Description:         "The number of Apache Kafka broker nodes in the Amazon MSK cluster.",
						MarkdownDescription: "The number of Apache Kafka broker nodes in the Amazon MSK cluster.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"open_monitoring": schema.SingleNestedAttribute{
						Description:         "The settings for open monitoring.",
						MarkdownDescription: "The settings for open monitoring.",
						Attributes: map[string]schema.Attribute{
							"prometheus": schema.SingleNestedAttribute{
								Description:         "Prometheus settings.",
								MarkdownDescription: "Prometheus settings.",
								Attributes: map[string]schema.Attribute{
									"jmx_exporter": schema.SingleNestedAttribute{
										Description:         "Indicates whether you want to enable or disable the JMX Exporter.",
										MarkdownDescription: "Indicates whether you want to enable or disable the JMX Exporter.",
										Attributes: map[string]schema.Attribute{
											"enabled_in_broker": schema.BoolAttribute{
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

									"node_exporter": schema.SingleNestedAttribute{
										Description:         "Indicates whether you want to enable or disable the Node Exporter.",
										MarkdownDescription: "Indicates whether you want to enable or disable the Node Exporter.",
										Attributes: map[string]schema.Attribute{
											"enabled_in_broker": schema.BoolAttribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"storage_mode": schema.StringAttribute{
						Description:         "This controls storage mode for supported storage tiers.",
						MarkdownDescription: "This controls storage mode for supported storage tiers.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.MapAttribute{
						Description:         "Create tags when creating the cluster.",
						MarkdownDescription: "Create tags when creating the cluster.",
						ElementType:         types.StringType,
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

func (r *KafkaServicesK8SAwsClusterV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kafka_services_k8s_aws_cluster_v1alpha1_manifest")

	var model KafkaServicesK8SAwsClusterV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kafka.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Cluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
