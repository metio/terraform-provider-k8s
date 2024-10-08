/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package projectcontour_io_v1alpha1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &ProjectcontourIoExtensionServiceV1Alpha1Manifest{}
)

func NewProjectcontourIoExtensionServiceV1Alpha1Manifest() datasource.DataSource {
	return &ProjectcontourIoExtensionServiceV1Alpha1Manifest{}
}

type ProjectcontourIoExtensionServiceV1Alpha1Manifest struct{}

type ProjectcontourIoExtensionServiceV1Alpha1ManifestData struct {
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
		CircuitBreakerPolicy *struct {
			MaxConnections        *int64 `tfsdk:"max_connections" json:"maxConnections,omitempty"`
			MaxPendingRequests    *int64 `tfsdk:"max_pending_requests" json:"maxPendingRequests,omitempty"`
			MaxRequests           *int64 `tfsdk:"max_requests" json:"maxRequests,omitempty"`
			MaxRetries            *int64 `tfsdk:"max_retries" json:"maxRetries,omitempty"`
			PerHostMaxConnections *int64 `tfsdk:"per_host_max_connections" json:"perHostMaxConnections,omitempty"`
		} `tfsdk:"circuit_breaker_policy" json:"circuitBreakerPolicy,omitempty"`
		LoadBalancerPolicy *struct {
			RequestHashPolicies *[]struct {
				HashSourceIP      *bool `tfsdk:"hash_source_ip" json:"hashSourceIP,omitempty"`
				HeaderHashOptions *struct {
					HeaderName *string `tfsdk:"header_name" json:"headerName,omitempty"`
				} `tfsdk:"header_hash_options" json:"headerHashOptions,omitempty"`
				QueryParameterHashOptions *struct {
					ParameterName *string `tfsdk:"parameter_name" json:"parameterName,omitempty"`
				} `tfsdk:"query_parameter_hash_options" json:"queryParameterHashOptions,omitempty"`
				Terminal *bool `tfsdk:"terminal" json:"terminal,omitempty"`
			} `tfsdk:"request_hash_policies" json:"requestHashPolicies,omitempty"`
			Strategy *string `tfsdk:"strategy" json:"strategy,omitempty"`
		} `tfsdk:"load_balancer_policy" json:"loadBalancerPolicy,omitempty"`
		Protocol        *string `tfsdk:"protocol" json:"protocol,omitempty"`
		ProtocolVersion *string `tfsdk:"protocol_version" json:"protocolVersion,omitempty"`
		Services        *[]struct {
			Name   *string `tfsdk:"name" json:"name,omitempty"`
			Port   *int64  `tfsdk:"port" json:"port,omitempty"`
			Weight *int64  `tfsdk:"weight" json:"weight,omitempty"`
		} `tfsdk:"services" json:"services,omitempty"`
		TimeoutPolicy *struct {
			Idle           *string `tfsdk:"idle" json:"idle,omitempty"`
			IdleConnection *string `tfsdk:"idle_connection" json:"idleConnection,omitempty"`
			Response       *string `tfsdk:"response" json:"response,omitempty"`
		} `tfsdk:"timeout_policy" json:"timeoutPolicy,omitempty"`
		Validation *struct {
			CaSecret     *string   `tfsdk:"ca_secret" json:"caSecret,omitempty"`
			SubjectName  *string   `tfsdk:"subject_name" json:"subjectName,omitempty"`
			SubjectNames *[]string `tfsdk:"subject_names" json:"subjectNames,omitempty"`
		} `tfsdk:"validation" json:"validation,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ProjectcontourIoExtensionServiceV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_projectcontour_io_extension_service_v1alpha1_manifest"
}

func (r *ProjectcontourIoExtensionServiceV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ExtensionService is the schema for the Contour extension services API. An ExtensionService resource binds a network service to the Contour API so that Contour API features can be implemented by collaborating components.",
		MarkdownDescription: "ExtensionService is the schema for the Contour extension services API. An ExtensionService resource binds a network service to the Contour API so that Contour API features can be implemented by collaborating components.",
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
				Description:         "ExtensionServiceSpec defines the desired state of an ExtensionService resource.",
				MarkdownDescription: "ExtensionServiceSpec defines the desired state of an ExtensionService resource.",
				Attributes: map[string]schema.Attribute{
					"circuit_breaker_policy": schema.SingleNestedAttribute{
						Description:         "CircuitBreakerPolicy specifies the circuit breaker budget across the extension service. If defined this overrides the global circuit breaker budget.",
						MarkdownDescription: "CircuitBreakerPolicy specifies the circuit breaker budget across the extension service. If defined this overrides the global circuit breaker budget.",
						Attributes: map[string]schema.Attribute{
							"max_connections": schema.Int64Attribute{
								Description:         "The maximum number of connections that a single Envoy instance allows to the Kubernetes Service; defaults to 1024.",
								MarkdownDescription: "The maximum number of connections that a single Envoy instance allows to the Kubernetes Service; defaults to 1024.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_pending_requests": schema.Int64Attribute{
								Description:         "The maximum number of pending requests that a single Envoy instance allows to the Kubernetes Service; defaults to 1024.",
								MarkdownDescription: "The maximum number of pending requests that a single Envoy instance allows to the Kubernetes Service; defaults to 1024.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_requests": schema.Int64Attribute{
								Description:         "The maximum parallel requests a single Envoy instance allows to the Kubernetes Service; defaults to 1024",
								MarkdownDescription: "The maximum parallel requests a single Envoy instance allows to the Kubernetes Service; defaults to 1024",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_retries": schema.Int64Attribute{
								Description:         "The maximum number of parallel retries a single Envoy instance allows to the Kubernetes Service; defaults to 3.",
								MarkdownDescription: "The maximum number of parallel retries a single Envoy instance allows to the Kubernetes Service; defaults to 3.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"per_host_max_connections": schema.Int64Attribute{
								Description:         "PerHostMaxConnections is the maximum number of connections that Envoy will allow to each individual host in a cluster.",
								MarkdownDescription: "PerHostMaxConnections is the maximum number of connections that Envoy will allow to each individual host in a cluster.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"load_balancer_policy": schema.SingleNestedAttribute{
						Description:         "The policy for load balancing GRPC service requests. Note that the 'Cookie' and 'RequestHash' load balancing strategies cannot be used here.",
						MarkdownDescription: "The policy for load balancing GRPC service requests. Note that the 'Cookie' and 'RequestHash' load balancing strategies cannot be used here.",
						Attributes: map[string]schema.Attribute{
							"request_hash_policies": schema.ListNestedAttribute{
								Description:         "RequestHashPolicies contains a list of hash policies to apply when the 'RequestHash' load balancing strategy is chosen. If an element of the supplied list of hash policies is invalid, it will be ignored. If the list of hash policies is empty after validation, the load balancing strategy will fall back to the default 'RoundRobin'.",
								MarkdownDescription: "RequestHashPolicies contains a list of hash policies to apply when the 'RequestHash' load balancing strategy is chosen. If an element of the supplied list of hash policies is invalid, it will be ignored. If the list of hash policies is empty after validation, the load balancing strategy will fall back to the default 'RoundRobin'.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"hash_source_ip": schema.BoolAttribute{
											Description:         "HashSourceIP should be set to true when request source IP hash based load balancing is desired. It must be the only hash option field set, otherwise this request hash policy object will be ignored.",
											MarkdownDescription: "HashSourceIP should be set to true when request source IP hash based load balancing is desired. It must be the only hash option field set, otherwise this request hash policy object will be ignored.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"header_hash_options": schema.SingleNestedAttribute{
											Description:         "HeaderHashOptions should be set when request header hash based load balancing is desired. It must be the only hash option field set, otherwise this request hash policy object will be ignored.",
											MarkdownDescription: "HeaderHashOptions should be set when request header hash based load balancing is desired. It must be the only hash option field set, otherwise this request hash policy object will be ignored.",
											Attributes: map[string]schema.Attribute{
												"header_name": schema.StringAttribute{
													Description:         "HeaderName is the name of the HTTP request header that will be used to calculate the hash key. If the header specified is not present on a request, no hash will be produced.",
													MarkdownDescription: "HeaderName is the name of the HTTP request header that will be used to calculate the hash key. If the header specified is not present on a request, no hash will be produced.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"query_parameter_hash_options": schema.SingleNestedAttribute{
											Description:         "QueryParameterHashOptions should be set when request query parameter hash based load balancing is desired. It must be the only hash option field set, otherwise this request hash policy object will be ignored.",
											MarkdownDescription: "QueryParameterHashOptions should be set when request query parameter hash based load balancing is desired. It must be the only hash option field set, otherwise this request hash policy object will be ignored.",
											Attributes: map[string]schema.Attribute{
												"parameter_name": schema.StringAttribute{
													Description:         "ParameterName is the name of the HTTP request query parameter that will be used to calculate the hash key. If the query parameter specified is not present on a request, no hash will be produced.",
													MarkdownDescription: "ParameterName is the name of the HTTP request query parameter that will be used to calculate the hash key. If the query parameter specified is not present on a request, no hash will be produced.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"terminal": schema.BoolAttribute{
											Description:         "Terminal is a flag that allows for short-circuiting computing of a hash for a given request. If set to true, and the request attribute specified in the attribute hash options is present, no further hash policies will be used to calculate a hash for the request.",
											MarkdownDescription: "Terminal is a flag that allows for short-circuiting computing of a hash for a given request. If set to true, and the request attribute specified in the attribute hash options is present, no further hash policies will be used to calculate a hash for the request.",
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

							"strategy": schema.StringAttribute{
								Description:         "Strategy specifies the policy used to balance requests across the pool of backend pods. Valid policy names are 'Random', 'RoundRobin', 'WeightedLeastRequest', 'Cookie', and 'RequestHash'. If an unknown strategy name is specified or no policy is supplied, the default 'RoundRobin' policy is used.",
								MarkdownDescription: "Strategy specifies the policy used to balance requests across the pool of backend pods. Valid policy names are 'Random', 'RoundRobin', 'WeightedLeastRequest', 'Cookie', and 'RequestHash'. If an unknown strategy name is specified or no policy is supplied, the default 'RoundRobin' policy is used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"protocol": schema.StringAttribute{
						Description:         "Protocol may be used to specify (or override) the protocol used to reach this Service. Values may be h2 or h2c. If omitted, protocol-selection falls back on Service annotations.",
						MarkdownDescription: "Protocol may be used to specify (or override) the protocol used to reach this Service. Values may be h2 or h2c. If omitted, protocol-selection falls back on Service annotations.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("h2", "h2c"),
						},
					},

					"protocol_version": schema.StringAttribute{
						Description:         "This field sets the version of the GRPC protocol that Envoy uses to send requests to the extension service. Since Contour always uses the v3 Envoy API, this is currently fixed at 'v3'. However, other protocol options will be available in future.",
						MarkdownDescription: "This field sets the version of the GRPC protocol that Envoy uses to send requests to the extension service. Since Contour always uses the v3 Envoy API, this is currently fixed at 'v3'. However, other protocol options will be available in future.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("v3"),
						},
					},

					"services": schema.ListNestedAttribute{
						Description:         "Services specifies the set of Kubernetes Service resources that receive GRPC extension API requests. If no weights are specified for any of the entries in this array, traffic will be spread evenly across all the services. Otherwise, traffic is balanced proportionally to the Weight field in each entry.",
						MarkdownDescription: "Services specifies the set of Kubernetes Service resources that receive GRPC extension API requests. If no weights are specified for any of the entries in this array, traffic will be spread evenly across all the services. Otherwise, traffic is balanced proportionally to the Weight field in each entry.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name is the name of Kubernetes service that will accept service traffic.",
									MarkdownDescription: "Name is the name of Kubernetes service that will accept service traffic.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"port": schema.Int64Attribute{
									Description:         "Port (defined as Integer) to proxy traffic to since a service can have multiple defined.",
									MarkdownDescription: "Port (defined as Integer) to proxy traffic to since a service can have multiple defined.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(1),
										int64validator.AtMost(65535),
									},
								},

								"weight": schema.Int64Attribute{
									Description:         "Weight defines proportion of traffic to balance to the Kubernetes Service.",
									MarkdownDescription: "Weight defines proportion of traffic to balance to the Kubernetes Service.",
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

					"timeout_policy": schema.SingleNestedAttribute{
						Description:         "The timeout policy for requests to the services.",
						MarkdownDescription: "The timeout policy for requests to the services.",
						Attributes: map[string]schema.Attribute{
							"idle": schema.StringAttribute{
								Description:         "Timeout for how long the proxy should wait while there is no activity during single request/response (for HTTP/1.1) or stream (for HTTP/2). Timeout will not trigger while HTTP/1.1 connection is idle between two consecutive requests. If not specified, there is no per-route idle timeout, though a connection manager-wide stream_idle_timeout default of 5m still applies.",
								MarkdownDescription: "Timeout for how long the proxy should wait while there is no activity during single request/response (for HTTP/1.1) or stream (for HTTP/2). Timeout will not trigger while HTTP/1.1 connection is idle between two consecutive requests. If not specified, there is no per-route idle timeout, though a connection manager-wide stream_idle_timeout default of 5m still applies.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+|infinity|infinite)$`), ""),
								},
							},

							"idle_connection": schema.StringAttribute{
								Description:         "Timeout for how long connection from the proxy to the upstream service is kept when there are no active requests. If not supplied, Envoy's default value of 1h applies.",
								MarkdownDescription: "Timeout for how long connection from the proxy to the upstream service is kept when there are no active requests. If not supplied, Envoy's default value of 1h applies.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+|infinity|infinite)$`), ""),
								},
							},

							"response": schema.StringAttribute{
								Description:         "Timeout for receiving a response from the server after processing a request from client. If not supplied, Envoy's default value of 15s applies.",
								MarkdownDescription: "Timeout for receiving a response from the server after processing a request from client. If not supplied, Envoy's default value of 15s applies.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+|infinity|infinite)$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"validation": schema.SingleNestedAttribute{
						Description:         "UpstreamValidation defines how to verify the backend service's certificate",
						MarkdownDescription: "UpstreamValidation defines how to verify the backend service's certificate",
						Attributes: map[string]schema.Attribute{
							"ca_secret": schema.StringAttribute{
								Description:         "Name or namespaced name of the Kubernetes secret used to validate the certificate presented by the backend. The secret must contain key named ca.crt. The name can be optionally prefixed with namespace 'namespace/name'. When cross-namespace reference is used, TLSCertificateDelegation resource must exist in the namespace to grant access to the secret. Max length should be the actual max possible length of a namespaced name (63 + 253 + 1 = 317)",
								MarkdownDescription: "Name or namespaced name of the Kubernetes secret used to validate the certificate presented by the backend. The secret must contain key named ca.crt. The name can be optionally prefixed with namespace 'namespace/name'. When cross-namespace reference is used, TLSCertificateDelegation resource must exist in the namespace to grant access to the secret. Max length should be the actual max possible length of a namespaced name (63 + 253 + 1 = 317)",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(317),
								},
							},

							"subject_name": schema.StringAttribute{
								Description:         "Key which is expected to be present in the 'subjectAltName' of the presented certificate. Deprecated: migrate to using the plural field subjectNames.",
								MarkdownDescription: "Key which is expected to be present in the 'subjectAltName' of the presented certificate. Deprecated: migrate to using the plural field subjectNames.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(250),
								},
							},

							"subject_names": schema.ListAttribute{
								Description:         "List of keys, of which at least one is expected to be present in the 'subjectAltName of the presented certificate.",
								MarkdownDescription: "List of keys, of which at least one is expected to be present in the 'subjectAltName of the presented certificate.",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *ProjectcontourIoExtensionServiceV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_projectcontour_io_extension_service_v1alpha1_manifest")

	var model ProjectcontourIoExtensionServiceV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("projectcontour.io/v1alpha1")
	model.Kind = pointer.String("ExtensionService")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
