/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ipam_cluster_x_k8s_io_v1beta1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &IpamClusterXK8SIoIpaddressV1Beta1Manifest{}
)

func NewIpamClusterXK8SIoIpaddressV1Beta1Manifest() datasource.DataSource {
	return &IpamClusterXK8SIoIpaddressV1Beta1Manifest{}
}

type IpamClusterXK8SIoIpaddressV1Beta1Manifest struct{}

type IpamClusterXK8SIoIpaddressV1Beta1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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

func (r *IpamClusterXK8SIoIpaddressV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ipam_cluster_x_k8s_io_ip_address_v1beta1_manifest"
}

func (r *IpamClusterXK8SIoIpaddressV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "IPAddress is the Schema for the ipaddress API.",
		MarkdownDescription: "IPAddress is the Schema for the ipaddress API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "IPAddressSpec is the desired state of an IPAddress.",
				MarkdownDescription: "IPAddressSpec is the desired state of an IPAddress.",
				Attributes: map[string]schema.Attribute{
					"address": schema.StringAttribute{
						Description:         "Address is the IP address.",
						MarkdownDescription: "Address is the IP address.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"claim_ref": schema.SingleNestedAttribute{
						Description:         "ClaimRef is a reference to the claim this IPAddress was created for.",
						MarkdownDescription: "ClaimRef is a reference to the claim this IPAddress was created for.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"gateway": schema.StringAttribute{
						Description:         "Gateway is the network gateway of the network the address is from.",
						MarkdownDescription: "Gateway is the network gateway of the network the address is from.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pool_ref": schema.SingleNestedAttribute{
						Description:         "PoolRef is a reference to the pool that this IPAddress was created from.",
						MarkdownDescription: "PoolRef is a reference to the pool that this IPAddress was created from.",
						Attributes: map[string]schema.Attribute{
							"api_group": schema.StringAttribute{
								Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
								MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind is the type of resource being referenced",
								MarkdownDescription: "Kind is the type of resource being referenced",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of resource being referenced",
								MarkdownDescription: "Name is the name of resource being referenced",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"prefix": schema.Int64Attribute{
						Description:         "Prefix is the prefix of the address.",
						MarkdownDescription: "Prefix is the prefix of the address.",
						Required:            true,
						Optional:            false,
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

func (r *IpamClusterXK8SIoIpaddressV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ipam_cluster_x_k8s_io_ip_address_v1beta1_manifest")

	var model IpamClusterXK8SIoIpaddressV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("ipam.cluster.x-k8s.io/v1beta1")
	model.Kind = pointer.String("IPAddress")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
