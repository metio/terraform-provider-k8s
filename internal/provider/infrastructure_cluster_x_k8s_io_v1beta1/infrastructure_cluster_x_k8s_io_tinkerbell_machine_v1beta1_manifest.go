/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1beta1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &InfrastructureClusterXK8SIoTinkerbellMachineV1Beta1Manifest{}
)

func NewInfrastructureClusterXK8SIoTinkerbellMachineV1Beta1Manifest() datasource.DataSource {
	return &InfrastructureClusterXK8SIoTinkerbellMachineV1Beta1Manifest{}
}

type InfrastructureClusterXK8SIoTinkerbellMachineV1Beta1Manifest struct{}

type InfrastructureClusterXK8SIoTinkerbellMachineV1Beta1ManifestData struct {
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
		HardwareAffinity *struct {
			Preferred *[]struct {
				HardwareAffinityTerm *struct {
					LabelSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
				} `tfsdk:"hardware_affinity_term" json:"hardwareAffinityTerm,omitempty"`
				Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
			} `tfsdk:"preferred" json:"preferred,omitempty"`
			Required *[]struct {
				LabelSelector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
			} `tfsdk:"required" json:"required,omitempty"`
		} `tfsdk:"hardware_affinity" json:"hardwareAffinity,omitempty"`
		HardwareName            *string `tfsdk:"hardware_name" json:"hardwareName,omitempty"`
		ImageLookupBaseRegistry *string `tfsdk:"image_lookup_base_registry" json:"imageLookupBaseRegistry,omitempty"`
		ImageLookupFormat       *string `tfsdk:"image_lookup_format" json:"imageLookupFormat,omitempty"`
		ImageLookupOSDistro     *string `tfsdk:"image_lookup_os_distro" json:"imageLookupOSDistro,omitempty"`
		ImageLookupOSVersion    *string `tfsdk:"image_lookup_os_version" json:"imageLookupOSVersion,omitempty"`
		ProviderID              *string `tfsdk:"provider_id" json:"providerID,omitempty"`
		TemplateOverride        *string `tfsdk:"template_override" json:"templateOverride,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoTinkerbellMachineV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_tinkerbell_machine_v1beta1_manifest"
}

func (r *InfrastructureClusterXK8SIoTinkerbellMachineV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TinkerbellMachine is the Schema for the tinkerbellmachines API.",
		MarkdownDescription: "TinkerbellMachine is the Schema for the tinkerbellmachines API.",
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
				Description:         "TinkerbellMachineSpec defines the desired state of TinkerbellMachine.",
				MarkdownDescription: "TinkerbellMachineSpec defines the desired state of TinkerbellMachine.",
				Attributes: map[string]schema.Attribute{
					"hardware_affinity": schema.SingleNestedAttribute{
						Description:         "HardwareAffinity allows filtering for hardware.",
						MarkdownDescription: "HardwareAffinity allows filtering for hardware.",
						Attributes: map[string]schema.Attribute{
							"preferred": schema.ListNestedAttribute{
								Description:         "Preferred are the preferred hardware affinity terms. Hardware matching these terms are preferred according to theweights provided, but are not required.",
								MarkdownDescription: "Preferred are the preferred hardware affinity terms. Hardware matching these terms are preferred according to theweights provided, but are not required.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"hardware_affinity_term": schema.SingleNestedAttribute{
											Description:         "HardwareAffinityTerm is the term associated with the corresponding weight.",
											MarkdownDescription: "HardwareAffinityTerm is the term associated with the corresponding weight.",
											Attributes: map[string]schema.Attribute{
												"label_selector": schema.SingleNestedAttribute{
													Description:         "LabelSelector is used to select for particular hardware by label.",
													MarkdownDescription: "LabelSelector is used to select for particular hardware by label.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																		ElementType:         types.StringType,
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

														"match_labels": schema.MapAttribute{
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
											Required: true,
											Optional: false,
											Computed: false,
										},

										"weight": schema.Int64Attribute{
											Description:         "Weight associated with matching the corresponding hardwareAffinityTerm, in the range 1-100.",
											MarkdownDescription: "Weight associated with matching the corresponding hardwareAffinityTerm, in the range 1-100.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
												int64validator.AtMost(100),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"required": schema.ListNestedAttribute{
								Description:         "Required are the required hardware affinity terms.  The terms are OR'd together, hardware must match one term tobe considered.",
								MarkdownDescription: "Required are the required hardware affinity terms.  The terms are OR'd together, hardware must match one term tobe considered.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"label_selector": schema.SingleNestedAttribute{
											Description:         "LabelSelector is used to select for particular hardware by label.",
											MarkdownDescription: "LabelSelector is used to select for particular hardware by label.",
											Attributes: map[string]schema.Attribute{
												"match_expressions": schema.ListNestedAttribute{
													Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"values": schema.ListAttribute{
																Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																ElementType:         types.StringType,
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

												"match_labels": schema.MapAttribute{
													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"hardware_name": schema.StringAttribute{
						Description:         "Those fields are set programmatically, but they cannot be re-constructed from 'state of the world', sowe put them in spec instead of status.",
						MarkdownDescription: "Those fields are set programmatically, but they cannot be re-constructed from 'state of the world', sowe put them in spec instead of status.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_lookup_base_registry": schema.StringAttribute{
						Description:         "ImageLookupBaseRegistry is the base Registry URL that is used for pulling images,if not set, the default will be to use ghcr.io/tinkerbell/cluster-api-provider-tinkerbell.",
						MarkdownDescription: "ImageLookupBaseRegistry is the base Registry URL that is used for pulling images,if not set, the default will be to use ghcr.io/tinkerbell/cluster-api-provider-tinkerbell.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_lookup_format": schema.StringAttribute{
						Description:         "ImageLookupFormat is the URL naming format to use for machine images whena machine does not specify. When set, this will be used for all cluster machinesunless a machine specifies a different ImageLookupFormat. Supports substitutionsfor {{.BaseRegistry}}, {{.OSDistro}}, {{.OSVersion}} and {{.KubernetesVersion}} withthe basse URL, OS distribution, OS version, and kubernetes version, respectively.BaseRegistry will be the value in ImageLookupBaseRegistry or ghcr.io/tinkerbell/cluster-api-provider-tinkerbell(the default), OSDistro will be the value in ImageLookupOSDistro or ubuntu (the default),OSVersion will be the value in ImageLookupOSVersion or default based on the OSDistro(if known), and the kubernetes version as defined by the packages produced bykubernetes/release: v1.13.0, v1.12.5-mybuild.1, or v1.17.3. For example, the defaultimage format of {{.BaseRegistry}}/{{.OSDistro}}-{{.OSVersion}}:{{.KubernetesVersion}}.gz willattempt to pull the image from that location. See also: https://golang.org/pkg/text/template/",
						MarkdownDescription: "ImageLookupFormat is the URL naming format to use for machine images whena machine does not specify. When set, this will be used for all cluster machinesunless a machine specifies a different ImageLookupFormat. Supports substitutionsfor {{.BaseRegistry}}, {{.OSDistro}}, {{.OSVersion}} and {{.KubernetesVersion}} withthe basse URL, OS distribution, OS version, and kubernetes version, respectively.BaseRegistry will be the value in ImageLookupBaseRegistry or ghcr.io/tinkerbell/cluster-api-provider-tinkerbell(the default), OSDistro will be the value in ImageLookupOSDistro or ubuntu (the default),OSVersion will be the value in ImageLookupOSVersion or default based on the OSDistro(if known), and the kubernetes version as defined by the packages produced bykubernetes/release: v1.13.0, v1.12.5-mybuild.1, or v1.17.3. For example, the defaultimage format of {{.BaseRegistry}}/{{.OSDistro}}-{{.OSVersion}}:{{.KubernetesVersion}}.gz willattempt to pull the image from that location. See also: https://golang.org/pkg/text/template/",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_lookup_os_distro": schema.StringAttribute{
						Description:         "ImageLookupOSDistro is the name of the OS distro to use when fetching machine images,if not set it will default to ubuntu.",
						MarkdownDescription: "ImageLookupOSDistro is the name of the OS distro to use when fetching machine images,if not set it will default to ubuntu.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_lookup_os_version": schema.StringAttribute{
						Description:         "ImageLookupOSVersion is the version of the OS distribution to use when fetching machineimages. If not set it will default based on ImageLookupOSDistro.",
						MarkdownDescription: "ImageLookupOSVersion is the version of the OS distribution to use when fetching machineimages. If not set it will default based on ImageLookupOSDistro.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"provider_id": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"template_override": schema.StringAttribute{
						Description:         "TemplateOverride overrides the default Tinkerbell template used by CAPT.You can learn more about Tinkerbell templates here: https://docs.tinkerbell.org/templates/",
						MarkdownDescription: "TemplateOverride overrides the default Tinkerbell template used by CAPT.You can learn more about Tinkerbell templates here: https://docs.tinkerbell.org/templates/",
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

func (r *InfrastructureClusterXK8SIoTinkerbellMachineV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infrastructure_cluster_x_k8s_io_tinkerbell_machine_v1beta1_manifest")

	var model InfrastructureClusterXK8SIoTinkerbellMachineV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1beta1")
	model.Kind = pointer.String("TinkerbellMachine")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
