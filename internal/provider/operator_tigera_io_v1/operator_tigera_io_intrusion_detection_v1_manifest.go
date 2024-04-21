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
	_ datasource.DataSource = &OperatorTigeraIoIntrusionDetectionV1Manifest{}
)

func NewOperatorTigeraIoIntrusionDetectionV1Manifest() datasource.DataSource {
	return &OperatorTigeraIoIntrusionDetectionV1Manifest{}
}

type OperatorTigeraIoIntrusionDetectionV1Manifest struct{}

type OperatorTigeraIoIntrusionDetectionV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AnomalyDetection *struct {
			StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
		} `tfsdk:"anomaly_detection" json:"anomalyDetection,omitempty"`
		ComponentResources *[]struct {
			ComponentName        *string `tfsdk:"component_name" json:"componentName,omitempty"`
			ResourceRequirements *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resource_requirements" json:"resourceRequirements,omitempty"`
		} `tfsdk:"component_resources" json:"componentResources,omitempty"`
		IntrusionDetectionControllerDeployment *struct {
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
		} `tfsdk:"intrusion_detection_controller_deployment" json:"intrusionDetectionControllerDeployment,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorTigeraIoIntrusionDetectionV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_tigera_io_intrusion_detection_v1_manifest"
}

func (r *OperatorTigeraIoIntrusionDetectionV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "IntrusionDetection installs the components required for Tigera intrusion detection. At most one instanceof this resource is supported. It must be named 'tigera-secure'.",
		MarkdownDescription: "IntrusionDetection installs the components required for Tigera intrusion detection. At most one instanceof this resource is supported. It must be named 'tigera-secure'.",
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
				Description:         "Specification of the desired state for Tigera intrusion detection.",
				MarkdownDescription: "Specification of the desired state for Tigera intrusion detection.",
				Attributes: map[string]schema.Attribute{
					"anomaly_detection": schema.SingleNestedAttribute{
						Description:         "AnomalyDetection is now deprecated, and configuring it has no effect.",
						MarkdownDescription: "AnomalyDetection is now deprecated, and configuring it has no effect.",
						Attributes: map[string]schema.Attribute{
							"storage_class_name": schema.StringAttribute{
								Description:         "StorageClassName is now deprecated, and configuring it has no effect.",
								MarkdownDescription: "StorageClassName is now deprecated, and configuring it has no effect.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"component_resources": schema.ListNestedAttribute{
						Description:         "ComponentResources can be used to customize the resource requirements for each component.Only DeepPacketInspection is supported for this spec.",
						MarkdownDescription: "ComponentResources can be used to customize the resource requirements for each component.Only DeepPacketInspection is supported for this spec.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "ComponentName is an enum which identifies the component",
									MarkdownDescription: "ComponentName is an enum which identifies the component",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("DeepPacketInspection"),
									},
								},

								"resource_requirements": schema.SingleNestedAttribute{
									Description:         "ResourceRequirements allows customization of limits and requests for compute resources such as cpu and memory.",
									MarkdownDescription: "ResourceRequirements allows customization of limits and requests for compute resources such as cpu and memory.",
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
									Required: true,
									Optional: false,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"intrusion_detection_controller_deployment": schema.SingleNestedAttribute{
						Description:         "IntrusionDetectionControllerDeployment configures the IntrusionDetection Controller Deployment.",
						MarkdownDescription: "IntrusionDetectionControllerDeployment configures the IntrusionDetection Controller Deployment.",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "Spec is the specification of the IntrusionDetectionController Deployment.",
								MarkdownDescription: "Spec is the specification of the IntrusionDetectionController Deployment.",
								Attributes: map[string]schema.Attribute{
									"template": schema.SingleNestedAttribute{
										Description:         "Template describes the IntrusionDetectionController Deployment pod that will be created.",
										MarkdownDescription: "Template describes the IntrusionDetectionController Deployment pod that will be created.",
										Attributes: map[string]schema.Attribute{
											"spec": schema.SingleNestedAttribute{
												Description:         "Spec is the IntrusionDetectionController Deployment's PodSpec.",
												MarkdownDescription: "Spec is the IntrusionDetectionController Deployment's PodSpec.",
												Attributes: map[string]schema.Attribute{
													"containers": schema.ListNestedAttribute{
														Description:         "Containers is a list of IntrusionDetectionController containers.If specified, this overrides the specified IntrusionDetectionController Deployment containers.If omitted, the IntrusionDetectionController Deployment will use its default values for its containers.",
														MarkdownDescription: "Containers is a list of IntrusionDetectionController containers.If specified, this overrides the specified IntrusionDetectionController Deployment containers.If omitted, the IntrusionDetectionController Deployment will use its default values for its containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the IntrusionDetectionController Deployment container by name.Supported values are: controller, webhooks-processor",
																	MarkdownDescription: "Name is an enum which identifies the IntrusionDetectionController Deployment container by name.Supported values are: controller, webhooks-processor",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("controller", "webhooks-processor"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named IntrusionDetectionController Deployment container's resources.If omitted, the IntrusionDetection Deployment will use its default value for this container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named IntrusionDetectionController Deployment container's resources.If omitted, the IntrusionDetection Deployment will use its default value for this container's resources.",
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
														Description:         "InitContainers is a list of IntrusionDetectionController init containers.If specified, this overrides the specified IntrusionDetectionController Deployment init containers.If omitted, the IntrusionDetectionController Deployment will use its default values for its init containers.",
														MarkdownDescription: "InitContainers is a list of IntrusionDetectionController init containers.If specified, this overrides the specified IntrusionDetectionController Deployment init containers.If omitted, the IntrusionDetectionController Deployment will use its default values for its init containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is an enum which identifies the IntrusionDetectionController Deployment init container by name.Supported values are: intrusion-detection-tls-key-cert-provisioner",
																	MarkdownDescription: "Name is an enum which identifies the IntrusionDetectionController Deployment init container by name.Supported values are: intrusion-detection-tls-key-cert-provisioner",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("intrusion-detection-tls-key-cert-provisioner"),
																	},
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named IntrusionDetectionController Deployment init container's resources.If omitted, the IntrusionDetectionController Deployment will use its default value for this init container's resources.",
																	MarkdownDescription: "Resources allows customization of limits and requests for compute resources such as cpu and memory.If specified, this overrides the named IntrusionDetectionController Deployment init container's resources.If omitted, the IntrusionDetectionController Deployment will use its default value for this init container's resources.",
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

func (r *OperatorTigeraIoIntrusionDetectionV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_tigera_io_intrusion_detection_v1_manifest")

	var model OperatorTigeraIoIntrusionDetectionV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.tigera.io/v1")
	model.Kind = pointer.String("IntrusionDetection")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
