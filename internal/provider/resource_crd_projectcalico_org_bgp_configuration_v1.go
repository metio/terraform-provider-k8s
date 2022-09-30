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

type CrdProjectcalicoOrgBGPConfigurationV1Resource struct{}

var (
	_ resource.Resource = (*CrdProjectcalicoOrgBGPConfigurationV1Resource)(nil)
)

type CrdProjectcalicoOrgBGPConfigurationV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CrdProjectcalicoOrgBGPConfigurationV1GoModel struct {
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
		NodeToNodeMeshEnabled *bool `tfsdk:"node_to_node_mesh_enabled" yaml:"nodeToNodeMeshEnabled,omitempty"`

		ServiceExternalIPs *[]struct {
			Cidr *string `tfsdk:"cidr" yaml:"cidr,omitempty"`
		} `tfsdk:"service_external_i_ps" yaml:"serviceExternalIPs,omitempty"`

		AsNumber *int64 `tfsdk:"as_number" yaml:"asNumber,omitempty"`

		BindMode *string `tfsdk:"bind_mode" yaml:"bindMode,omitempty"`

		Communities *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"communities" yaml:"communities,omitempty"`

		NodeMeshPassword *struct {
			SecretKeyRef *struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
			} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
		} `tfsdk:"node_mesh_password" yaml:"nodeMeshPassword,omitempty"`

		PrefixAdvertisements *[]struct {
			Cidr *string `tfsdk:"cidr" yaml:"cidr,omitempty"`

			Communities *[]string `tfsdk:"communities" yaml:"communities,omitempty"`
		} `tfsdk:"prefix_advertisements" yaml:"prefixAdvertisements,omitempty"`

		ServiceClusterIPs *[]struct {
			Cidr *string `tfsdk:"cidr" yaml:"cidr,omitempty"`
		} `tfsdk:"service_cluster_i_ps" yaml:"serviceClusterIPs,omitempty"`

		ServiceLoadBalancerIPs *[]struct {
			Cidr *string `tfsdk:"cidr" yaml:"cidr,omitempty"`
		} `tfsdk:"service_load_balancer_i_ps" yaml:"serviceLoadBalancerIPs,omitempty"`

		ListenPort *int64 `tfsdk:"listen_port" yaml:"listenPort,omitempty"`

		LogSeverityScreen *string `tfsdk:"log_severity_screen" yaml:"logSeverityScreen,omitempty"`

		NodeMeshMaxRestartTime *string `tfsdk:"node_mesh_max_restart_time" yaml:"nodeMeshMaxRestartTime,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCrdProjectcalicoOrgBGPConfigurationV1Resource() resource.Resource {
	return &CrdProjectcalicoOrgBGPConfigurationV1Resource{}
}

func (r *CrdProjectcalicoOrgBGPConfigurationV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_crd_projectcalico_org_bgp_configuration_v1"
}

func (r *CrdProjectcalicoOrgBGPConfigurationV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "BGPConfiguration contains the configuration for any BGP routing.",
		MarkdownDescription: "BGPConfiguration contains the configuration for any BGP routing.",
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
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
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
				Description:         "BGPConfigurationSpec contains the values of the BGP configuration.",
				MarkdownDescription: "BGPConfigurationSpec contains the values of the BGP configuration.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"node_to_node_mesh_enabled": {
						Description:         "NodeToNodeMeshEnabled sets whether full node to node BGP mesh is enabled. [Default: true]",
						MarkdownDescription: "NodeToNodeMeshEnabled sets whether full node to node BGP mesh is enabled. [Default: true]",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_external_i_ps": {
						Description:         "ServiceExternalIPs are the CIDR blocks for Kubernetes Service External IPs. Kubernetes Service ExternalIPs will only be advertised if they are within one of these blocks.",
						MarkdownDescription: "ServiceExternalIPs are the CIDR blocks for Kubernetes Service External IPs. Kubernetes Service ExternalIPs will only be advertised if they are within one of these blocks.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"cidr": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"as_number": {
						Description:         "ASNumber is the default AS number used by a node. [Default: 64512]",
						MarkdownDescription: "ASNumber is the default AS number used by a node. [Default: 64512]",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bind_mode": {
						Description:         "BindMode indicates whether to listen for BGP connections on all addresses (None) or only on the node's canonical IP address Node.Spec.BGP.IPvXAddress (NodeIP). Default behaviour is to listen for BGP connections on all addresses.",
						MarkdownDescription: "BindMode indicates whether to listen for BGP connections on all addresses (None) or only on the node's canonical IP address Node.Spec.BGP.IPvXAddress (NodeIP). Default behaviour is to listen for BGP connections on all addresses.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"communities": {
						Description:         "Communities is a list of BGP community values and their arbitrary names for tagging routes.",
						MarkdownDescription: "Communities is a list of BGP community values and their arbitrary names for tagging routes.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name given to community value.",
								MarkdownDescription: "Name given to community value.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"value": {
								Description:         "Value must be of format 'aa:nn' or 'aa:nn:mm'. For standard community use 'aa:nn' format, where 'aa' and 'nn' are 16 bit number. For large community use 'aa:nn:mm' format, where 'aa', 'nn' and 'mm' are 32 bit number. Where, 'aa' is an AS Number, 'nn' and 'mm' are per-AS identifier.",
								MarkdownDescription: "Value must be of format 'aa:nn' or 'aa:nn:mm'. For standard community use 'aa:nn' format, where 'aa' and 'nn' are 16 bit number. For large community use 'aa:nn:mm' format, where 'aa', 'nn' and 'mm' are 32 bit number. Where, 'aa' is an AS Number, 'nn' and 'mm' are per-AS identifier.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_mesh_password": {
						Description:         "Optional BGP password for full node-to-mesh peerings. This field can only be set on the default BGPConfiguration instance and requires that NodeMesh is enabled",
						MarkdownDescription: "Optional BGP password for full node-to-mesh peerings. This field can only be set on the default BGPConfiguration instance and requires that NodeMesh is enabled",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"secret_key_ref": {
								Description:         "Selects a key of a secret in the node pod's namespace.",
								MarkdownDescription: "Selects a key of a secret in the node pod's namespace.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": {
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"optional": {
										Description:         "Specify whether the Secret or its key must be defined",
										MarkdownDescription: "Specify whether the Secret or its key must be defined",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"prefix_advertisements": {
						Description:         "PrefixAdvertisements contains per-prefix advertisement configuration.",
						MarkdownDescription: "PrefixAdvertisements contains per-prefix advertisement configuration.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"cidr": {
								Description:         "CIDR for which properties should be advertised.",
								MarkdownDescription: "CIDR for which properties should be advertised.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"communities": {
								Description:         "Communities can be list of either community names already defined in 'Specs.Communities' or community value of format 'aa:nn' or 'aa:nn:mm'. For standard community use 'aa:nn' format, where 'aa' and 'nn' are 16 bit number. For large community use 'aa:nn:mm' format, where 'aa', 'nn' and 'mm' are 32 bit number. Where,'aa' is an AS Number, 'nn' and 'mm' are per-AS identifier.",
								MarkdownDescription: "Communities can be list of either community names already defined in 'Specs.Communities' or community value of format 'aa:nn' or 'aa:nn:mm'. For standard community use 'aa:nn' format, where 'aa' and 'nn' are 16 bit number. For large community use 'aa:nn:mm' format, where 'aa', 'nn' and 'mm' are 32 bit number. Where,'aa' is an AS Number, 'nn' and 'mm' are per-AS identifier.",

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

					"service_cluster_i_ps": {
						Description:         "ServiceClusterIPs are the CIDR blocks from which service cluster IPs are allocated. If specified, Calico will advertise these blocks, as well as any cluster IPs within them.",
						MarkdownDescription: "ServiceClusterIPs are the CIDR blocks from which service cluster IPs are allocated. If specified, Calico will advertise these blocks, as well as any cluster IPs within them.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"cidr": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_load_balancer_i_ps": {
						Description:         "ServiceLoadBalancerIPs are the CIDR blocks for Kubernetes Service LoadBalancer IPs. Kubernetes Service status.LoadBalancer.Ingress IPs will only be advertised if they are within one of these blocks.",
						MarkdownDescription: "ServiceLoadBalancerIPs are the CIDR blocks for Kubernetes Service LoadBalancer IPs. Kubernetes Service status.LoadBalancer.Ingress IPs will only be advertised if they are within one of these blocks.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"cidr": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"listen_port": {
						Description:         "ListenPort is the port where BGP protocol should listen. Defaults to 179",
						MarkdownDescription: "ListenPort is the port where BGP protocol should listen. Defaults to 179",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"log_severity_screen": {
						Description:         "LogSeverityScreen is the log severity above which logs are sent to the stdout. [Default: INFO]",
						MarkdownDescription: "LogSeverityScreen is the log severity above which logs are sent to the stdout. [Default: INFO]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_mesh_max_restart_time": {
						Description:         "Time to allow for software restart for node-to-mesh peerings.  When specified, this is configured as the graceful restart timeout.  When not specified, the BIRD default of 120s is used. This field can only be set on the default BGPConfiguration instance and requires that NodeMesh is enabled",
						MarkdownDescription: "Time to allow for software restart for node-to-mesh peerings.  When specified, this is configured as the graceful restart timeout.  When not specified, the BIRD default of 120s is used. This field can only be set on the default BGPConfiguration instance and requires that NodeMesh is enabled",

						Type: types.StringType,

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

func (r *CrdProjectcalicoOrgBGPConfigurationV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_crd_projectcalico_org_bgp_configuration_v1")

	var state CrdProjectcalicoOrgBGPConfigurationV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CrdProjectcalicoOrgBGPConfigurationV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("crd.projectcalico.org/v1")
	goModel.Kind = utilities.Ptr("BGPConfiguration")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CrdProjectcalicoOrgBGPConfigurationV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_crd_projectcalico_org_bgp_configuration_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *CrdProjectcalicoOrgBGPConfigurationV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_crd_projectcalico_org_bgp_configuration_v1")

	var state CrdProjectcalicoOrgBGPConfigurationV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CrdProjectcalicoOrgBGPConfigurationV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("crd.projectcalico.org/v1")
	goModel.Kind = utilities.Ptr("BGPConfiguration")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CrdProjectcalicoOrgBGPConfigurationV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_crd_projectcalico_org_bgp_configuration_v1")
	// NO-OP: Terraform removes the state automatically for us
}
