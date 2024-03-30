/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package crd_projectcalico_org_v1

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
	_ datasource.DataSource = &CrdProjectcalicoOrgHostEndpointV1Manifest{}
)

func NewCrdProjectcalicoOrgHostEndpointV1Manifest() datasource.DataSource {
	return &CrdProjectcalicoOrgHostEndpointV1Manifest{}
}

type CrdProjectcalicoOrgHostEndpointV1Manifest struct{}

type CrdProjectcalicoOrgHostEndpointV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ExpectedIPs   *[]string `tfsdk:"expected_i_ps" json:"expectedIPs,omitempty"`
		InterfaceName *string   `tfsdk:"interface_name" json:"interfaceName,omitempty"`
		Node          *string   `tfsdk:"node" json:"node,omitempty"`
		Ports         *[]struct {
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Port     *int64  `tfsdk:"port" json:"port,omitempty"`
			Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
		} `tfsdk:"ports" json:"ports,omitempty"`
		Profiles *[]string `tfsdk:"profiles" json:"profiles,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CrdProjectcalicoOrgHostEndpointV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_crd_projectcalico_org_host_endpoint_v1_manifest"
}

func (r *CrdProjectcalicoOrgHostEndpointV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "HostEndpointSpec contains the specification for a HostEndpoint resource.",
				MarkdownDescription: "HostEndpointSpec contains the specification for a HostEndpoint resource.",
				Attributes: map[string]schema.Attribute{
					"expected_i_ps": schema.ListAttribute{
						Description:         "The expected IP addresses (IPv4 and IPv6) of the endpoint. If 'InterfaceName' is not present, Calico will look for an interface matching any of the IPs in the list and apply policy to that. Note: 	When using the selector match criteria in an ingress or egress security Policy 	or Profile, Calico converts the selector into a set of IP addresses. For host 	endpoints, the ExpectedIPs field is used for that purpose. (If only the interface 	name is specified, Calico does not learn the IPs of the interface for use in match 	criteria.)",
						MarkdownDescription: "The expected IP addresses (IPv4 and IPv6) of the endpoint. If 'InterfaceName' is not present, Calico will look for an interface matching any of the IPs in the list and apply policy to that. Note: 	When using the selector match criteria in an ingress or egress security Policy 	or Profile, Calico converts the selector into a set of IP addresses. For host 	endpoints, the ExpectedIPs field is used for that purpose. (If only the interface 	name is specified, Calico does not learn the IPs of the interface for use in match 	criteria.)",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"interface_name": schema.StringAttribute{
						Description:         "Either '*', or the name of a specific Linux interface to apply policy to; or empty.  '*' indicates that this HostEndpoint governs all traffic to, from or through the default network namespace of the host named by the 'Node' field; entering and leaving that namespace via any interface, including those from/to non-host-networked local workloads.  If InterfaceName is not '*', this HostEndpoint only governs traffic that enters or leaves the host through the specific interface named by InterfaceName, or - when InterfaceName is empty - through the specific interface that has one of the IPs in ExpectedIPs. Therefore, when InterfaceName is empty, at least one expected IP must be specified.  Only external interfaces (such as 'eth0') are supported here; it isn't possible for a HostEndpoint to protect traffic through a specific local workload interface.  Note: Only some kinds of policy are implemented for '*' HostEndpoints; initially just pre-DNAT policy.  Please check Calico documentation for the latest position.",
						MarkdownDescription: "Either '*', or the name of a specific Linux interface to apply policy to; or empty.  '*' indicates that this HostEndpoint governs all traffic to, from or through the default network namespace of the host named by the 'Node' field; entering and leaving that namespace via any interface, including those from/to non-host-networked local workloads.  If InterfaceName is not '*', this HostEndpoint only governs traffic that enters or leaves the host through the specific interface named by InterfaceName, or - when InterfaceName is empty - through the specific interface that has one of the IPs in ExpectedIPs. Therefore, when InterfaceName is empty, at least one expected IP must be specified.  Only external interfaces (such as 'eth0') are supported here; it isn't possible for a HostEndpoint to protect traffic through a specific local workload interface.  Note: Only some kinds of policy are implemented for '*' HostEndpoints; initially just pre-DNAT policy.  Please check Calico documentation for the latest position.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"node": schema.StringAttribute{
						Description:         "The node name identifying the Calico node instance.",
						MarkdownDescription: "The node name identifying the Calico node instance.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ports": schema.ListNestedAttribute{
						Description:         "Ports contains the endpoint's named ports, which may be referenced in security policy rules.",
						MarkdownDescription: "Ports contains the endpoint's named ports, which may be referenced in security policy rules.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"port": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"protocol": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"profiles": schema.ListAttribute{
						Description:         "A list of identifiers of security Profile objects that apply to this endpoint. Each profile is applied in the order that they appear in this list.  Profile rules are applied after the selector-based security policy.",
						MarkdownDescription: "A list of identifiers of security Profile objects that apply to this endpoint. Each profile is applied in the order that they appear in this list.  Profile rules are applied after the selector-based security policy.",
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
	}
}

func (r *CrdProjectcalicoOrgHostEndpointV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_crd_projectcalico_org_host_endpoint_v1_manifest")

	var model CrdProjectcalicoOrgHostEndpointV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("crd.projectcalico.org/v1")
	model.Kind = pointer.String("HostEndpoint")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
