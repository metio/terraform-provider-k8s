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
	_ datasource.DataSource = &OperatorTigeraIoTenantV1Manifest{}
)

func NewOperatorTigeraIoTenantV1Manifest() datasource.DataSource {
	return &OperatorTigeraIoTenantV1Manifest{}
}

type OperatorTigeraIoTenantV1Manifest struct{}

type OperatorTigeraIoTenantV1ManifestData struct {
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
		ControlPlaneReplicas *int64 `tfsdk:"control_plane_replicas" json:"controlPlaneReplicas,omitempty"`
		DashboardsJob        *struct {
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
					} `tfsdk:"spec" json:"spec,omitempty"`
				} `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"dashboards_job" json:"dashboardsJob,omitempty"`
		Elastic *struct {
			KibanaURL *string `tfsdk:"kibana_url" json:"kibanaURL,omitempty"`
			MutualTLS *bool   `tfsdk:"mutual_tls" json:"mutualTLS,omitempty"`
			Url       *string `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"elastic" json:"elastic,omitempty"`
		Id      *string `tfsdk:"id" json:"id,omitempty"`
		Indices *[]struct {
			BaseIndexName *string `tfsdk:"base_index_name" json:"baseIndexName,omitempty"`
			DataType      *string `tfsdk:"data_type" json:"dataType,omitempty"`
		} `tfsdk:"indices" json:"indices,omitempty"`
		LinseedDeployment *struct {
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
		} `tfsdk:"linseed_deployment" json:"linseedDeployment,omitempty"`
		Name *string `tfsdk:"name" json:"name,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorTigeraIoTenantV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_tigera_io_tenant_v1_manifest"
}

func (r *OperatorTigeraIoTenantV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Tenant is the Schema for the tenants API",
		MarkdownDescription: "Tenant is the Schema for the tenants API",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"control_plane_replicas": schema.Int64Attribute{
						Description:         "ControlPlaneReplicas defines how many replicas of the control plane core components will be deployed in the Tenant's namespace. Defaults to the controlPlaneReplicas in Installation CR",
						MarkdownDescription: "ControlPlaneReplicas defines how many replicas of the control plane core components will be deployed in the Tenant's namespace. Defaults to the controlPlaneReplicas in Installation CR",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"dashboards_job": schema.SingleNestedAttribute{
						Description:         "DashboardsJob configures the Dashboards job",
						MarkdownDescription: "DashboardsJob configures the Dashboards job",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "Spec is the specification of the dashboards job.",
								MarkdownDescription: "Spec is the specification of the dashboards job.",
								Attributes: map[string]schema.Attribute{
									"template": schema.SingleNestedAttribute{
										Description:         "Template describes the Dashboards job pod that will be created.",
										MarkdownDescription: "Template describes the Dashboards job pod that will be created.",
										Attributes: map[string]schema.Attribute{
											"spec": schema.SingleNestedAttribute{
												Description:         "Spec is the Dashboard job's PodSpec.",
												MarkdownDescription: "Spec is the Dashboard job's PodSpec.",
												Attributes: map[string]schema.Attribute{
													"containers": schema.ListNestedAttribute{
														Description:         "Containers is a list of dashboards job containers. If specified, this overrides the specified Dashboard job containers. If omitted, the Dashboard job will use its default values for its containers.",
														MarkdownDescription: "Containers is a list of dashboards job containers. If specified, this overrides the specified Dashboard job containers. If omitted, the Dashboard job will use its default values for its containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the Dashboard Job container by name.",
																	MarkdownDescription: "Name is an enum which identifies the Dashboard Job container by name.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("dashboards-installer"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named Dashboard Job container's resources. If omitted, the Dashboard Job will use its default value for this container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named Dashboard Job container's resources. If omitted, the Dashboard Job will use its default value for this container's resources.",
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

					"elastic": schema.SingleNestedAttribute{
						Description:         "Elastic configures per-tenant ElasticSearch and Kibana parameters. This field is required for clusters using external ES.",
						MarkdownDescription: "Elastic configures per-tenant ElasticSearch and Kibana parameters. This field is required for clusters using external ES.",
						Attributes: map[string]schema.Attribute{
							"kibana_url": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mutual_tls": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"url": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"id": schema.StringAttribute{
						Description:         "ID is the unique identifier for this tenant.",
						MarkdownDescription: "ID is the unique identifier for this tenant.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"indices": schema.ListNestedAttribute{
						Description:         "Indices defines the how to store a tenant's data",
						MarkdownDescription: "Indices defines the how to store a tenant's data",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"base_index_name": schema.StringAttribute{
									Description:         "BaseIndexName defines the name of the index that will be used to store data (this name excludes the numerical identifier suffix)",
									MarkdownDescription: "BaseIndexName defines the name of the index that will be used to store data (this name excludes the numerical identifier suffix)",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"data_type": schema.StringAttribute{
									Description:         "DataType represents the type of data stored in the defined index",
									MarkdownDescription: "DataType represents the type of data stored in the defined index",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Alerts", "AuditLogs", "BGPLogs", "ComplianceBenchmarks", "ComplianceReports", "ComplianceSnapshots", "DNSLogs", "FlowLogs", "L7Logs", "RuntimeReports", "ThreatFeedsDomainSet", "ThreatFeedsIPSet", "WAFLogs"),
									},
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"linseed_deployment": schema.SingleNestedAttribute{
						Description:         "LinseedDeployment configures the linseed Deployment.",
						MarkdownDescription: "LinseedDeployment configures the linseed Deployment.",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "Spec is the specification of the linseed Deployment.",
								MarkdownDescription: "Spec is the specification of the linseed Deployment.",
								Attributes: map[string]schema.Attribute{
									"template": schema.SingleNestedAttribute{
										Description:         "Template describes the linseed Deployment pod that will be created.",
										MarkdownDescription: "Template describes the linseed Deployment pod that will be created.",
										Attributes: map[string]schema.Attribute{
											"spec": schema.SingleNestedAttribute{
												Description:         "Spec is the linseed Deployment's PodSpec.",
												MarkdownDescription: "Spec is the linseed Deployment's PodSpec.",
												Attributes: map[string]schema.Attribute{
													"containers": schema.ListNestedAttribute{
														Description:         "Containers is a list of linseed containers. If specified, this overrides the specified linseed Deployment containers. If omitted, the linseed Deployment will use its default values for its containers.",
														MarkdownDescription: "Containers is a list of linseed containers. If specified, this overrides the specified linseed Deployment containers. If omitted, the linseed Deployment will use its default values for its containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the linseed Deployment container by name.",
																	MarkdownDescription: "Name is an enum which identifies the linseed Deployment container by name.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("tigera-linseed"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named linseed Deployment container's resources. If omitted, the linseed Deployment will use its default value for this container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named linseed Deployment container's resources. If omitted, the linseed Deployment will use its default value for this container's resources.",
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
														Description:         "InitContainers is a list of linseed init containers. If specified, this overrides the specified linseed Deployment init containers. If omitted, the linseed Deployment will use its default values for its init containers.",
														MarkdownDescription: "InitContainers is a list of linseed init containers. If specified, this overrides the specified linseed Deployment init containers. If omitted, the linseed Deployment will use its default values for its init containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the linseed Deployment init container by name.",
																	MarkdownDescription: "Name is an enum which identifies the linseed Deployment init container by name.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("tigera-secure-linseed-token-tls-key-cert-provisioner", "tigera-secure-linseed-cert-key-cert-provisioner"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named linseed Deployment init container's resources. If omitted, the linseed Deployment will use its default value for this init container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory. If specified, this overrides the named linseed Deployment init container's resources. If omitted, the linseed Deployment will use its default value for this init container's resources.",
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

					"name": schema.StringAttribute{
						Description:         "Name is a human readable name for this tenant.",
						MarkdownDescription: "Name is a human readable name for this tenant.",
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

func (r *OperatorTigeraIoTenantV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_tigera_io_tenant_v1_manifest")

	var model OperatorTigeraIoTenantV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.tigera.io/v1")
	model.Kind = pointer.String("Tenant")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
