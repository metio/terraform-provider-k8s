/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type CrdProjectcalicoOrgHostEndpointV1Resource struct{}

var (
	_ resource.Resource = (*CrdProjectcalicoOrgHostEndpointV1Resource)(nil)
)

type CrdProjectcalicoOrgHostEndpointV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CrdProjectcalicoOrgHostEndpointV1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		ExpectedIPs *[]string `tfsdk:"expected_i_ps" yaml:"expectedIPs,omitempty"`

		InterfaceName *string `tfsdk:"interface_name" yaml:"interfaceName,omitempty"`

		Node *string `tfsdk:"node" yaml:"node,omitempty"`

		Ports *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

			Protocol utilities.IntOrString `tfsdk:"protocol" yaml:"protocol,omitempty"`
		} `tfsdk:"ports" yaml:"ports,omitempty"`

		Profiles *[]string `tfsdk:"profiles" yaml:"profiles,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCrdProjectcalicoOrgHostEndpointV1Resource() resource.Resource {
	return &CrdProjectcalicoOrgHostEndpointV1Resource{}
}

func (r *CrdProjectcalicoOrgHostEndpointV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_crd_projectcalico_org_host_endpoint_v1"
}

func (r *CrdProjectcalicoOrgHostEndpointV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "",
		MarkdownDescription: "",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "HostEndpointSpec contains the specification for a HostEndpoint resource.",
				MarkdownDescription: "HostEndpointSpec contains the specification for a HostEndpoint resource.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"expected_i_ps": {
						Description:         "The expected IP addresses (IPv4 and IPv6) of the endpoint. If 'InterfaceName' is not present, Calico will look for an interface matching any of the IPs in the list and apply policy to that. Note: 	When using the selector match criteria in an ingress or egress security Policy 	or Profile, Calico converts the selector into a set of IP addresses. For host 	endpoints, the ExpectedIPs field is used for that purpose. (If only the interface 	name is specified, Calico does not learn the IPs of the interface for use in match 	criteria.)",
						MarkdownDescription: "The expected IP addresses (IPv4 and IPv6) of the endpoint. If 'InterfaceName' is not present, Calico will look for an interface matching any of the IPs in the list and apply policy to that. Note: 	When using the selector match criteria in an ingress or egress security Policy 	or Profile, Calico converts the selector into a set of IP addresses. For host 	endpoints, the ExpectedIPs field is used for that purpose. (If only the interface 	name is specified, Calico does not learn the IPs of the interface for use in match 	criteria.)",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"interface_name": {
						Description:         "Either '*', or the name of a specific Linux interface to apply policy to; or empty.  '*' indicates that this HostEndpoint governs all traffic to, from or through the default network namespace of the host named by the 'Node' field; entering and leaving that namespace via any interface, including those from/to non-host-networked local workloads.  If InterfaceName is not '*', this HostEndpoint only governs traffic that enters or leaves the host through the specific interface named by InterfaceName, or - when InterfaceName is empty - through the specific interface that has one of the IPs in ExpectedIPs. Therefore, when InterfaceName is empty, at least one expected IP must be specified.  Only external interfaces (such as 'eth0') are supported here; it isn't possible for a HostEndpoint to protect traffic through a specific local workload interface.  Note: Only some kinds of policy are implemented for '*' HostEndpoints; initially just pre-DNAT policy.  Please check Calico documentation for the latest position.",
						MarkdownDescription: "Either '*', or the name of a specific Linux interface to apply policy to; or empty.  '*' indicates that this HostEndpoint governs all traffic to, from or through the default network namespace of the host named by the 'Node' field; entering and leaving that namespace via any interface, including those from/to non-host-networked local workloads.  If InterfaceName is not '*', this HostEndpoint only governs traffic that enters or leaves the host through the specific interface named by InterfaceName, or - when InterfaceName is empty - through the specific interface that has one of the IPs in ExpectedIPs. Therefore, when InterfaceName is empty, at least one expected IP must be specified.  Only external interfaces (such as 'eth0') are supported here; it isn't possible for a HostEndpoint to protect traffic through a specific local workload interface.  Note: Only some kinds of policy are implemented for '*' HostEndpoints; initially just pre-DNAT policy.  Please check Calico documentation for the latest position.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"node": {
						Description:         "The node name identifying the Calico node instance.",
						MarkdownDescription: "The node name identifying the Calico node instance.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ports": {
						Description:         "Ports contains the endpoint's named ports, which may be referenced in security policy rules.",
						MarkdownDescription: "Ports contains the endpoint's named ports, which may be referenced in security policy rules.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"port": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"protocol": {
								Description:         "",
								MarkdownDescription: "",

								Type: utilities.IntOrStringType{},

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"profiles": {
						Description:         "A list of identifiers of security Profile objects that apply to this endpoint. Each profile is applied in the order that they appear in this list.  Profile rules are applied after the selector-based security policy.",
						MarkdownDescription: "A list of identifiers of security Profile objects that apply to this endpoint. Each profile is applied in the order that they appear in this list.  Profile rules are applied after the selector-based security policy.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *CrdProjectcalicoOrgHostEndpointV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_crd_projectcalico_org_host_endpoint_v1")

	var state CrdProjectcalicoOrgHostEndpointV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CrdProjectcalicoOrgHostEndpointV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("crd.projectcalico.org/v1")
	goModel.Kind = utilities.Ptr("HostEndpoint")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CrdProjectcalicoOrgHostEndpointV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_crd_projectcalico_org_host_endpoint_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *CrdProjectcalicoOrgHostEndpointV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_crd_projectcalico_org_host_endpoint_v1")

	var state CrdProjectcalicoOrgHostEndpointV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CrdProjectcalicoOrgHostEndpointV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("crd.projectcalico.org/v1")
	goModel.Kind = utilities.Ptr("HostEndpoint")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CrdProjectcalicoOrgHostEndpointV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_crd_projectcalico_org_host_endpoint_v1")
	// NO-OP: Terraform removes the state automatically for us
}
