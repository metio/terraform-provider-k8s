/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package networking_karmada_io_v1alpha1

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
	_ datasource.DataSource = &NetworkingKarmadaIoMultiClusterIngressV1Alpha1Manifest{}
)

func NewNetworkingKarmadaIoMultiClusterIngressV1Alpha1Manifest() datasource.DataSource {
	return &NetworkingKarmadaIoMultiClusterIngressV1Alpha1Manifest{}
}

type NetworkingKarmadaIoMultiClusterIngressV1Alpha1Manifest struct{}

type NetworkingKarmadaIoMultiClusterIngressV1Alpha1ManifestData struct {
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
		DefaultBackend *struct {
			Resource *struct {
				ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
				Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"resource" json:"resource,omitempty"`
			Service *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
				Port *struct {
					Name   *string `tfsdk:"name" json:"name,omitempty"`
					Number *int64  `tfsdk:"number" json:"number,omitempty"`
				} `tfsdk:"port" json:"port,omitempty"`
			} `tfsdk:"service" json:"service,omitempty"`
		} `tfsdk:"default_backend" json:"defaultBackend,omitempty"`
		IngressClassName *string `tfsdk:"ingress_class_name" json:"ingressClassName,omitempty"`
		Rules            *[]struct {
			Host *string `tfsdk:"host" json:"host,omitempty"`
			Http *struct {
				Paths *[]struct {
					Backend *struct {
						Resource *struct {
							ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
							Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"resource" json:"resource,omitempty"`
						Service *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
							Port *struct {
								Name   *string `tfsdk:"name" json:"name,omitempty"`
								Number *int64  `tfsdk:"number" json:"number,omitempty"`
							} `tfsdk:"port" json:"port,omitempty"`
						} `tfsdk:"service" json:"service,omitempty"`
					} `tfsdk:"backend" json:"backend,omitempty"`
					Path     *string `tfsdk:"path" json:"path,omitempty"`
					PathType *string `tfsdk:"path_type" json:"pathType,omitempty"`
				} `tfsdk:"paths" json:"paths,omitempty"`
			} `tfsdk:"http" json:"http,omitempty"`
		} `tfsdk:"rules" json:"rules,omitempty"`
		Tls *[]struct {
			Hosts      *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
			SecretName *string   `tfsdk:"secret_name" json:"secretName,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NetworkingKarmadaIoMultiClusterIngressV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_networking_karmada_io_multi_cluster_ingress_v1alpha1_manifest"
}

func (r *NetworkingKarmadaIoMultiClusterIngressV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "MultiClusterIngress is a collection of rules that allow inbound connections to reach theendpoints defined by a backend. The structure of MultiClusterIngress is same as Ingress,indicates the Ingress in multi-clusters.",
		MarkdownDescription: "MultiClusterIngress is a collection of rules that allow inbound connections to reach theendpoints defined by a backend. The structure of MultiClusterIngress is same as Ingress,indicates the Ingress in multi-clusters.",
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
				Description:         "Spec is the desired state of the MultiClusterIngress.",
				MarkdownDescription: "Spec is the desired state of the MultiClusterIngress.",
				Attributes: map[string]schema.Attribute{
					"default_backend": schema.SingleNestedAttribute{
						Description:         "defaultBackend is the backend that should handle requests that don'tmatch any rule. If Rules are not specified, DefaultBackend must be specified.If DefaultBackend is not set, the handling of requests that do not match anyof the rules will be up to the Ingress controller.",
						MarkdownDescription: "defaultBackend is the backend that should handle requests that don'tmatch any rule. If Rules are not specified, DefaultBackend must be specified.If DefaultBackend is not set, the handling of requests that do not match anyof the rules will be up to the Ingress controller.",
						Attributes: map[string]schema.Attribute{
							"resource": schema.SingleNestedAttribute{
								Description:         "resource is an ObjectRef to another Kubernetes resource in the namespaceof the Ingress object. If resource is specified, a service.Name andservice.Port must not be specified.This is a mutually exclusive setting with 'Service'.",
								MarkdownDescription: "resource is an ObjectRef to another Kubernetes resource in the namespaceof the Ingress object. If resource is specified, a service.Name andservice.Port must not be specified.This is a mutually exclusive setting with 'Service'.",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"service": schema.SingleNestedAttribute{
								Description:         "service references a service as a backend.This is a mutually exclusive setting with 'Resource'.",
								MarkdownDescription: "service references a service as a backend.This is a mutually exclusive setting with 'Resource'.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "name is the referenced service. The service must exist inthe same namespace as the Ingress object.",
										MarkdownDescription: "name is the referenced service. The service must exist inthe same namespace as the Ingress object.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"port": schema.SingleNestedAttribute{
										Description:         "port of the referenced service. A port name or port numberis required for a IngressServiceBackend.",
										MarkdownDescription: "port of the referenced service. A port name or port numberis required for a IngressServiceBackend.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "name is the name of the port on the Service.This is a mutually exclusive setting with 'Number'.",
												MarkdownDescription: "name is the name of the port on the Service.This is a mutually exclusive setting with 'Number'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"number": schema.Int64Attribute{
												Description:         "number is the numerical port number (e.g. 80) on the Service.This is a mutually exclusive setting with 'Name'.",
												MarkdownDescription: "number is the numerical port number (e.g. 80) on the Service.This is a mutually exclusive setting with 'Name'.",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress_class_name": schema.StringAttribute{
						Description:         "ingressClassName is the name of an IngressClass cluster resource. Ingresscontroller implementations use this field to know whether they should beserving this Ingress resource, by a transitive connection(controller -> IngressClass -> Ingress resource). Although the'kubernetes.io/ingress.class' annotation (simple constant name) was neverformally defined, it was widely supported by Ingress controllers to createa direct binding between Ingress controller and Ingress resources. Newlycreated Ingress resources should prefer using the field. However, eventhough the annotation is officially deprecated, for backwards compatibilityreasons, ingress controllers should still honor that annotation if present.",
						MarkdownDescription: "ingressClassName is the name of an IngressClass cluster resource. Ingresscontroller implementations use this field to know whether they should beserving this Ingress resource, by a transitive connection(controller -> IngressClass -> Ingress resource). Although the'kubernetes.io/ingress.class' annotation (simple constant name) was neverformally defined, it was widely supported by Ingress controllers to createa direct binding between Ingress controller and Ingress resources. Newlycreated Ingress resources should prefer using the field. However, eventhough the annotation is officially deprecated, for backwards compatibilityreasons, ingress controllers should still honor that annotation if present.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"rules": schema.ListNestedAttribute{
						Description:         "rules is a list of host rules used to configure the Ingress. If unspecified,or no rule matches, all traffic is sent to the default backend.",
						MarkdownDescription: "rules is a list of host rules used to configure the Ingress. If unspecified,or no rule matches, all traffic is sent to the default backend.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"host": schema.StringAttribute{
									Description:         "host is the fully qualified domain name of a network host, as defined by RFC 3986.Note the following deviations from the 'host' part of theURI as defined in RFC 3986:1. IPs are not allowed. Currently an IngressRuleValue can only apply to   the IP in the Spec of the parent Ingress.2. The ':' delimiter is not respected because ports are not allowed.	  Currently the port of an Ingress is implicitly :80 for http and	  :443 for https.Both these may change in the future.Incoming requests are matched against the host before theIngressRuleValue. If the host is unspecified, the Ingress routes alltraffic based on the specified IngressRuleValue.host can be 'precise' which is a domain name without the terminating dot ofa network host (e.g. 'foo.bar.com') or 'wildcard', which is a domain nameprefixed with a single wildcard label (e.g. '*.foo.com').The wildcard character '*' must appear by itself as the first DNS label andmatches only a single label. You cannot have a wildcard label by itself (e.g. Host == '*').Requests will be matched against the Host field in the following way:1. If host is precise, the request matches this rule if the http host header is equal to Host.2. If host is a wildcard, then the request matches this rule if the http host headeris to equal to the suffix (removing the first label) of the wildcard rule.",
									MarkdownDescription: "host is the fully qualified domain name of a network host, as defined by RFC 3986.Note the following deviations from the 'host' part of theURI as defined in RFC 3986:1. IPs are not allowed. Currently an IngressRuleValue can only apply to   the IP in the Spec of the parent Ingress.2. The ':' delimiter is not respected because ports are not allowed.	  Currently the port of an Ingress is implicitly :80 for http and	  :443 for https.Both these may change in the future.Incoming requests are matched against the host before theIngressRuleValue. If the host is unspecified, the Ingress routes alltraffic based on the specified IngressRuleValue.host can be 'precise' which is a domain name without the terminating dot ofa network host (e.g. 'foo.bar.com') or 'wildcard', which is a domain nameprefixed with a single wildcard label (e.g. '*.foo.com').The wildcard character '*' must appear by itself as the first DNS label andmatches only a single label. You cannot have a wildcard label by itself (e.g. Host == '*').Requests will be matched against the Host field in the following way:1. If host is precise, the request matches this rule if the http host header is equal to Host.2. If host is a wildcard, then the request matches this rule if the http host headeris to equal to the suffix (removing the first label) of the wildcard rule.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"http": schema.SingleNestedAttribute{
									Description:         "HTTPIngressRuleValue is a list of http selectors pointing to backends.In the example: http://<host>/<path>?<searchpart> -> backend wherewhere parts of the url correspond to RFC 3986, this resource will be usedto match against everything after the last '/' and before the first '?'or '#'.",
									MarkdownDescription: "HTTPIngressRuleValue is a list of http selectors pointing to backends.In the example: http://<host>/<path>?<searchpart> -> backend wherewhere parts of the url correspond to RFC 3986, this resource will be usedto match against everything after the last '/' and before the first '?'or '#'.",
									Attributes: map[string]schema.Attribute{
										"paths": schema.ListNestedAttribute{
											Description:         "paths is a collection of paths that map requests to backends.",
											MarkdownDescription: "paths is a collection of paths that map requests to backends.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"backend": schema.SingleNestedAttribute{
														Description:         "backend defines the referenced service endpoint to which the trafficwill be forwarded to.",
														MarkdownDescription: "backend defines the referenced service endpoint to which the trafficwill be forwarded to.",
														Attributes: map[string]schema.Attribute{
															"resource": schema.SingleNestedAttribute{
																Description:         "resource is an ObjectRef to another Kubernetes resource in the namespaceof the Ingress object. If resource is specified, a service.Name andservice.Port must not be specified.This is a mutually exclusive setting with 'Service'.",
																MarkdownDescription: "resource is an ObjectRef to another Kubernetes resource in the namespaceof the Ingress object. If resource is specified, a service.Name andservice.Port must not be specified.This is a mutually exclusive setting with 'Service'.",
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
																Required: false,
																Optional: true,
																Computed: false,
															},

															"service": schema.SingleNestedAttribute{
																Description:         "service references a service as a backend.This is a mutually exclusive setting with 'Resource'.",
																MarkdownDescription: "service references a service as a backend.This is a mutually exclusive setting with 'Resource'.",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "name is the referenced service. The service must exist inthe same namespace as the Ingress object.",
																		MarkdownDescription: "name is the referenced service. The service must exist inthe same namespace as the Ingress object.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"port": schema.SingleNestedAttribute{
																		Description:         "port of the referenced service. A port name or port numberis required for a IngressServiceBackend.",
																		MarkdownDescription: "port of the referenced service. A port name or port numberis required for a IngressServiceBackend.",
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "name is the name of the port on the Service.This is a mutually exclusive setting with 'Number'.",
																				MarkdownDescription: "name is the name of the port on the Service.This is a mutually exclusive setting with 'Number'.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"number": schema.Int64Attribute{
																				Description:         "number is the numerical port number (e.g. 80) on the Service.This is a mutually exclusive setting with 'Name'.",
																				MarkdownDescription: "number is the numerical port number (e.g. 80) on the Service.This is a mutually exclusive setting with 'Name'.",
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
														Required: true,
														Optional: false,
														Computed: false,
													},

													"path": schema.StringAttribute{
														Description:         "path is matched against the path of an incoming request. Currently it cancontain characters disallowed from the conventional 'path' part of a URLas defined by RFC 3986. Paths must begin with a '/' and must be presentwhen using PathType with value 'Exact' or 'Prefix'.",
														MarkdownDescription: "path is matched against the path of an incoming request. Currently it cancontain characters disallowed from the conventional 'path' part of a URLas defined by RFC 3986. Paths must begin with a '/' and must be presentwhen using PathType with value 'Exact' or 'Prefix'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"path_type": schema.StringAttribute{
														Description:         "pathType determines the interpretation of the path matching. PathType canbe one of the following values:* Exact: Matches the URL path exactly.* Prefix: Matches based on a URL path prefix split by '/'. Matching is  done on a path element by element basis. A path element refers is the  list of labels in the path split by the '/' separator. A request is a  match for path p if every p is an element-wise prefix of p of the  request path. Note that if the last element of the path is a substring  of the last element in request path, it is not a match (e.g. /foo/bar  matches /foo/bar/baz, but does not match /foo/barbaz).* ImplementationSpecific: Interpretation of the Path matching is up to  the IngressClass. Implementations can treat this as a separate PathType  or treat it identically to Prefix or Exact path types.Implementations are required to support all path types.",
														MarkdownDescription: "pathType determines the interpretation of the path matching. PathType canbe one of the following values:* Exact: Matches the URL path exactly.* Prefix: Matches based on a URL path prefix split by '/'. Matching is  done on a path element by element basis. A path element refers is the  list of labels in the path split by the '/' separator. A request is a  match for path p if every p is an element-wise prefix of p of the  request path. Note that if the last element of the path is a substring  of the last element in request path, it is not a match (e.g. /foo/bar  matches /foo/bar/baz, but does not match /foo/barbaz).* ImplementationSpecific: Interpretation of the Path matching is up to  the IngressClass. Implementations can treat this as a separate PathType  or treat it identically to Prefix or Exact path types.Implementations are required to support all path types.",
														Required:            true,
														Optional:            false,
														Computed:            false,
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tls": schema.ListNestedAttribute{
						Description:         "tls represents the TLS configuration. Currently the Ingress only supports asingle TLS port, 443. If multiple members of this list specify different hosts,they will be multiplexed on the same port according to the hostname specifiedthrough the SNI TLS extension, if the ingress controller fulfilling theingress supports SNI.",
						MarkdownDescription: "tls represents the TLS configuration. Currently the Ingress only supports asingle TLS port, 443. If multiple members of this list specify different hosts,they will be multiplexed on the same port according to the hostname specifiedthrough the SNI TLS extension, if the ingress controller fulfilling theingress supports SNI.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"hosts": schema.ListAttribute{
									Description:         "hosts is a list of hosts included in the TLS certificate. The values inthis list must match the name/s used in the tlsSecret. Defaults to thewildcard host setting for the loadbalancer controller fulfilling thisIngress, if left unspecified.",
									MarkdownDescription: "hosts is a list of hosts included in the TLS certificate. The values inthis list must match the name/s used in the tlsSecret. Defaults to thewildcard host setting for the loadbalancer controller fulfilling thisIngress, if left unspecified.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"secret_name": schema.StringAttribute{
									Description:         "secretName is the name of the secret used to terminate TLS traffic onport 443. Field is left optional to allow TLS routing based on SNIhostname alone. If the SNI host in a listener conflicts with the 'Host'header field used by an IngressRule, the SNI host is used for terminationand value of the 'Host' header is used for routing.",
									MarkdownDescription: "secretName is the name of the secret used to terminate TLS traffic onport 443. Field is left optional to allow TLS routing based on SNIhostname alone. If the SNI host in a listener conflicts with the 'Host'header field used by an IngressRule, the SNI host is used for terminationand value of the 'Host' header is used for routing.",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *NetworkingKarmadaIoMultiClusterIngressV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networking_karmada_io_multi_cluster_ingress_v1alpha1_manifest")

	var model NetworkingKarmadaIoMultiClusterIngressV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("networking.karmada.io/v1alpha1")
	model.Kind = pointer.String("MultiClusterIngress")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
