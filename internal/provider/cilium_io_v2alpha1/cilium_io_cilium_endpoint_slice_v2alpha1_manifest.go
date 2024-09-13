/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cilium_io_v2alpha1

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
	_ datasource.DataSource = &CiliumIoCiliumEndpointSliceV2Alpha1Manifest{}
)

func NewCiliumIoCiliumEndpointSliceV2Alpha1Manifest() datasource.DataSource {
	return &CiliumIoCiliumEndpointSliceV2Alpha1Manifest{}
}

type CiliumIoCiliumEndpointSliceV2Alpha1Manifest struct{}

type CiliumIoCiliumEndpointSliceV2Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Endpoints *[]struct {
		Encryption *struct {
			Key *int64 `tfsdk:"key" json:"key,omitempty"`
		} `tfsdk:"encryption" json:"encryption,omitempty"`
		Id          *int64  `tfsdk:"id" json:"id,omitempty"`
		Name        *string `tfsdk:"name" json:"name,omitempty"`
		Named_ports *[]struct {
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Port     *int64  `tfsdk:"port" json:"port,omitempty"`
			Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
		} `tfsdk:"named_ports" json:"named-ports,omitempty"`
		Networking *struct {
			Addressing *[]struct {
				Ipv4 *string `tfsdk:"ipv4" json:"ipv4,omitempty"`
				Ipv6 *string `tfsdk:"ipv6" json:"ipv6,omitempty"`
			} `tfsdk:"addressing" json:"addressing,omitempty"`
			Node *string `tfsdk:"node" json:"node,omitempty"`
		} `tfsdk:"networking" json:"networking,omitempty"`
	} `tfsdk:"endpoints" json:"endpoints,omitempty"`
	Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
}

func (r *CiliumIoCiliumEndpointSliceV2Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cilium_io_cilium_endpoint_slice_v2alpha1_manifest"
}

func (r *CiliumIoCiliumEndpointSliceV2Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CiliumEndpointSlice contains a group of CoreCiliumendpoints.",
		MarkdownDescription: "CiliumEndpointSlice contains a group of CoreCiliumendpoints.",
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

			"endpoints": schema.ListNestedAttribute{
				Description:         "Endpoints is a list of coreCEPs packed in a CiliumEndpointSlice",
				MarkdownDescription: "Endpoints is a list of coreCEPs packed in a CiliumEndpointSlice",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"encryption": schema.SingleNestedAttribute{
							Description:         "EncryptionSpec defines the encryption relevant configuration of a node.",
							MarkdownDescription: "EncryptionSpec defines the encryption relevant configuration of a node.",
							Attributes: map[string]schema.Attribute{
								"key": schema.Int64Attribute{
									Description:         "Key is the index to the key to use for encryption or 0 if encryption is disabled.",
									MarkdownDescription: "Key is the index to the key to use for encryption or 0 if encryption is disabled.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
							Required: false,
							Optional: true,
							Computed: false,
						},

						"id": schema.Int64Attribute{
							Description:         "IdentityID is the numeric identity of the endpoint",
							MarkdownDescription: "IdentityID is the numeric identity of the endpoint",
							Required:            false,
							Optional:            true,
							Computed:            false,
						},

						"name": schema.StringAttribute{
							Description:         "Name indicate as CiliumEndpoint name.",
							MarkdownDescription: "Name indicate as CiliumEndpoint name.",
							Required:            false,
							Optional:            true,
							Computed:            false,
						},

						"named_ports": schema.ListNestedAttribute{
							Description:         "NamedPorts List of named Layer 4 port and protocol pairs which will be used in Network Policy specs. swagger:model NamedPorts",
							MarkdownDescription: "NamedPorts List of named Layer 4 port and protocol pairs which will be used in Network Policy specs. swagger:model NamedPorts",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Optional layer 4 port name",
										MarkdownDescription: "Optional layer 4 port name",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port": schema.Int64Attribute{
										Description:         "Layer 4 port number",
										MarkdownDescription: "Layer 4 port number",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"protocol": schema.StringAttribute{
										Description:         "Layer 4 protocol Enum: [TCP UDP SCTP ICMP ICMPV6 ANY]",
										MarkdownDescription: "Layer 4 protocol Enum: [TCP UDP SCTP ICMP ICMPV6 ANY]",
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

						"networking": schema.SingleNestedAttribute{
							Description:         "EndpointNetworking is the addressing information of an endpoint.",
							MarkdownDescription: "EndpointNetworking is the addressing information of an endpoint.",
							Attributes: map[string]schema.Attribute{
								"addressing": schema.ListNestedAttribute{
									Description:         "IP4/6 addresses assigned to this Endpoint",
									MarkdownDescription: "IP4/6 addresses assigned to this Endpoint",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"ipv4": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ipv6": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"node": schema.StringAttribute{
									Description:         "NodeIP is the IP of the node the endpoint is running on. The IP must be reachable between nodes.",
									MarkdownDescription: "NodeIP is the IP of the node the endpoint is running on. The IP must be reachable between nodes.",
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
				Required: true,
				Optional: false,
				Computed: false,
			},

			"namespace": schema.StringAttribute{
				Description:         "Namespace indicate as CiliumEndpointSlice namespace. All the CiliumEndpoints within the same namespace are put together in CiliumEndpointSlice.",
				MarkdownDescription: "Namespace indicate as CiliumEndpointSlice namespace. All the CiliumEndpoints within the same namespace are put together in CiliumEndpointSlice.",
				Required:            false,
				Optional:            true,
				Computed:            false,
			},
		},
	}
}

func (r *CiliumIoCiliumEndpointSliceV2Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cilium_io_cilium_endpoint_slice_v2alpha1_manifest")

	var model CiliumIoCiliumEndpointSliceV2Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cilium.io/v2alpha1")
	model.Kind = pointer.String("CiliumEndpointSlice")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
