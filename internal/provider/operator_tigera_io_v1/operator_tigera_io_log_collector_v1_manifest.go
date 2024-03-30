/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_tigera_io_v1

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
	_ datasource.DataSource = &OperatorTigeraIoLogCollectorV1Manifest{}
)

func NewOperatorTigeraIoLogCollectorV1Manifest() datasource.DataSource {
	return &OperatorTigeraIoLogCollectorV1Manifest{}
}

type OperatorTigeraIoLogCollectorV1Manifest struct{}

type OperatorTigeraIoLogCollectorV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AdditionalSources *struct {
			EksCloudwatchLog *struct {
				FetchInterval *int64  `tfsdk:"fetch_interval" json:"fetchInterval,omitempty"`
				GroupName     *string `tfsdk:"group_name" json:"groupName,omitempty"`
				Region        *string `tfsdk:"region" json:"region,omitempty"`
				StreamPrefix  *string `tfsdk:"stream_prefix" json:"streamPrefix,omitempty"`
			} `tfsdk:"eks_cloudwatch_log" json:"eksCloudwatchLog,omitempty"`
		} `tfsdk:"additional_sources" json:"additionalSources,omitempty"`
		AdditionalStores *struct {
			S3 *struct {
				BucketName *string `tfsdk:"bucket_name" json:"bucketName,omitempty"`
				BucketPath *string `tfsdk:"bucket_path" json:"bucketPath,omitempty"`
				Region     *string `tfsdk:"region" json:"region,omitempty"`
			} `tfsdk:"s3" json:"s3,omitempty"`
			Splunk *struct {
				Endpoint *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
			} `tfsdk:"splunk" json:"splunk,omitempty"`
			Syslog *struct {
				Encryption *string   `tfsdk:"encryption" json:"encryption,omitempty"`
				Endpoint   *string   `tfsdk:"endpoint" json:"endpoint,omitempty"`
				LogTypes   *[]string `tfsdk:"log_types" json:"logTypes,omitempty"`
				PacketSize *int64    `tfsdk:"packet_size" json:"packetSize,omitempty"`
			} `tfsdk:"syslog" json:"syslog,omitempty"`
		} `tfsdk:"additional_stores" json:"additionalStores,omitempty"`
		CollectProcessPath        *string `tfsdk:"collect_process_path" json:"collectProcessPath,omitempty"`
		EksLogForwarderDeployment *struct {
			Spec *struct {
				Template *struct {
					Spec *struct {
						Containers *[]struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Resources *struct {
								Claims *[]struct {
									Name *string `tfsdk:"name" json:"name,omitempty"`
								} `tfsdk:"claims" json:"claims,omitempty"`
								Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
								Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
							} `tfsdk:"resources" json:"resources,omitempty"`
						} `tfsdk:"containers" json:"containers,omitempty"`
						InitContainers *[]struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Resources *struct {
								Claims *[]struct {
									Name *string `tfsdk:"name" json:"name,omitempty"`
								} `tfsdk:"claims" json:"claims,omitempty"`
								Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
								Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
							} `tfsdk:"resources" json:"resources,omitempty"`
						} `tfsdk:"init_containers" json:"initContainers,omitempty"`
					} `tfsdk:"spec" json:"spec,omitempty"`
				} `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"eks_log_forwarder_deployment" json:"eksLogForwarderDeployment,omitempty"`
		FluentdDaemonSet *struct {
			Spec *struct {
				Template *struct {
					Spec *struct {
						Containers *[]struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Resources *struct {
								Claims *[]struct {
									Name *string `tfsdk:"name" json:"name,omitempty"`
								} `tfsdk:"claims" json:"claims,omitempty"`
								Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
								Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
							} `tfsdk:"resources" json:"resources,omitempty"`
						} `tfsdk:"containers" json:"containers,omitempty"`
						InitContainers *[]struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Resources *struct {
								Claims *[]struct {
									Name *string `tfsdk:"name" json:"name,omitempty"`
								} `tfsdk:"claims" json:"claims,omitempty"`
								Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
								Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
							} `tfsdk:"resources" json:"resources,omitempty"`
						} `tfsdk:"init_containers" json:"initContainers,omitempty"`
					} `tfsdk:"spec" json:"spec,omitempty"`
				} `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"fluentd_daemon_set" json:"fluentdDaemonSet,omitempty"`
		MultiTenantManagementClusterNamespace *string `tfsdk:"multi_tenant_management_cluster_namespace" json:"multiTenantManagementClusterNamespace,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorTigeraIoLogCollectorV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_tigera_io_log_collector_v1_manifest"
}

func (r *OperatorTigeraIoLogCollectorV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "LogCollector installs the components required for Tigera flow and DNS log collection. At most one instance of this resource is supported. It must be named 'tigera-secure'. When created, this installs fluentd on all nodes configured to collect Tigera log data and export it to Tigera's Elasticsearch cluster as well as any additionally configured destinations.",
		MarkdownDescription: "LogCollector installs the components required for Tigera flow and DNS log collection. At most one instance of this resource is supported. It must be named 'tigera-secure'. When created, this installs fluentd on all nodes configured to collect Tigera log data and export it to Tigera's Elasticsearch cluster as well as any additionally configured destinations.",
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
				Description:         "Specification of the desired state for Tigera log collection.",
				MarkdownDescription: "Specification of the desired state for Tigera log collection.",
				Attributes: map[string]schema.Attribute{
					"additional_sources": schema.SingleNestedAttribute{
						Description:         "Configuration for importing audit logs from managed kubernetes cluster log sources.",
						MarkdownDescription: "Configuration for importing audit logs from managed kubernetes cluster log sources.",
						Attributes: map[string]schema.Attribute{
							"eks_cloudwatch_log": schema.SingleNestedAttribute{
								Description:         "If specified with EKS Provider in Installation, enables fetching EKS audit logs.",
								MarkdownDescription: "If specified with EKS Provider in Installation, enables fetching EKS audit logs.",
								Attributes: map[string]schema.Attribute{
									"fetch_interval": schema.Int64Attribute{
										Description:         "Cloudwatch audit logs fetching interval in seconds. Default: 60",
										MarkdownDescription: "Cloudwatch audit logs fetching interval in seconds. Default: 60",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"group_name": schema.StringAttribute{
										Description:         "Cloudwatch log-group name containing EKS audit logs.",
										MarkdownDescription: "Cloudwatch log-group name containing EKS audit logs.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"region": schema.StringAttribute{
										Description:         "AWS Region EKS cluster is hosted in.",
										MarkdownDescription: "AWS Region EKS cluster is hosted in.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"stream_prefix": schema.StringAttribute{
										Description:         "Prefix of Cloudwatch log stream containing EKS audit logs in the log-group. Default: kube-apiserver-audit-",
										MarkdownDescription: "Prefix of Cloudwatch log stream containing EKS audit logs in the log-group. Default: kube-apiserver-audit-",
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

					"additional_stores": schema.SingleNestedAttribute{
						Description:         "Configuration for exporting flow, audit, and DNS logs to external storage.",
						MarkdownDescription: "Configuration for exporting flow, audit, and DNS logs to external storage.",
						Attributes: map[string]schema.Attribute{
							"s3": schema.SingleNestedAttribute{
								Description:         "If specified, enables exporting of flow, audit, and DNS logs to Amazon S3 storage.",
								MarkdownDescription: "If specified, enables exporting of flow, audit, and DNS logs to Amazon S3 storage.",
								Attributes: map[string]schema.Attribute{
									"bucket_name": schema.StringAttribute{
										Description:         "Name of the S3 bucket to send logs",
										MarkdownDescription: "Name of the S3 bucket to send logs",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"bucket_path": schema.StringAttribute{
										Description:         "Path in the S3 bucket where to send logs",
										MarkdownDescription: "Path in the S3 bucket where to send logs",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"region": schema.StringAttribute{
										Description:         "AWS Region of the S3 bucket",
										MarkdownDescription: "AWS Region of the S3 bucket",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"splunk": schema.SingleNestedAttribute{
								Description:         "If specified, enables exporting of flow, audit, and DNS logs to splunk.",
								MarkdownDescription: "If specified, enables exporting of flow, audit, and DNS logs to splunk.",
								Attributes: map[string]schema.Attribute{
									"endpoint": schema.StringAttribute{
										Description:         "Location for splunk's http event collector end point. example 'https://1.2.3.4:8088'",
										MarkdownDescription: "Location for splunk's http event collector end point. example 'https://1.2.3.4:8088'",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"syslog": schema.SingleNestedAttribute{
								Description:         "If specified, enables exporting of flow, audit, and DNS logs to syslog.",
								MarkdownDescription: "If specified, enables exporting of flow, audit, and DNS logs to syslog.",
								Attributes: map[string]schema.Attribute{
									"encryption": schema.StringAttribute{
										Description:         "Encryption configures traffic encryption to the Syslog server. Default: None",
										MarkdownDescription: "Encryption configures traffic encryption to the Syslog server. Default: None",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("None", "TLS"),
										},
									},

									"endpoint": schema.StringAttribute{
										Description:         "Location of the syslog server. example: tcp://1.2.3.4:601",
										MarkdownDescription: "Location of the syslog server. example: tcp://1.2.3.4:601",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"log_types": schema.ListAttribute{
										Description:         "If no values are provided, the list will be updated to include log types Audit, DNS and Flows. Default: Audit, DNS, Flows",
										MarkdownDescription: "If no values are provided, the list will be updated to include log types Audit, DNS and Flows. Default: Audit, DNS, Flows",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"packet_size": schema.Int64Attribute{
										Description:         "PacketSize defines the maximum size of packets to send to syslog. In general this is only needed if you notice long logs being truncated. Default: 1024",
										MarkdownDescription: "PacketSize defines the maximum size of packets to send to syslog. In general this is only needed if you notice long logs being truncated. Default: 1024",
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

					"collect_process_path": schema.StringAttribute{
						Description:         "Configuration for enabling/disabling process path collection in flowlogs. If Enabled, this feature sets hostPID to true in order to read process cmdline. Default: Enabled",
						MarkdownDescription: "Configuration for enabling/disabling process path collection in flowlogs. If Enabled, this feature sets hostPID to true in order to read process cmdline. Default: Enabled",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Enabled", "Disabled"),
						},
					},

					"eks_log_forwarder_deployment": schema.SingleNestedAttribute{
						Description:         "EKSLogForwarderDeployment configures the EKSLogForwarderDeployment Deployment.",
						MarkdownDescription: "EKSLogForwarderDeployment configures the EKSLogForwarderDeployment Deployment.",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "Spec is the specification of the EKSLogForwarder Deployment.",
								MarkdownDescription: "Spec is the specification of the EKSLogForwarder Deployment.",
								Attributes: map[string]schema.Attribute{
									"template": schema.SingleNestedAttribute{
										Description:         "Template describes the EKSLogForwarder Deployment pod that will be created.",
										MarkdownDescription: "Template describes the EKSLogForwarder Deployment pod that will be created.",
										Attributes: map[string]schema.Attribute{
											"spec": schema.SingleNestedAttribute{
												Description:         "Spec is the EKSLogForwarder Deployment's PodSpec.",
												MarkdownDescription: "Spec is the EKSLogForwarder Deployment's PodSpec.",
												Attributes: map[string]schema.Attribute{
													"containers": schema.ListNestedAttribute{
														Description:         "Containers is a list of EKSLogForwarder containers. If specified, this overrides the specified EKSLogForwarder Deployment containers. If omitted, the EKSLogForwarder Deployment will use its default values for its containers.",
														MarkdownDescription: "Containers is a list of EKSLogForwarder containers. If specified, this overrides the specified EKSLogForwarder Deployment containers. If omitted, the EKSLogForwarder Deployment will use its default values for its containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the EKSLogForwarder Deployment container by name.",
																	MarkdownDescription: "Name is an enum which identifies the EKSLogForwarder Deployment container by name.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("eks-log-forwarder"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named EKSLogForwarder Deployment container's resources. If omitted, the EKSLogForwarder Deployment will use its default value for this container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named EKSLogForwarder Deployment container's resources. If omitted, the EKSLogForwarder Deployment will use its default value for this container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"limits": schema.MapAttribute{
																			Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"init_containers": schema.ListNestedAttribute{
														Description:         "InitContainers is a list of EKSLogForwarder init containers. If specified, this overrides the specified EKSLogForwarder Deployment init containers. If omitted, the EKSLogForwarder Deployment will use its default values for its init containers.",
														MarkdownDescription: "InitContainers is a list of EKSLogForwarder init containers. If specified, this overrides the specified EKSLogForwarder Deployment init containers. If omitted, the EKSLogForwarder Deployment will use its default values for its init containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the EKSLogForwarder Deployment init container by name.",
																	MarkdownDescription: "Name is an enum which identifies the EKSLogForwarder Deployment init container by name.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("eks-log-forwarder-startup"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named EKSLogForwarder Deployment init container's resources. If omitted, the EKSLogForwarder Deployment will use its default value for this init container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named EKSLogForwarder Deployment init container's resources. If omitted, the EKSLogForwarder Deployment will use its default value for this init container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"limits": schema.MapAttribute{
																			Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"fluentd_daemon_set": schema.SingleNestedAttribute{
						Description:         "FluentdDaemonSet configures the Fluentd DaemonSet.",
						MarkdownDescription: "FluentdDaemonSet configures the Fluentd DaemonSet.",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "Spec is the specification of the Fluentd DaemonSet.",
								MarkdownDescription: "Spec is the specification of the Fluentd DaemonSet.",
								Attributes: map[string]schema.Attribute{
									"template": schema.SingleNestedAttribute{
										Description:         "Template describes the Fluentd DaemonSet pod that will be created.",
										MarkdownDescription: "Template describes the Fluentd DaemonSet pod that will be created.",
										Attributes: map[string]schema.Attribute{
											"spec": schema.SingleNestedAttribute{
												Description:         "Spec is the Fluentd DaemonSet's PodSpec.",
												MarkdownDescription: "Spec is the Fluentd DaemonSet's PodSpec.",
												Attributes: map[string]schema.Attribute{
													"containers": schema.ListNestedAttribute{
														Description:         "Containers is a list of Fluentd DaemonSet containers. If specified, this overrides the specified Fluentd DaemonSet containers. If omitted, the Fluentd DaemonSet will use its default values for its containers.",
														MarkdownDescription: "Containers is a list of Fluentd DaemonSet containers. If specified, this overrides the specified Fluentd DaemonSet containers. If omitted, the Fluentd DaemonSet will use its default values for its containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the Fluentd DaemonSet container by name.",
																	MarkdownDescription: "Name is an enum which identifies the Fluentd DaemonSet container by name.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("fluentd"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named Fluentd DaemonSet container's resources. If omitted, the Fluentd DaemonSet will use its default value for this container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named Fluentd DaemonSet container's resources. If omitted, the Fluentd DaemonSet will use its default value for this container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"limits": schema.MapAttribute{
																			Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"init_containers": schema.ListNestedAttribute{
														Description:         "InitContainers is a list of Fluentd DaemonSet init containers. If specified, this overrides the specified Fluentd DaemonSet init containers. If omitted, the Fluentd DaemonSet will use its default values for its init containers.",
														MarkdownDescription: "InitContainers is a list of Fluentd DaemonSet init containers. If specified, this overrides the specified Fluentd DaemonSet init containers. If omitted, the Fluentd DaemonSet will use its default values for its init containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the Fluentd DaemonSet init container by name.",
																	MarkdownDescription: "Name is an enum which identifies the Fluentd DaemonSet init container by name.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("tigera-fluentd-prometheus-tls-key-cert-provisioner"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named Fluentd DaemonSet init container's resources. If omitted, the Fluentd DaemonSet will use its default value for this init container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named Fluentd DaemonSet init container's resources. If omitted, the Fluentd DaemonSet will use its default value for this init container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"limits": schema.MapAttribute{
																			Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"multi_tenant_management_cluster_namespace": schema.StringAttribute{
						Description:         "If running as a multi-tenant management cluster, the namespace in which the management cluster's tenant services are running.",
						MarkdownDescription: "If running as a multi-tenant management cluster, the namespace in which the management cluster's tenant services are running.",
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

func (r *OperatorTigeraIoLogCollectorV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_tigera_io_log_collector_v1_manifest")

	var model OperatorTigeraIoLogCollectorV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.tigera.io/v1")
	model.Kind = pointer.String("LogCollector")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
