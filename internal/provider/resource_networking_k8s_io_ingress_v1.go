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

type NetworkingK8SIoIngressV1Resource struct{}

var (
	_ resource.Resource = (*NetworkingK8SIoIngressV1Resource)(nil)
)

type NetworkingK8SIoIngressV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type NetworkingK8SIoIngressV1GoModel struct {
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
		DefaultBackend *struct {
			Resource *struct {
				ApiGroup *string `tfsdk:"api_group" yaml:"apiGroup,omitempty"`

				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"resource" yaml:"resource,omitempty"`

			Service *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Port *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Number *int64 `tfsdk:"number" yaml:"number,omitempty"`
				} `tfsdk:"port" yaml:"port,omitempty"`
			} `tfsdk:"service" yaml:"service,omitempty"`
		} `tfsdk:"default_backend" yaml:"defaultBackend,omitempty"`

		IngressClassName *string `tfsdk:"ingress_class_name" yaml:"ingressClassName,omitempty"`

		Rules *[]struct {
			Host *string `tfsdk:"host" yaml:"host,omitempty"`

			Http *struct {
				Paths *[]struct {
					Backend *struct {
						Resource *struct {
							ApiGroup *string `tfsdk:"api_group" yaml:"apiGroup,omitempty"`

							Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"resource" yaml:"resource,omitempty"`

						Service *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Port *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Number *int64 `tfsdk:"number" yaml:"number,omitempty"`
							} `tfsdk:"port" yaml:"port,omitempty"`
						} `tfsdk:"service" yaml:"service,omitempty"`
					} `tfsdk:"backend" yaml:"backend,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					PathType *string `tfsdk:"path_type" yaml:"pathType,omitempty"`
				} `tfsdk:"paths" yaml:"paths,omitempty"`
			} `tfsdk:"http" yaml:"http,omitempty"`
		} `tfsdk:"rules" yaml:"rules,omitempty"`

		Tls *[]struct {
			Hosts *[]string `tfsdk:"hosts" yaml:"hosts,omitempty"`

			SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
		} `tfsdk:"tls" yaml:"tls,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewNetworkingK8SIoIngressV1Resource() resource.Resource {
	return &NetworkingK8SIoIngressV1Resource{}
}

func (r *NetworkingK8SIoIngressV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networking_k8s_io_ingress_v1"
}

func (r *NetworkingK8SIoIngressV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend. An Ingress can be configured to give services externally-reachable urls, load balance traffic, terminate SSL, offer name based virtual hosting etc.",
		MarkdownDescription: "Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend. An Ingress can be configured to give services externally-reachable urls, load balance traffic, terminate SSL, offer name based virtual hosting etc.",
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
				Description:         "IngressSpec describes the Ingress the user wishes to exist.",
				MarkdownDescription: "IngressSpec describes the Ingress the user wishes to exist.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"default_backend": {
						Description:         "IngressBackend describes all endpoints for a given service and port.",
						MarkdownDescription: "IngressBackend describes all endpoints for a given service and port.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"resource": {
								Description:         "TypedLocalObjectReference contains enough information to let you locate the typed referenced object inside the same namespace.",
								MarkdownDescription: "TypedLocalObjectReference contains enough information to let you locate the typed referenced object inside the same namespace.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"api_group": {
										Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
										MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kind": {
										Description:         "Kind is the type of resource being referenced",
										MarkdownDescription: "Kind is the type of resource being referenced",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": {
										Description:         "Name is the name of resource being referenced",
										MarkdownDescription: "Name is the name of resource being referenced",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service": {
								Description:         "IngressServiceBackend references a Kubernetes Service as a Backend.",
								MarkdownDescription: "IngressServiceBackend references a Kubernetes Service as a Backend.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name is the referenced service. The service must exist in the same namespace as the Ingress object.",
										MarkdownDescription: "Name is the referenced service. The service must exist in the same namespace as the Ingress object.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"port": {
										Description:         "ServiceBackendPort is the service port being referenced.",
										MarkdownDescription: "ServiceBackendPort is the service port being referenced.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name is the name of the port on the Service. This is a mutually exclusive setting with 'Number'.",
												MarkdownDescription: "Name is the name of the port on the Service. This is a mutually exclusive setting with 'Number'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"number": {
												Description:         "Number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with 'Name'.",
												MarkdownDescription: "Number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with 'Name'.",

												Type: types.Int64Type,

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress_class_name": {
						Description:         "IngressClassName is the name of an IngressClass cluster resource. Ingress controller implementations use this field to know whether they should be serving this Ingress resource, by a transitive connection (controller -> IngressClass -> Ingress resource). Although the 'kubernetes.io/ingress.class' annotation (simple constant name) was never formally defined, it was widely supported by Ingress controllers to create a direct binding between Ingress controller and Ingress resources. Newly created Ingress resources should prefer using the field. However, even though the annotation is officially deprecated, for backwards compatibility reasons, ingress controllers should still honor that annotation if present.",
						MarkdownDescription: "IngressClassName is the name of an IngressClass cluster resource. Ingress controller implementations use this field to know whether they should be serving this Ingress resource, by a transitive connection (controller -> IngressClass -> Ingress resource). Although the 'kubernetes.io/ingress.class' annotation (simple constant name) was never formally defined, it was widely supported by Ingress controllers to create a direct binding between Ingress controller and Ingress resources. Newly created Ingress resources should prefer using the field. However, even though the annotation is officially deprecated, for backwards compatibility reasons, ingress controllers should still honor that annotation if present.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"rules": {
						Description:         "A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.",
						MarkdownDescription: "A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"host": {
								Description:         "Host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the 'host' part of the URI as defined in RFC 3986: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to   the IP in the Spec of the parent Ingress.2. The ':' delimiter is not respected because ports are not allowed.	  Currently the port of an Ingress is implicitly :80 for http and	  :443 for https.Both these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue.Host can be 'precise' which is a domain name without the terminating dot of a network host (e.g. 'foo.bar.com') or 'wildcard', which is a domain name prefixed with a single wildcard label (e.g. '*.foo.com'). The wildcard character '*' must appear by itself as the first DNS label and matches only a single label. You cannot have a wildcard label by itself (e.g. Host == '*'). Requests will be matched against the Host field in the following way: 1. If Host is precise, the request matches this rule if the http host header is equal to Host. 2. If Host is a wildcard, then the request matches this rule if the http host header is to equal to the suffix (removing the first label) of the wildcard rule.",
								MarkdownDescription: "Host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the 'host' part of the URI as defined in RFC 3986: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to   the IP in the Spec of the parent Ingress.2. The ':' delimiter is not respected because ports are not allowed.	  Currently the port of an Ingress is implicitly :80 for http and	  :443 for https.Both these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue.Host can be 'precise' which is a domain name without the terminating dot of a network host (e.g. 'foo.bar.com') or 'wildcard', which is a domain name prefixed with a single wildcard label (e.g. '*.foo.com'). The wildcard character '*' must appear by itself as the first DNS label and matches only a single label. You cannot have a wildcard label by itself (e.g. Host == '*'). Requests will be matched against the Host field in the following way: 1. If Host is precise, the request matches this rule if the http host header is equal to Host. 2. If Host is a wildcard, then the request matches this rule if the http host header is to equal to the suffix (removing the first label) of the wildcard rule.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"http": {
								Description:         "HTTPIngressRuleValue is a list of http selectors pointing to backends. In the example: http://<host>/<path>?<searchpart> -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'.",
								MarkdownDescription: "HTTPIngressRuleValue is a list of http selectors pointing to backends. In the example: http://<host>/<path>?<searchpart> -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"paths": {
										Description:         "A collection of paths that map requests to backends.",
										MarkdownDescription: "A collection of paths that map requests to backends.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"backend": {
												Description:         "IngressBackend describes all endpoints for a given service and port.",
												MarkdownDescription: "IngressBackend describes all endpoints for a given service and port.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"resource": {
														Description:         "TypedLocalObjectReference contains enough information to let you locate the typed referenced object inside the same namespace.",
														MarkdownDescription: "TypedLocalObjectReference contains enough information to let you locate the typed referenced object inside the same namespace.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"api_group": {
																Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kind": {
																Description:         "Kind is the type of resource being referenced",
																MarkdownDescription: "Kind is the type of resource being referenced",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
																Description:         "Name is the name of resource being referenced",
																MarkdownDescription: "Name is the name of resource being referenced",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"service": {
														Description:         "IngressServiceBackend references a Kubernetes Service as a Backend.",
														MarkdownDescription: "IngressServiceBackend references a Kubernetes Service as a Backend.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name is the referenced service. The service must exist in the same namespace as the Ingress object.",
																MarkdownDescription: "Name is the referenced service. The service must exist in the same namespace as the Ingress object.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"port": {
																Description:         "ServiceBackendPort is the service port being referenced.",
																MarkdownDescription: "ServiceBackendPort is the service port being referenced.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "Name is the name of the port on the Service. This is a mutually exclusive setting with 'Number'.",
																		MarkdownDescription: "Name is the name of the port on the Service. This is a mutually exclusive setting with 'Number'.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"number": {
																		Description:         "Number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with 'Name'.",
																		MarkdownDescription: "Number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with 'Name'.",

																		Type: types.Int64Type,

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
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"path": {
												Description:         "Path is matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986. Paths must begin with a '/' and must be present when using PathType with value 'Exact' or 'Prefix'.",
												MarkdownDescription: "Path is matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986. Paths must begin with a '/' and must be present when using PathType with value 'Exact' or 'Prefix'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"path_type": {
												Description:         "PathType determines the interpretation of the Path matching. PathType can be one of the following values: * Exact: Matches the URL path exactly. * Prefix: Matches based on a URL path prefix split by '/'. Matching is  done on a path element by element basis. A path element refers is the  list of labels in the path split by the '/' separator. A request is a  match for path p if every p is an element-wise prefix of p of the  request path. Note that if the last element of the path is a substring  of the last element in request path, it is not a match (e.g. /foo/bar  matches /foo/bar/baz, but does not match /foo/barbaz).* ImplementationSpecific: Interpretation of the Path matching is up to  the IngressClass. Implementations can treat this as a separate PathType  or treat it identically to Prefix or Exact path types.Implementations are required to support all path types.",
												MarkdownDescription: "PathType determines the interpretation of the Path matching. PathType can be one of the following values: * Exact: Matches the URL path exactly. * Prefix: Matches based on a URL path prefix split by '/'. Matching is  done on a path element by element basis. A path element refers is the  list of labels in the path split by the '/' separator. A request is a  match for path p if every p is an element-wise prefix of p of the  request path. Note that if the last element of the path is a substring  of the last element in request path, it is not a match (e.g. /foo/bar  matches /foo/bar/baz, but does not match /foo/barbaz).* ImplementationSpecific: Interpretation of the Path matching is up to  the IngressClass. Implementations can treat this as a separate PathType  or treat it identically to Prefix or Exact path types.Implementations are required to support all path types.",

												Type: types.StringType,

												Required: true,
												Optional: false,
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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tls": {
						Description:         "TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.",
						MarkdownDescription: "TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"hosts": {
								Description:         "Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
								MarkdownDescription: "Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"secret_name": {
								Description:         "SecretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.",
								MarkdownDescription: "SecretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.",

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
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *NetworkingK8SIoIngressV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_networking_k8s_io_ingress_v1")

	var state NetworkingK8SIoIngressV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel NetworkingK8SIoIngressV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("networking.k8s.io/v1")
	goModel.Kind = utilities.Ptr("Ingress")

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

func (r *NetworkingK8SIoIngressV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networking_k8s_io_ingress_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *NetworkingK8SIoIngressV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_networking_k8s_io_ingress_v1")

	var state NetworkingK8SIoIngressV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel NetworkingK8SIoIngressV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("networking.k8s.io/v1")
	goModel.Kind = utilities.Ptr("Ingress")

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

func (r *NetworkingK8SIoIngressV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_networking_k8s_io_ingress_v1")
	// NO-OP: Terraform removes the state automatically for us
}
