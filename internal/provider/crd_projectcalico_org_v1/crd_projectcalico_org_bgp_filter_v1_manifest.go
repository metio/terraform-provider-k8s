/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package crd_projectcalico_org_v1

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
	_ datasource.DataSource = &CrdProjectcalicoOrgBgpfilterV1Manifest{}
)

func NewCrdProjectcalicoOrgBgpfilterV1Manifest() datasource.DataSource {
	return &CrdProjectcalicoOrgBgpfilterV1Manifest{}
}

type CrdProjectcalicoOrgBgpfilterV1Manifest struct{}

type CrdProjectcalicoOrgBgpfilterV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ExportV4 *[]struct {
			Action        *string `tfsdk:"action" json:"action,omitempty"`
			Cidr          *string `tfsdk:"cidr" json:"cidr,omitempty"`
			Interface     *string `tfsdk:"interface" json:"interface,omitempty"`
			MatchOperator *string `tfsdk:"match_operator" json:"matchOperator,omitempty"`
			PrefixLength  *struct {
				Max *int64 `tfsdk:"max" json:"max,omitempty"`
				Min *int64 `tfsdk:"min" json:"min,omitempty"`
			} `tfsdk:"prefix_length" json:"prefixLength,omitempty"`
			Source *string `tfsdk:"source" json:"source,omitempty"`
		} `tfsdk:"export_v4" json:"exportV4,omitempty"`
		ExportV6 *[]struct {
			Action        *string `tfsdk:"action" json:"action,omitempty"`
			Cidr          *string `tfsdk:"cidr" json:"cidr,omitempty"`
			Interface     *string `tfsdk:"interface" json:"interface,omitempty"`
			MatchOperator *string `tfsdk:"match_operator" json:"matchOperator,omitempty"`
			PrefixLength  *struct {
				Max *int64 `tfsdk:"max" json:"max,omitempty"`
				Min *int64 `tfsdk:"min" json:"min,omitempty"`
			} `tfsdk:"prefix_length" json:"prefixLength,omitempty"`
			Source *string `tfsdk:"source" json:"source,omitempty"`
		} `tfsdk:"export_v6" json:"exportV6,omitempty"`
		ImportV4 *[]struct {
			Action        *string `tfsdk:"action" json:"action,omitempty"`
			Cidr          *string `tfsdk:"cidr" json:"cidr,omitempty"`
			Interface     *string `tfsdk:"interface" json:"interface,omitempty"`
			MatchOperator *string `tfsdk:"match_operator" json:"matchOperator,omitempty"`
			PrefixLength  *struct {
				Max *int64 `tfsdk:"max" json:"max,omitempty"`
				Min *int64 `tfsdk:"min" json:"min,omitempty"`
			} `tfsdk:"prefix_length" json:"prefixLength,omitempty"`
			Source *string `tfsdk:"source" json:"source,omitempty"`
		} `tfsdk:"import_v4" json:"importV4,omitempty"`
		ImportV6 *[]struct {
			Action        *string `tfsdk:"action" json:"action,omitempty"`
			Cidr          *string `tfsdk:"cidr" json:"cidr,omitempty"`
			Interface     *string `tfsdk:"interface" json:"interface,omitempty"`
			MatchOperator *string `tfsdk:"match_operator" json:"matchOperator,omitempty"`
			PrefixLength  *struct {
				Max *int64 `tfsdk:"max" json:"max,omitempty"`
				Min *int64 `tfsdk:"min" json:"min,omitempty"`
			} `tfsdk:"prefix_length" json:"prefixLength,omitempty"`
			Source *string `tfsdk:"source" json:"source,omitempty"`
		} `tfsdk:"import_v6" json:"importV6,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CrdProjectcalicoOrgBgpfilterV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_crd_projectcalico_org_bgp_filter_v1_manifest"
}

func (r *CrdProjectcalicoOrgBgpfilterV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "BGPFilterSpec contains the IPv4 and IPv6 filter rules of the BGP Filter.",
				MarkdownDescription: "BGPFilterSpec contains the IPv4 and IPv6 filter rules of the BGP Filter.",
				Attributes: map[string]schema.Attribute{
					"export_v4": schema.ListNestedAttribute{
						Description:         "The ordered set of IPv4 BGPFilter rules acting on exporting routes to a peer.",
						MarkdownDescription: "The ordered set of IPv4 BGPFilter rules acting on exporting routes to a peer.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"cidr": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"interface": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"match_operator": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"prefix_length": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"max": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
												int64validator.AtMost(32),
											},
										},

										"min": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
												int64validator.AtMost(32),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"source": schema.StringAttribute{
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

					"export_v6": schema.ListNestedAttribute{
						Description:         "The ordered set of IPv6 BGPFilter rules acting on exporting routes to a peer.",
						MarkdownDescription: "The ordered set of IPv6 BGPFilter rules acting on exporting routes to a peer.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"cidr": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"interface": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"match_operator": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"prefix_length": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"max": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
												int64validator.AtMost(128),
											},
										},

										"min": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
												int64validator.AtMost(128),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"source": schema.StringAttribute{
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

					"import_v4": schema.ListNestedAttribute{
						Description:         "The ordered set of IPv4 BGPFilter rules acting on importing routes from a peer.",
						MarkdownDescription: "The ordered set of IPv4 BGPFilter rules acting on importing routes from a peer.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"cidr": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"interface": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"match_operator": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"prefix_length": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"max": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
												int64validator.AtMost(32),
											},
										},

										"min": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
												int64validator.AtMost(32),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"source": schema.StringAttribute{
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

					"import_v6": schema.ListNestedAttribute{
						Description:         "The ordered set of IPv6 BGPFilter rules acting on importing routes from a peer.",
						MarkdownDescription: "The ordered set of IPv6 BGPFilter rules acting on importing routes from a peer.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"cidr": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"interface": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"match_operator": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"prefix_length": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"max": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
												int64validator.AtMost(128),
											},
										},

										"min": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
												int64validator.AtMost(128),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"source": schema.StringAttribute{
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *CrdProjectcalicoOrgBgpfilterV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_crd_projectcalico_org_bgp_filter_v1_manifest")

	var model CrdProjectcalicoOrgBgpfilterV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("crd.projectcalico.org/v1")
	model.Kind = pointer.String("BGPFilter")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
