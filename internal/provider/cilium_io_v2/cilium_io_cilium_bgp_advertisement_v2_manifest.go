/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cilium_io_v2

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
	_ datasource.DataSource = &CiliumIoCiliumBgpadvertisementV2Manifest{}
)

func NewCiliumIoCiliumBgpadvertisementV2Manifest() datasource.DataSource {
	return &CiliumIoCiliumBgpadvertisementV2Manifest{}
}

type CiliumIoCiliumBgpadvertisementV2Manifest struct{}

type CiliumIoCiliumBgpadvertisementV2ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Advertisements *[]struct {
			AdvertisementType *string `tfsdk:"advertisement_type" json:"advertisementType,omitempty"`
			Attributes        *struct {
				Communities *struct {
					Large     *[]string `tfsdk:"large" json:"large,omitempty"`
					Standard  *[]string `tfsdk:"standard" json:"standard,omitempty"`
					WellKnown *[]string `tfsdk:"well_known" json:"wellKnown,omitempty"`
				} `tfsdk:"communities" json:"communities,omitempty"`
				LocalPreference *int64 `tfsdk:"local_preference" json:"localPreference,omitempty"`
			} `tfsdk:"attributes" json:"attributes,omitempty"`
			Selector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"selector" json:"selector,omitempty"`
			Service *struct {
				Addresses             *[]string `tfsdk:"addresses" json:"addresses,omitempty"`
				AggregationLengthIPv4 *int64    `tfsdk:"aggregation_length_i_pv4" json:"aggregationLengthIPv4,omitempty"`
				AggregationLengthIPv6 *int64    `tfsdk:"aggregation_length_i_pv6" json:"aggregationLengthIPv6,omitempty"`
			} `tfsdk:"service" json:"service,omitempty"`
		} `tfsdk:"advertisements" json:"advertisements,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CiliumIoCiliumBgpadvertisementV2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cilium_io_cilium_bgp_advertisement_v2_manifest"
}

func (r *CiliumIoCiliumBgpadvertisementV2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CiliumBGPAdvertisement is the Schema for the ciliumbgpadvertisements API",
		MarkdownDescription: "CiliumBGPAdvertisement is the Schema for the ciliumbgpadvertisements API",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"advertisements": schema.ListNestedAttribute{
						Description:         "Advertisements is a list of BGP advertisements.",
						MarkdownDescription: "Advertisements is a list of BGP advertisements.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"advertisement_type": schema.StringAttribute{
									Description:         "AdvertisementType defines type of advertisement which has to be advertised.",
									MarkdownDescription: "AdvertisementType defines type of advertisement which has to be advertised.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("PodCIDR", "CiliumPodIPPool", "Service"),
									},
								},

								"attributes": schema.SingleNestedAttribute{
									Description:         "Attributes defines additional attributes to set to the advertised routes. If not specified, no additional attributes are set.",
									MarkdownDescription: "Attributes defines additional attributes to set to the advertised routes. If not specified, no additional attributes are set.",
									Attributes: map[string]schema.Attribute{
										"communities": schema.SingleNestedAttribute{
											Description:         "Communities sets the community attributes in the route. If not specified, no community attribute is set.",
											MarkdownDescription: "Communities sets the community attributes in the route. If not specified, no community attribute is set.",
											Attributes: map[string]schema.Attribute{
												"large": schema.ListAttribute{
													Description:         "Large holds a list of the BGP Large Communities Attribute (RFC 8092) values.",
													MarkdownDescription: "Large holds a list of the BGP Large Communities Attribute (RFC 8092) values.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"standard": schema.ListAttribute{
													Description:         "Standard holds a list of 'standard' 32-bit BGP Communities Attribute (RFC 1997) values defined as numeric values.",
													MarkdownDescription: "Standard holds a list of 'standard' 32-bit BGP Communities Attribute (RFC 1997) values defined as numeric values.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"well_known": schema.ListAttribute{
													Description:         "WellKnown holds a list 'standard' 32-bit BGP Communities Attribute (RFC 1997) values defined as well-known string aliases to their numeric values.",
													MarkdownDescription: "WellKnown holds a list 'standard' 32-bit BGP Communities Attribute (RFC 1997) values defined as well-known string aliases to their numeric values.",
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

										"local_preference": schema.Int64Attribute{
											Description:         "LocalPreference sets the local preference attribute in the route. If not specified, no local preference attribute is set.",
											MarkdownDescription: "LocalPreference sets the local preference attribute in the route. If not specified, no local preference attribute is set.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"selector": schema.SingleNestedAttribute{
									Description:         "Selector is a label selector to select objects of the type specified by AdvertisementType. For the PodCIDR AdvertisementType it is not applicable. For other advertisement types, if not specified, no objects of the type specified by AdvertisementType are selected for advertisement.",
									MarkdownDescription: "Selector is a label selector to select objects of the type specified by AdvertisementType. For the PodCIDR AdvertisementType it is not applicable. For other advertisement types, if not specified, no objects of the type specified by AdvertisementType are selected for advertisement.",
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
														Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
														},
													},

													"values": schema.ListAttribute{
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
											Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

								"service": schema.SingleNestedAttribute{
									Description:         "Service defines configuration options for advertisementType service.",
									MarkdownDescription: "Service defines configuration options for advertisementType service.",
									Attributes: map[string]schema.Attribute{
										"addresses": schema.ListAttribute{
											Description:         "Addresses is a list of service address types which needs to be advertised via BGP.",
											MarkdownDescription: "Addresses is a list of service address types which needs to be advertised via BGP.",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"aggregation_length_i_pv4": schema.Int64Attribute{
											Description:         "IPv4 mask to aggregate BGP route advertisements of service",
											MarkdownDescription: "IPv4 mask to aggregate BGP route advertisements of service",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
												int64validator.AtMost(31),
											},
										},

										"aggregation_length_i_pv6": schema.Int64Attribute{
											Description:         "IPv6 mask to aggregate BGP route advertisements of service",
											MarkdownDescription: "IPv6 mask to aggregate BGP route advertisements of service",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
												int64validator.AtMost(127),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},
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
		},
	}
}

func (r *CiliumIoCiliumBgpadvertisementV2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cilium_io_cilium_bgp_advertisement_v2_manifest")

	var model CiliumIoCiliumBgpadvertisementV2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cilium.io/v2")
	model.Kind = pointer.String("CiliumBGPAdvertisement")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
