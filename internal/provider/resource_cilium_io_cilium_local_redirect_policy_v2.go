/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"regexp"

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

type CiliumIoCiliumLocalRedirectPolicyV2Resource struct{}

var (
	_ resource.Resource = (*CiliumIoCiliumLocalRedirectPolicyV2Resource)(nil)
)

type CiliumIoCiliumLocalRedirectPolicyV2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CiliumIoCiliumLocalRedirectPolicyV2GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Description *string `tfsdk:"description" yaml:"description,omitempty"`

		RedirectBackend *struct {
			LocalEndpointSelector *struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"local_endpoint_selector" yaml:"localEndpointSelector,omitempty"`

			ToPorts *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Port *string `tfsdk:"port" yaml:"port,omitempty"`

				Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
			} `tfsdk:"to_ports" yaml:"toPorts,omitempty"`
		} `tfsdk:"redirect_backend" yaml:"redirectBackend,omitempty"`

		RedirectFrontend *struct {
			AddressMatcher *struct {
				Ip *string `tfsdk:"ip" yaml:"ip,omitempty"`

				ToPorts *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Port *string `tfsdk:"port" yaml:"port,omitempty"`

					Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
				} `tfsdk:"to_ports" yaml:"toPorts,omitempty"`
			} `tfsdk:"address_matcher" yaml:"addressMatcher,omitempty"`

			ServiceMatcher *struct {
				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

				ServiceName *string `tfsdk:"service_name" yaml:"serviceName,omitempty"`

				ToPorts *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Port *string `tfsdk:"port" yaml:"port,omitempty"`

					Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
				} `tfsdk:"to_ports" yaml:"toPorts,omitempty"`
			} `tfsdk:"service_matcher" yaml:"serviceMatcher,omitempty"`
		} `tfsdk:"redirect_frontend" yaml:"redirectFrontend,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCiliumIoCiliumLocalRedirectPolicyV2Resource() resource.Resource {
	return &CiliumIoCiliumLocalRedirectPolicyV2Resource{}
}

func (r *CiliumIoCiliumLocalRedirectPolicyV2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cilium_io_cilium_local_redirect_policy_v2"
}

func (r *CiliumIoCiliumLocalRedirectPolicyV2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "CiliumLocalRedirectPolicy is a Kubernetes Custom Resource that contains a specification to redirect traffic locally within a node.",
		MarkdownDescription: "CiliumLocalRedirectPolicy is a Kubernetes Custom Resource that contains a specification to redirect traffic locally within a node.",
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

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
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
				Description:         "Spec is the desired behavior of the local redirect policy.",
				MarkdownDescription: "Spec is the desired behavior of the local redirect policy.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"description": {
						Description:         "Description can be used by the creator of the policy to describe the purpose of this policy.",
						MarkdownDescription: "Description can be used by the creator of the policy to describe the purpose of this policy.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"redirect_backend": {
						Description:         "RedirectBackend specifies backend configuration to redirect traffic to. It can not be empty.",
						MarkdownDescription: "RedirectBackend specifies backend configuration to redirect traffic to. It can not be empty.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"local_endpoint_selector": {
								Description:         "LocalEndpointSelector selects node local pod(s) where traffic is redirected to.",
								MarkdownDescription: "LocalEndpointSelector selects node local pod(s) where traffic is redirected to.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
												},
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

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

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"to_ports": {
								Description:         "ToPorts is a list of L4 ports with protocol of node local pod(s) where traffic is redirected to. When multiple ports are specified, the ports must be named.",
								MarkdownDescription: "ToPorts is a list of L4 ports with protocol of node local pod(s) where traffic is redirected to. When multiple ports are specified, the ports must be named.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name is a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
										MarkdownDescription: "Name is a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]{1,4})|([a-zA-Z0-9]-?)*[a-zA-Z](-?[a-zA-Z0-9])*$`), ""),
										},
									},

									"port": {
										Description:         "Port is an L4 port number. The string will be strictly parsed as a single uint16.",
										MarkdownDescription: "Port is an L4 port number. The string will be strictly parsed as a single uint16.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^()([1-9]|[1-5]?[0-9]{2,4}|6[1-4][0-9]{3}|65[1-4][0-9]{2}|655[1-2][0-9]|6553[1-5])$`), ""),
										},
									},

									"protocol": {
										Description:         "Protocol is the L4 protocol. Accepted values: 'TCP', 'UDP'",
										MarkdownDescription: "Protocol is the L4 protocol. Accepted values: 'TCP', 'UDP'",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("TCP", "UDP"),
										},
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"redirect_frontend": {
						Description:         "RedirectFrontend specifies frontend configuration to redirect traffic from. It can not be empty.",
						MarkdownDescription: "RedirectFrontend specifies frontend configuration to redirect traffic from. It can not be empty.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"address_matcher": {
								Description:         "AddressMatcher is a tuple {IP, port, protocol} that matches traffic to be redirected.",
								MarkdownDescription: "AddressMatcher is a tuple {IP, port, protocol} that matches traffic to be redirected.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ip": {
										Description:         "IP is a destination ip address for traffic to be redirected.  Example: When it is set to '169.254.169.254', traffic destined to '169.254.169.254' is redirected.",
										MarkdownDescription: "IP is a destination ip address for traffic to be redirected.  Example: When it is set to '169.254.169.254', traffic destined to '169.254.169.254' is redirected.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`((^\s*((([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]))\s*$)|(^\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?\s*$))`), ""),
										},
									},

									"to_ports": {
										Description:         "ToPorts is a list of destination L4 ports with protocol for traffic to be redirected. When multiple ports are specified, the ports must be named.  Example: When set to Port: '53' and Protocol: UDP, traffic destined to port '53' with UDP protocol is redirected.",
										MarkdownDescription: "ToPorts is a list of destination L4 ports with protocol for traffic to be redirected. When multiple ports are specified, the ports must be named.  Example: When set to Port: '53' and Protocol: UDP, traffic destined to port '53' with UDP protocol is redirected.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name is a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
												MarkdownDescription: "Name is a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]{1,4})|([a-zA-Z0-9]-?)*[a-zA-Z](-?[a-zA-Z0-9])*$`), ""),
												},
											},

											"port": {
												Description:         "Port is an L4 port number. The string will be strictly parsed as a single uint16.",
												MarkdownDescription: "Port is an L4 port number. The string will be strictly parsed as a single uint16.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^()([1-9]|[1-5]?[0-9]{2,4}|6[1-4][0-9]{3}|65[1-4][0-9]{2}|655[1-2][0-9]|6553[1-5])$`), ""),
												},
											},

											"protocol": {
												Description:         "Protocol is the L4 protocol. Accepted values: 'TCP', 'UDP'",
												MarkdownDescription: "Protocol is the L4 protocol. Accepted values: 'TCP', 'UDP'",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("TCP", "UDP"),
												},
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_matcher": {
								Description:         "ServiceMatcher specifies Kubernetes service and port that matches traffic to be redirected.",
								MarkdownDescription: "ServiceMatcher specifies Kubernetes service and port that matches traffic to be redirected.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"namespace": {
										Description:         "Namespace is the Kubernetes service namespace. The service namespace must match the namespace of the parent Local Redirect Policy.  For Cluster-wide Local Redirect Policy, this can be any namespace.",
										MarkdownDescription: "Namespace is the Kubernetes service namespace. The service namespace must match the namespace of the parent Local Redirect Policy.  For Cluster-wide Local Redirect Policy, this can be any namespace.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"service_name": {
										Description:         "Name is the name of a destination Kubernetes service that identifies traffic to be redirected. The service type needs to be ClusterIP.  Example: When this field is populated with 'serviceName:myService', all the traffic destined to the cluster IP of this service at the (specified) service port(s) will be redirected.",
										MarkdownDescription: "Name is the name of a destination Kubernetes service that identifies traffic to be redirected. The service type needs to be ClusterIP.  Example: When this field is populated with 'serviceName:myService', all the traffic destined to the cluster IP of this service at the (specified) service port(s) will be redirected.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"to_ports": {
										Description:         "ToPorts is a list of destination service L4 ports with protocol for traffic to be redirected. If not specified, traffic for all the service ports will be redirected. When multiple ports are specified, the ports must be named.",
										MarkdownDescription: "ToPorts is a list of destination service L4 ports with protocol for traffic to be redirected. If not specified, traffic for all the service ports will be redirected. When multiple ports are specified, the ports must be named.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name is a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
												MarkdownDescription: "Name is a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]{1,4})|([a-zA-Z0-9]-?)*[a-zA-Z](-?[a-zA-Z0-9])*$`), ""),
												},
											},

											"port": {
												Description:         "Port is an L4 port number. The string will be strictly parsed as a single uint16.",
												MarkdownDescription: "Port is an L4 port number. The string will be strictly parsed as a single uint16.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^()([1-9]|[1-5]?[0-9]{2,4}|6[1-4][0-9]{3}|65[1-4][0-9]{2}|655[1-2][0-9]|6553[1-5])$`), ""),
												},
											},

											"protocol": {
												Description:         "Protocol is the L4 protocol. Accepted values: 'TCP', 'UDP'",
												MarkdownDescription: "Protocol is the L4 protocol. Accepted values: 'TCP', 'UDP'",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("TCP", "UDP"),
												},
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
						}),

						Required: true,
						Optional: false,
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

func (r *CiliumIoCiliumLocalRedirectPolicyV2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_cilium_io_cilium_local_redirect_policy_v2")

	var state CiliumIoCiliumLocalRedirectPolicyV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CiliumIoCiliumLocalRedirectPolicyV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("cilium.io/v2")
	goModel.Kind = utilities.Ptr("CiliumLocalRedirectPolicy")

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

func (r *CiliumIoCiliumLocalRedirectPolicyV2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cilium_io_cilium_local_redirect_policy_v2")
	// NO-OP: All data is already in Terraform state
}

func (r *CiliumIoCiliumLocalRedirectPolicyV2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_cilium_io_cilium_local_redirect_policy_v2")

	var state CiliumIoCiliumLocalRedirectPolicyV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CiliumIoCiliumLocalRedirectPolicyV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("cilium.io/v2")
	goModel.Kind = utilities.Ptr("CiliumLocalRedirectPolicy")

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

func (r *CiliumIoCiliumLocalRedirectPolicyV2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_cilium_io_cilium_local_redirect_policy_v2")
	// NO-OP: Terraform removes the state automatically for us
}
