/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cilium_io_v2

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
	_ datasource.DataSource = &CiliumIoCiliumLocalRedirectPolicyV2Manifest{}
)

func NewCiliumIoCiliumLocalRedirectPolicyV2Manifest() datasource.DataSource {
	return &CiliumIoCiliumLocalRedirectPolicyV2Manifest{}
}

type CiliumIoCiliumLocalRedirectPolicyV2Manifest struct{}

type CiliumIoCiliumLocalRedirectPolicyV2ManifestData struct {
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
		Description     *string `tfsdk:"description" json:"description,omitempty"`
		RedirectBackend *struct {
			LocalEndpointSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"local_endpoint_selector" json:"localEndpointSelector,omitempty"`
			ToPorts *[]struct {
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Port     *string `tfsdk:"port" json:"port,omitempty"`
				Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
			} `tfsdk:"to_ports" json:"toPorts,omitempty"`
		} `tfsdk:"redirect_backend" json:"redirectBackend,omitempty"`
		RedirectFrontend *struct {
			AddressMatcher *struct {
				Ip      *string `tfsdk:"ip" json:"ip,omitempty"`
				ToPorts *[]struct {
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Port     *string `tfsdk:"port" json:"port,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
				} `tfsdk:"to_ports" json:"toPorts,omitempty"`
			} `tfsdk:"address_matcher" json:"addressMatcher,omitempty"`
			ServiceMatcher *struct {
				Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
				ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
				ToPorts     *[]struct {
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Port     *string `tfsdk:"port" json:"port,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
				} `tfsdk:"to_ports" json:"toPorts,omitempty"`
			} `tfsdk:"service_matcher" json:"serviceMatcher,omitempty"`
		} `tfsdk:"redirect_frontend" json:"redirectFrontend,omitempty"`
		SkipRedirectFromBackend *bool `tfsdk:"skip_redirect_from_backend" json:"skipRedirectFromBackend,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CiliumIoCiliumLocalRedirectPolicyV2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cilium_io_cilium_local_redirect_policy_v2_manifest"
}

func (r *CiliumIoCiliumLocalRedirectPolicyV2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CiliumLocalRedirectPolicy is a Kubernetes Custom Resource that contains a specification to redirect traffic locally within a node.",
		MarkdownDescription: "CiliumLocalRedirectPolicy is a Kubernetes Custom Resource that contains a specification to redirect traffic locally within a node.",
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
				Description:         "Spec is the desired behavior of the local redirect policy.",
				MarkdownDescription: "Spec is the desired behavior of the local redirect policy.",
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description:         "Description can be used by the creator of the policy to describe the purpose of this policy.",
						MarkdownDescription: "Description can be used by the creator of the policy to describe the purpose of this policy.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"redirect_backend": schema.SingleNestedAttribute{
						Description:         "RedirectBackend specifies backend configuration to redirect traffic to. It can not be empty.",
						MarkdownDescription: "RedirectBackend specifies backend configuration to redirect traffic to. It can not be empty.",
						Attributes: map[string]schema.Attribute{
							"local_endpoint_selector": schema.SingleNestedAttribute{
								Description:         "LocalEndpointSelector selects node local pod(s) where traffic is redirected to.",
								MarkdownDescription: "LocalEndpointSelector selects node local pod(s) where traffic is redirected to.",
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
								Required: true,
								Optional: false,
								Computed: false,
							},

							"to_ports": schema.ListNestedAttribute{
								Description:         "ToPorts is a list of L4 ports with protocol of node local pod(s) where traffic is redirected to. When multiple ports are specified, the ports must be named.",
								MarkdownDescription: "ToPorts is a list of L4 ports with protocol of node local pod(s) where traffic is redirected to. When multiple ports are specified, the ports must be named.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name is a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
											MarkdownDescription: "Name is a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]{1,4})|([a-zA-Z0-9]-?)*[a-zA-Z](-?[a-zA-Z0-9])*$`), ""),
											},
										},

										"port": schema.StringAttribute{
											Description:         "Port is an L4 port number. The string will be strictly parsed as a single uint16.",
											MarkdownDescription: "Port is an L4 port number. The string will be strictly parsed as a single uint16.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^()([1-9]|[1-5]?[0-9]{2,4}|6[1-4][0-9]{3}|65[1-4][0-9]{2}|655[1-2][0-9]|6553[1-5])$`), ""),
											},
										},

										"protocol": schema.StringAttribute{
											Description:         "Protocol is the L4 protocol. Accepted values: 'TCP', 'UDP'",
											MarkdownDescription: "Protocol is the L4 protocol. Accepted values: 'TCP', 'UDP'",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("TCP", "UDP"),
											},
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

					"redirect_frontend": schema.SingleNestedAttribute{
						Description:         "RedirectFrontend specifies frontend configuration to redirect traffic from. It can not be empty.",
						MarkdownDescription: "RedirectFrontend specifies frontend configuration to redirect traffic from. It can not be empty.",
						Attributes: map[string]schema.Attribute{
							"address_matcher": schema.SingleNestedAttribute{
								Description:         "AddressMatcher is a tuple {IP, port, protocol} that matches traffic to be redirected.",
								MarkdownDescription: "AddressMatcher is a tuple {IP, port, protocol} that matches traffic to be redirected.",
								Attributes: map[string]schema.Attribute{
									"ip": schema.StringAttribute{
										Description:         "IP is a destination ip address for traffic to be redirected. Example: When it is set to '169.254.169.254', traffic destined to '169.254.169.254' is redirected.",
										MarkdownDescription: "IP is a destination ip address for traffic to be redirected. Example: When it is set to '169.254.169.254', traffic destined to '169.254.169.254' is redirected.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`((^\s*((([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]))\s*$)|(^\s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)(\.(25[0-5]|2[0-4]\d|1\d\d|[1-9]?\d)){3}))|:)))(%.+)?\s*$))`), ""),
										},
									},

									"to_ports": schema.ListNestedAttribute{
										Description:         "ToPorts is a list of destination L4 ports with protocol for traffic to be redirected. When multiple ports are specified, the ports must be named. Example: When set to Port: '53' and Protocol: UDP, traffic destined to port '53' with UDP protocol is redirected.",
										MarkdownDescription: "ToPorts is a list of destination L4 ports with protocol for traffic to be redirected. When multiple ports are specified, the ports must be named. Example: When set to Port: '53' and Protocol: UDP, traffic destined to port '53' with UDP protocol is redirected.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name is a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
													MarkdownDescription: "Name is a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]{1,4})|([a-zA-Z0-9]-?)*[a-zA-Z](-?[a-zA-Z0-9])*$`), ""),
													},
												},

												"port": schema.StringAttribute{
													Description:         "Port is an L4 port number. The string will be strictly parsed as a single uint16.",
													MarkdownDescription: "Port is an L4 port number. The string will be strictly parsed as a single uint16.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^()([1-9]|[1-5]?[0-9]{2,4}|6[1-4][0-9]{3}|65[1-4][0-9]{2}|655[1-2][0-9]|6553[1-5])$`), ""),
													},
												},

												"protocol": schema.StringAttribute{
													Description:         "Protocol is the L4 protocol. Accepted values: 'TCP', 'UDP'",
													MarkdownDescription: "Protocol is the L4 protocol. Accepted values: 'TCP', 'UDP'",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("TCP", "UDP"),
													},
												},
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_matcher": schema.SingleNestedAttribute{
								Description:         "ServiceMatcher specifies Kubernetes service and port that matches traffic to be redirected.",
								MarkdownDescription: "ServiceMatcher specifies Kubernetes service and port that matches traffic to be redirected.",
								Attributes: map[string]schema.Attribute{
									"namespace": schema.StringAttribute{
										Description:         "Namespace is the Kubernetes service namespace. The service namespace must match the namespace of the parent Local Redirect Policy. For Cluster-wide Local Redirect Policy, this can be any namespace.",
										MarkdownDescription: "Namespace is the Kubernetes service namespace. The service namespace must match the namespace of the parent Local Redirect Policy. For Cluster-wide Local Redirect Policy, this can be any namespace.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"service_name": schema.StringAttribute{
										Description:         "Name is the name of a destination Kubernetes service that identifies traffic to be redirected. The service type needs to be ClusterIP. Example: When this field is populated with 'serviceName:myService', all the traffic destined to the cluster IP of this service at the (specified) service port(s) will be redirected.",
										MarkdownDescription: "Name is the name of a destination Kubernetes service that identifies traffic to be redirected. The service type needs to be ClusterIP. Example: When this field is populated with 'serviceName:myService', all the traffic destined to the cluster IP of this service at the (specified) service port(s) will be redirected.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"to_ports": schema.ListNestedAttribute{
										Description:         "ToPorts is a list of destination service L4 ports with protocol for traffic to be redirected. If not specified, traffic for all the service ports will be redirected. When multiple ports are specified, the ports must be named.",
										MarkdownDescription: "ToPorts is a list of destination service L4 ports with protocol for traffic to be redirected. If not specified, traffic for all the service ports will be redirected. When multiple ports are specified, the ports must be named.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name is a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
													MarkdownDescription: "Name is a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]{1,4})|([a-zA-Z0-9]-?)*[a-zA-Z](-?[a-zA-Z0-9])*$`), ""),
													},
												},

												"port": schema.StringAttribute{
													Description:         "Port is an L4 port number. The string will be strictly parsed as a single uint16.",
													MarkdownDescription: "Port is an L4 port number. The string will be strictly parsed as a single uint16.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^()([1-9]|[1-5]?[0-9]{2,4}|6[1-4][0-9]{3}|65[1-4][0-9]{2}|655[1-2][0-9]|6553[1-5])$`), ""),
													},
												},

												"protocol": schema.StringAttribute{
													Description:         "Protocol is the L4 protocol. Accepted values: 'TCP', 'UDP'",
													MarkdownDescription: "Protocol is the L4 protocol. Accepted values: 'TCP', 'UDP'",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("TCP", "UDP"),
													},
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"skip_redirect_from_backend": schema.BoolAttribute{
						Description:         "SkipRedirectFromBackend indicates whether traffic matching RedirectFrontend from RedirectBackend should skip redirection, and hence the traffic will be forwarded as-is. The default is false which means traffic matching RedirectFrontend will get redirected from all pods, including the RedirectBackend(s). Example: If RedirectFrontend is configured to '169.254.169.254:80' as the traffic that needs to be redirected to backends selected by RedirectBackend, if SkipRedirectFromBackend is set to true, traffic going to '169.254.169.254:80' from such backends will not be redirected back to the backends. Instead, the matched traffic from the backends will be forwarded to the original destination '169.254.169.254:80'.",
						MarkdownDescription: "SkipRedirectFromBackend indicates whether traffic matching RedirectFrontend from RedirectBackend should skip redirection, and hence the traffic will be forwarded as-is. The default is false which means traffic matching RedirectFrontend will get redirected from all pods, including the RedirectBackend(s). Example: If RedirectFrontend is configured to '169.254.169.254:80' as the traffic that needs to be redirected to backends selected by RedirectBackend, if SkipRedirectFromBackend is set to true, traffic going to '169.254.169.254:80' from such backends will not be redirected back to the backends. Instead, the matched traffic from the backends will be forwarded to the original destination '169.254.169.254:80'.",
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

func (r *CiliumIoCiliumLocalRedirectPolicyV2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cilium_io_cilium_local_redirect_policy_v2_manifest")

	var model CiliumIoCiliumLocalRedirectPolicyV2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cilium.io/v2")
	model.Kind = pointer.String("CiliumLocalRedirectPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
