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
	_ datasource.DataSource = &OperatorTigeraIoComplianceV1Manifest{}
)

func NewOperatorTigeraIoComplianceV1Manifest() datasource.DataSource {
	return &OperatorTigeraIoComplianceV1Manifest{}
}

type OperatorTigeraIoComplianceV1Manifest struct{}

type OperatorTigeraIoComplianceV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ComplianceBenchmarkerDaemonSet *struct {
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
		} `tfsdk:"compliance_benchmarker_daemon_set" json:"complianceBenchmarkerDaemonSet,omitempty"`
		ComplianceControllerDeployment *struct {
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
		} `tfsdk:"compliance_controller_deployment" json:"complianceControllerDeployment,omitempty"`
		ComplianceReporterPodTemplate *struct {
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
		} `tfsdk:"compliance_reporter_pod_template" json:"complianceReporterPodTemplate,omitempty"`
		ComplianceServerDeployment *struct {
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
		} `tfsdk:"compliance_server_deployment" json:"complianceServerDeployment,omitempty"`
		ComplianceSnapshotterDeployment *struct {
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
		} `tfsdk:"compliance_snapshotter_deployment" json:"complianceSnapshotterDeployment,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorTigeraIoComplianceV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_tigera_io_compliance_v1_manifest"
}

func (r *OperatorTigeraIoComplianceV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Compliance installs the components required for Tigera compliance reporting. At most one instanceof this resource is supported. It must be named 'tigera-secure'.",
		MarkdownDescription: "Compliance installs the components required for Tigera compliance reporting. At most one instanceof this resource is supported. It must be named 'tigera-secure'.",
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
				Description:         "Specification of the desired state for Tigera compliance reporting.",
				MarkdownDescription: "Specification of the desired state for Tigera compliance reporting.",
				Attributes: map[string]schema.Attribute{
					"compliance_benchmarker_daemon_set": schema.SingleNestedAttribute{
						Description:         "ComplianceBenchmarkerDaemonSet configures the Compliance Benchmarker DaemonSet.",
						MarkdownDescription: "ComplianceBenchmarkerDaemonSet configures the Compliance Benchmarker DaemonSet.",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "Spec is the specification of the Compliance Benchmarker DaemonSet.",
								MarkdownDescription: "Spec is the specification of the Compliance Benchmarker DaemonSet.",
								Attributes: map[string]schema.Attribute{
									"template": schema.SingleNestedAttribute{
										Description:         "Template describes the Compliance Benchmarker DaemonSet pod that will be created.",
										MarkdownDescription: "Template describes the Compliance Benchmarker DaemonSet pod that will be created.",
										Attributes: map[string]schema.Attribute{
											"spec": schema.SingleNestedAttribute{
												Description:         "Spec is the Compliance Benchmarker DaemonSet's PodSpec.",
												MarkdownDescription: "Spec is the Compliance Benchmarker DaemonSet's PodSpec.",
												Attributes: map[string]schema.Attribute{
													"containers": schema.ListNestedAttribute{
														Description:         "Containers is a list of Compliance benchmark containers.If specified, this overrides the specified Compliance Benchmarker DaemonSet containers.If omitted, the Compliance Benchmarker DaemonSet will use its default values for its containers.",
														MarkdownDescription: "Containers is a list of Compliance benchmark containers.If specified, this overrides the specified Compliance Benchmarker DaemonSet containers.If omitted, the Compliance Benchmarker DaemonSet will use its default values for its containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the Compliance Benchmarker DaemonSet container by name.Supported values are: compliance-benchmarker",
																	MarkdownDescription: "Name is an enum which identifies the Compliance Benchmarker DaemonSet container by name.Supported values are: compliance-benchmarker",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("compliance-benchmarker"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named Compliance Benchmarker DaemonSet container's resources.If omitted, the Compliance Benchmarker DaemonSet will use its default value for this container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named Compliance Benchmarker DaemonSet container's resources.If omitted, the Compliance Benchmarker DaemonSet will use its default value for this container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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
																			Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
														Description:         "InitContainers is a list of Compliance benchmark init containers.If specified, this overrides the specified Compliance Benchmarker DaemonSet init containers.If omitted, the Compliance Benchmarker DaemonSet will use its default values for its init containers.",
														MarkdownDescription: "InitContainers is a list of Compliance benchmark init containers.If specified, this overrides the specified Compliance Benchmarker DaemonSet init containers.If omitted, the Compliance Benchmarker DaemonSet will use its default values for its init containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the Compliance Benchmarker DaemonSet init container by name.Supported values are: tigera-compliance-benchmarker-tls-key-cert-provisioner",
																	MarkdownDescription: "Name is an enum which identifies the Compliance Benchmarker DaemonSet init container by name.Supported values are: tigera-compliance-benchmarker-tls-key-cert-provisioner",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("tigera-compliance-benchmarker-tls-key-cert-provisioner"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named Compliance Benchmarker DaemonSet init container's resources.If omitted, the Compliance Benchmarker DaemonSet will use its default value for this init container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named Compliance Benchmarker DaemonSet init container's resources.If omitted, the Compliance Benchmarker DaemonSet will use its default value for this init container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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
																			Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"compliance_controller_deployment": schema.SingleNestedAttribute{
						Description:         "ComplianceControllerDeployment configures the Compliance Controller Deployment.",
						MarkdownDescription: "ComplianceControllerDeployment configures the Compliance Controller Deployment.",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "Spec is the specification of the compliance controller Deployment.",
								MarkdownDescription: "Spec is the specification of the compliance controller Deployment.",
								Attributes: map[string]schema.Attribute{
									"template": schema.SingleNestedAttribute{
										Description:         "Template describes the compliance controller Deployment pod that will be created.",
										MarkdownDescription: "Template describes the compliance controller Deployment pod that will be created.",
										Attributes: map[string]schema.Attribute{
											"spec": schema.SingleNestedAttribute{
												Description:         "Spec is the compliance controller Deployment's PodSpec.",
												MarkdownDescription: "Spec is the compliance controller Deployment's PodSpec.",
												Attributes: map[string]schema.Attribute{
													"containers": schema.ListNestedAttribute{
														Description:         "Containers is a list of compliance controller containers.If specified, this overrides the specified compliance controller Deployment containers.If omitted, the compliance controller Deployment will use its default values for its containers.",
														MarkdownDescription: "Containers is a list of compliance controller containers.If specified, this overrides the specified compliance controller Deployment containers.If omitted, the compliance controller Deployment will use its default values for its containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the compliance controller Deployment container by name.Supported values are: compliance-controller",
																	MarkdownDescription: "Name is an enum which identifies the compliance controller Deployment container by name.Supported values are: compliance-controller",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("compliance-controller"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named compliance controller Deployment container's resources.If omitted, the compliance controller Deployment will use its default value for this container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named compliance controller Deployment container's resources.If omitted, the compliance controller Deployment will use its default value for this container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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
																			Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
														Description:         "InitContainers is a list of compliance controller init containers.If specified, this overrides the specified compliance controller Deployment init containers.If omitted, the compliance controller Deployment will use its default values for its init containers.",
														MarkdownDescription: "InitContainers is a list of compliance controller init containers.If specified, this overrides the specified compliance controller Deployment init containers.If omitted, the compliance controller Deployment will use its default values for its init containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the compliance controller Deployment init container by name.Supported values are: tigera-compliance-controller-tls-key-cert-provisioner",
																	MarkdownDescription: "Name is an enum which identifies the compliance controller Deployment init container by name.Supported values are: tigera-compliance-controller-tls-key-cert-provisioner",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("tigera-compliance-controller-tls-key-cert-provisioner"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named compliance controller Deployment init container's resources.If omitted, the compliance controller Deployment will use its default value for this init container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named compliance controller Deployment init container's resources.If omitted, the compliance controller Deployment will use its default value for this init container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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
																			Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"compliance_reporter_pod_template": schema.SingleNestedAttribute{
						Description:         "ComplianceReporterPodTemplate configures the Compliance Reporter PodTemplate.",
						MarkdownDescription: "ComplianceReporterPodTemplate configures the Compliance Reporter PodTemplate.",
						Attributes: map[string]schema.Attribute{
							"template": schema.SingleNestedAttribute{
								Description:         "Spec is the specification of the ComplianceReporter PodTemplateSpec.",
								MarkdownDescription: "Spec is the specification of the ComplianceReporter PodTemplateSpec.",
								Attributes: map[string]schema.Attribute{
									"spec": schema.SingleNestedAttribute{
										Description:         "Spec is the ComplianceReporter PodTemplate's PodSpec.",
										MarkdownDescription: "Spec is the ComplianceReporter PodTemplate's PodSpec.",
										Attributes: map[string]schema.Attribute{
											"containers": schema.ListNestedAttribute{
												Description:         "Containers is a list of ComplianceServer containers.If specified, this overrides the specified ComplianceReporter PodSpec containers.If omitted, the ComplianceServer Deployment will use its default values for its containers.",
												MarkdownDescription: "Containers is a list of ComplianceServer containers.If specified, this overrides the specified ComplianceReporter PodSpec containers.If omitted, the ComplianceServer Deployment will use its default values for its containers.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name is an enum which identifies the ComplianceServer Deployment container by name.Supported values are: reporter",
															MarkdownDescription: "Name is an enum which identifies the ComplianceServer Deployment container by name.Supported values are: reporter",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("reporter"),
															},
														},

														"resources": schema.SingleNestedAttribute{
															Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named ComplianceServer Deployment container's resources.If omitted, the ComplianceServer Deployment will use its default value for this container's resources.",
															MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named ComplianceServer Deployment container's resources.If omitted, the ComplianceServer Deployment will use its default value for this container's resources.",
															Attributes: map[string]schema.Attribute{
																"claims": schema.ListNestedAttribute{
																	Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																	MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
																				MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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
																	Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																	MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"requests": schema.MapAttribute{
																	Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																	MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
												Description:         "InitContainers is a list of ComplianceReporter PodSpec init containers.If specified, this overrides the specified ComplianceReporter PodSpec init containers.If omitted, the ComplianceServer Deployment will use its default values for its init containers.",
												MarkdownDescription: "InitContainers is a list of ComplianceReporter PodSpec init containers.If specified, this overrides the specified ComplianceReporter PodSpec init containers.If omitted, the ComplianceServer Deployment will use its default values for its init containers.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name is an enum which identifies the ComplianceReporter PodSpec init container by name.Supported values are: tigera-compliance-reporter-tls-key-cert-provisioner",
															MarkdownDescription: "Name is an enum which identifies the ComplianceReporter PodSpec init container by name.Supported values are: tigera-compliance-reporter-tls-key-cert-provisioner",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("tigera-compliance-reporter-tls-key-cert-provisioner"),
															},
														},

														"resources": schema.SingleNestedAttribute{
															Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named ComplianceReporter PodSpec init container's resources.If omitted, the ComplianceServer Deployment will use its default value for this init container's resources.",
															MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named ComplianceReporter PodSpec init container's resources.If omitted, the ComplianceServer Deployment will use its default value for this init container's resources.",
															Attributes: map[string]schema.Attribute{
																"claims": schema.ListNestedAttribute{
																	Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																	MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
																				MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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
																	Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																	MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"requests": schema.MapAttribute{
																	Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																	MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"compliance_server_deployment": schema.SingleNestedAttribute{
						Description:         "ComplianceServerDeployment configures the Compliance Server Deployment.",
						MarkdownDescription: "ComplianceServerDeployment configures the Compliance Server Deployment.",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "Spec is the specification of the ComplianceServer Deployment.",
								MarkdownDescription: "Spec is the specification of the ComplianceServer Deployment.",
								Attributes: map[string]schema.Attribute{
									"template": schema.SingleNestedAttribute{
										Description:         "Template describes the ComplianceServer Deployment pod that will be created.",
										MarkdownDescription: "Template describes the ComplianceServer Deployment pod that will be created.",
										Attributes: map[string]schema.Attribute{
											"spec": schema.SingleNestedAttribute{
												Description:         "Spec is the ComplianceServer Deployment's PodSpec.",
												MarkdownDescription: "Spec is the ComplianceServer Deployment's PodSpec.",
												Attributes: map[string]schema.Attribute{
													"containers": schema.ListNestedAttribute{
														Description:         "Containers is a list of ComplianceServer containers.If specified, this overrides the specified ComplianceServer Deployment containers.If omitted, the ComplianceServer Deployment will use its default values for its containers.",
														MarkdownDescription: "Containers is a list of ComplianceServer containers.If specified, this overrides the specified ComplianceServer Deployment containers.If omitted, the ComplianceServer Deployment will use its default values for its containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the ComplianceServer Deployment container by name.Supported values are: compliance-server",
																	MarkdownDescription: "Name is an enum which identifies the ComplianceServer Deployment container by name.Supported values are: compliance-server",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("compliance-server"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named ComplianceServer Deployment container's resources.If omitted, the ComplianceServer Deployment will use its default value for this container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named ComplianceServer Deployment container's resources.If omitted, the ComplianceServer Deployment will use its default value for this container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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
																			Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
														Description:         "InitContainers is a list of ComplianceServer init containers.If specified, this overrides the specified ComplianceServer Deployment init containers.If omitted, the ComplianceServer Deployment will use its default values for its init containers.",
														MarkdownDescription: "InitContainers is a list of ComplianceServer init containers.If specified, this overrides the specified ComplianceServer Deployment init containers.If omitted, the ComplianceServer Deployment will use its default values for its init containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the ComplianceServer Deployment init container by name.Supported values are: tigera-compliance-server-tls-key-cert-provisioner",
																	MarkdownDescription: "Name is an enum which identifies the ComplianceServer Deployment init container by name.Supported values are: tigera-compliance-server-tls-key-cert-provisioner",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("tigera-compliance-server-tls-key-cert-provisioner"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named ComplianceServer Deployment init container's resources.If omitted, the ComplianceServer Deployment will use its default value for this init container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named ComplianceServer Deployment init container's resources.If omitted, the ComplianceServer Deployment will use its default value for this init container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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
																			Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"compliance_snapshotter_deployment": schema.SingleNestedAttribute{
						Description:         "ComplianceSnapshotterDeployment configures the Compliance Snapshotter Deployment.",
						MarkdownDescription: "ComplianceSnapshotterDeployment configures the Compliance Snapshotter Deployment.",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "Spec is the specification of the compliance snapshotter Deployment.",
								MarkdownDescription: "Spec is the specification of the compliance snapshotter Deployment.",
								Attributes: map[string]schema.Attribute{
									"template": schema.SingleNestedAttribute{
										Description:         "Template describes the compliance snapshotter Deployment pod that will be created.",
										MarkdownDescription: "Template describes the compliance snapshotter Deployment pod that will be created.",
										Attributes: map[string]schema.Attribute{
											"spec": schema.SingleNestedAttribute{
												Description:         "Spec is the compliance snapshotter Deployment's PodSpec.",
												MarkdownDescription: "Spec is the compliance snapshotter Deployment's PodSpec.",
												Attributes: map[string]schema.Attribute{
													"containers": schema.ListNestedAttribute{
														Description:         "Containers is a list of compliance snapshotter containers.If specified, this overrides the specified compliance snapshotter Deployment containers.If omitted, the compliance snapshotter Deployment will use its default values for its containers.",
														MarkdownDescription: "Containers is a list of compliance snapshotter containers.If specified, this overrides the specified compliance snapshotter Deployment containers.If omitted, the compliance snapshotter Deployment will use its default values for its containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the compliance snapshotter Deployment container by name.Supported values are: compliance-snapshotter",
																	MarkdownDescription: "Name is an enum which identifies the compliance snapshotter Deployment container by name.Supported values are: compliance-snapshotter",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("compliance-snapshotter"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named compliance snapshotter Deployment container's resources.If omitted, the compliance snapshotter Deployment will use its default value for this container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named compliance snapshotter Deployment container's resources.If omitted, the compliance snapshotter Deployment will use its default value for this container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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
																			Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
														Description:         "InitContainers is a list of compliance snapshotter init containers.If specified, this overrides the specified compliance snapshotter Deployment init containers.If omitted, the compliance snapshotter Deployment will use its default values for its init containers.",
														MarkdownDescription: "InitContainers is a list of compliance snapshotter init containers.If specified, this overrides the specified compliance snapshotter Deployment init containers.If omitted, the compliance snapshotter Deployment will use its default values for its init containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the compliance snapshotter Deployment init container by name.Supported values are: tigera-compliance-snapshotter-tls-key-cert-provisioner",
																	MarkdownDescription: "Name is an enum which identifies the compliance snapshotter Deployment init container by name.Supported values are: tigera-compliance-snapshotter-tls-key-cert-provisioner",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("tigera-compliance-snapshotter-tls-key-cert-provisioner"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named compliance snapshotter Deployment init container's resources.If omitted, the compliance snapshotter Deployment will use its default value for this init container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named compliance snapshotter Deployment init container's resources.If omitted, the compliance snapshotter Deployment will use its default value for this init container's resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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
																			Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *OperatorTigeraIoComplianceV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_tigera_io_compliance_v1_manifest")

	var model OperatorTigeraIoComplianceV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.tigera.io/v1")
	model.Kind = pointer.String("Compliance")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
