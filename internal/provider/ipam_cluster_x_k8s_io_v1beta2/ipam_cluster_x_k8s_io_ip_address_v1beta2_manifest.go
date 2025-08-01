/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ipam_cluster_x_k8s_io_v1beta2

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &IpamClusterXK8SIoIpaddressV1Beta2Manifest{}
)

func NewIpamClusterXK8SIoIpaddressV1Beta2Manifest() datasource.DataSource {
	return &IpamClusterXK8SIoIpaddressV1Beta2Manifest{}
}

type IpamClusterXK8SIoIpaddressV1Beta2Manifest struct{}

type IpamClusterXK8SIoIpaddressV1Beta2ManifestData struct {
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
		Address  *string `tfsdk:"address" json:"address,omitempty"`
		ClaimRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"claim_ref" json:"claimRef,omitempty"`
		Gateway *string `tfsdk:"gateway" json:"gateway,omitempty"`
		PoolRef *struct {
			ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
			Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"pool_ref" json:"poolRef,omitempty"`
		Prefix *int64 `tfsdk:"prefix" json:"prefix,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *IpamClusterXK8SIoIpaddressV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ipam_cluster_x_k8s_io_ip_address_v1beta2_manifest"
}

func (r *IpamClusterXK8SIoIpaddressV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "IPAddress is the Schema for the ipaddress API.",
		MarkdownDescription: "IPAddress is the Schema for the ipaddress API.",
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
				Description:         "spec is the desired state of IPAddress.",
				MarkdownDescription: "spec is the desired state of IPAddress.",
				Attributes: map[string]schema.Attribute{
					"address": schema.StringAttribute{
						Description:         "address is the IP address.",
						MarkdownDescription: "address is the IP address.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.LengthAtMost(39),
						},
					},

					"claim_ref": schema.SingleNestedAttribute{
						Description:         "claimRef is a reference to the claim this IPAddress was created for.",
						MarkdownDescription: "claimRef is a reference to the claim this IPAddress was created for.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name of the IPAddressClaim. name must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character.",
								MarkdownDescription: "name of the IPAddressClaim. name must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"gateway": schema.StringAttribute{
						Description:         "gateway is the network gateway of the network the address is from.",
						MarkdownDescription: "gateway is the network gateway of the network the address is from.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.LengthAtMost(39),
						},
					},

					"pool_ref": schema.SingleNestedAttribute{
						Description:         "poolRef is a reference to the pool that this IPAddress was created from.",
						MarkdownDescription: "poolRef is a reference to the pool that this IPAddress was created from.",
						Attributes: map[string]schema.Attribute{
							"api_group": schema.StringAttribute{
								Description:         "apiGroup of the IPPool. apiGroup must be fully qualified domain name.",
								MarkdownDescription: "apiGroup of the IPPool. apiGroup must be fully qualified domain name.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
								},
							},

							"kind": schema.StringAttribute{
								Description:         "kind of the IPPool. kind must consist of alphanumeric characters or '-', start with an alphabetic character, and end with an alphanumeric character.",
								MarkdownDescription: "kind of the IPPool. kind must consist of alphanumeric characters or '-', start with an alphabetic character, and end with an alphanumeric character.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
								},
							},

							"name": schema.StringAttribute{
								Description:         "name of the IPPool. name must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character.",
								MarkdownDescription: "name of the IPPool. name must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"prefix": schema.Int64Attribute{
						Description:         "prefix is the prefix of the address.",
						MarkdownDescription: "prefix is the prefix of the address.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *IpamClusterXK8SIoIpaddressV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ipam_cluster_x_k8s_io_ip_address_v1beta2_manifest")

	var model IpamClusterXK8SIoIpaddressV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("ipam.cluster.x-k8s.io/v1beta2")
	model.Kind = pointer.String("IPAddress")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
