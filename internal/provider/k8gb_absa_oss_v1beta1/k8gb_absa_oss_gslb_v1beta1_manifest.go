/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package k8gb_absa_oss_v1beta1

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
	_ datasource.DataSource = &K8GbAbsaOssGslbV1Beta1Manifest{}
)

func NewK8GbAbsaOssGslbV1Beta1Manifest() datasource.DataSource {
	return &K8GbAbsaOssGslbV1Beta1Manifest{}
}

type K8GbAbsaOssGslbV1Beta1Manifest struct{}

type K8GbAbsaOssGslbV1Beta1ManifestData struct {
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
		Ingress *struct {
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
		} `tfsdk:"ingress" json:"ingress,omitempty"`
		ResourceRef *struct {
			ApiVersion       *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind             *string `tfsdk:"kind" json:"kind,omitempty"`
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"resource_ref" json:"resourceRef,omitempty"`
		Strategy *struct {
			DnsTtlSeconds              *int64             `tfsdk:"dns_ttl_seconds" json:"dnsTtlSeconds,omitempty"`
			PrimaryGeoTag              *string            `tfsdk:"primary_geo_tag" json:"primaryGeoTag,omitempty"`
			SplitBrainThresholdSeconds *int64             `tfsdk:"split_brain_threshold_seconds" json:"splitBrainThresholdSeconds,omitempty"`
			Type                       *string            `tfsdk:"type" json:"type,omitempty"`
			Weight                     *map[string]string `tfsdk:"weight" json:"weight,omitempty"`
		} `tfsdk:"strategy" json:"strategy,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *K8GbAbsaOssGslbV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_k8gb_absa_oss_gslb_v1beta1_manifest"
}

func (r *K8GbAbsaOssGslbV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Gslb is the Schema for the gslbs API",
		MarkdownDescription: "Gslb is the Schema for the gslbs API",
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
				Description:         "GslbSpec defines the desired state of Gslb",
				MarkdownDescription: "GslbSpec defines the desired state of Gslb",
				Attributes: map[string]schema.Attribute{
					"ingress": schema.SingleNestedAttribute{
						Description:         "Gslb-enabled Ingress Spec",
						MarkdownDescription: "Gslb-enabled Ingress Spec",
						Attributes: map[string]schema.Attribute{
							"backend": schema.SingleNestedAttribute{
								Description:         "A default backend capable of servicing requests that don't match any rule. At least one of 'backend' or 'rules' must be specified. This field is optional to allow the loadbalancer controller or defaulting logic to specify a global default.",
								MarkdownDescription: "A default backend capable of servicing requests that don't match any rule. At least one of 'backend' or 'rules' must be specified. This field is optional to allow the loadbalancer controller or defaulting logic to specify a global default.",
								Attributes: map[string]schema.Attribute{
									"resource": schema.SingleNestedAttribute{
										Description:         "resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, a service.Name and service.Port must not be specified. This is a mutually exclusive setting with 'Service'.",
										MarkdownDescription: "resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, a service.Name and service.Port must not be specified. This is a mutually exclusive setting with 'Service'.",
										Attributes: map[string]schema.Attribute{
											"api_group": schema.StringAttribute{
												Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
												MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
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
										Description:         "service references a service as a backend. This is a mutually exclusive setting with 'Resource'.",
										MarkdownDescription: "service references a service as a backend. This is a mutually exclusive setting with 'Resource'.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "name is the referenced service. The service must exist in the same namespace as the Ingress object.",
												MarkdownDescription: "name is the referenced service. The service must exist in the same namespace as the Ingress object.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"port": schema.SingleNestedAttribute{
												Description:         "port of the referenced service. A port name or port number is required for a IngressServiceBackend.",
												MarkdownDescription: "port of the referenced service. A port name or port number is required for a IngressServiceBackend.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "name is the name of the port on the Service. This is a mutually exclusive setting with 'Number'.",
														MarkdownDescription: "name is the name of the port on the Service. This is a mutually exclusive setting with 'Number'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"number": schema.Int64Attribute{
														Description:         "number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with 'Name'.",
														MarkdownDescription: "number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with 'Name'.",
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
								Description:         "IngressClassName is the name of the IngressClass cluster resource. The associated IngressClass defines which controller will implement the resource. This replaces the deprecated 'kubernetes.io/ingress.class' annotation. For backwards compatibility, when that annotation is set, it must be given precedence over this field. The controller may emit a warning if the field and annotation have different values. Implementations of this API should ignore Ingresses without a class specified. An IngressClass resource may be marked as default, which can be used to set a default value for this field. For more information, refer to the IngressClass documentation.",
								MarkdownDescription: "IngressClassName is the name of the IngressClass cluster resource. The associated IngressClass defines which controller will implement the resource. This replaces the deprecated 'kubernetes.io/ingress.class' annotation. For backwards compatibility, when that annotation is set, it must be given precedence over this field. The controller may emit a warning if the field and annotation have different values. Implementations of this API should ignore Ingresses without a class specified. An IngressClass resource may be marked as default, which can be used to set a default value for this field. For more information, refer to the IngressClass documentation.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rules": schema.ListNestedAttribute{
								Description:         "A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.",
								MarkdownDescription: "A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"host": schema.StringAttribute{
											Description:         "Host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the 'host' part of the URI as defined in RFC 3986: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to the IP in the Spec of the parent Ingress. 2. The ':' delimiter is not respected because ports are not allowed. Currently the port of an Ingress is implicitly :80 for http and :443 for https. Both these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue. Host can be 'precise' which is a domain name without the terminating dot of a network host (e.g. 'foo.bar.com') or 'wildcard', which is a domain name prefixed with a single wildcard label (e.g. '*.foo.com'). The wildcard character '*' must appear by itself as the first DNS label and matches only a single label. You cannot have a wildcard label by itself (e.g. Host == '*'). Requests will be matched against the Host field in the following way: 1. If Host is precise, the request matches this rule if the http host header is equal to Host. 2. If Host is a wildcard, then the request matches this rule if the http host header is to equal to the suffix (removing the first label) of the wildcard rule.",
											MarkdownDescription: "Host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the 'host' part of the URI as defined in RFC 3986: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to the IP in the Spec of the parent Ingress. 2. The ':' delimiter is not respected because ports are not allowed. Currently the port of an Ingress is implicitly :80 for http and :443 for https. Both these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue. Host can be 'precise' which is a domain name without the terminating dot of a network host (e.g. 'foo.bar.com') or 'wildcard', which is a domain name prefixed with a single wildcard label (e.g. '*.foo.com'). The wildcard character '*' must appear by itself as the first DNS label and matches only a single label. You cannot have a wildcard label by itself (e.g. Host == '*'). Requests will be matched against the Host field in the following way: 1. If Host is precise, the request matches this rule if the http host header is equal to Host. 2. If Host is a wildcard, then the request matches this rule if the http host header is to equal to the suffix (removing the first label) of the wildcard rule.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"http": schema.SingleNestedAttribute{
											Description:         "HTTPIngressRuleValue is a list of http selectors pointing to backends. In the example: http://<host>/<path>?<searchpart> -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'.",
											MarkdownDescription: "HTTPIngressRuleValue is a list of http selectors pointing to backends. In the example: http://<host>/<path>?<searchpart> -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'.",
											Attributes: map[string]schema.Attribute{
												"paths": schema.ListNestedAttribute{
													Description:         "paths is a collection of paths that map requests to backends.",
													MarkdownDescription: "paths is a collection of paths that map requests to backends.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"backend": schema.SingleNestedAttribute{
																Description:         "backend defines the referenced service endpoint to which the traffic will be forwarded to.",
																MarkdownDescription: "backend defines the referenced service endpoint to which the traffic will be forwarded to.",
																Attributes: map[string]schema.Attribute{
																	"resource": schema.SingleNestedAttribute{
																		Description:         "resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, a service.Name and service.Port must not be specified. This is a mutually exclusive setting with 'Service'.",
																		MarkdownDescription: "resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, a service.Name and service.Port must not be specified. This is a mutually exclusive setting with 'Service'.",
																		Attributes: map[string]schema.Attribute{
																			"api_group": schema.StringAttribute{
																				Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																				MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
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
																		Description:         "service references a service as a backend. This is a mutually exclusive setting with 'Resource'.",
																		MarkdownDescription: "service references a service as a backend. This is a mutually exclusive setting with 'Resource'.",
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "name is the referenced service. The service must exist in the same namespace as the Ingress object.",
																				MarkdownDescription: "name is the referenced service. The service must exist in the same namespace as the Ingress object.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"port": schema.SingleNestedAttribute{
																				Description:         "port of the referenced service. A port name or port number is required for a IngressServiceBackend.",
																				MarkdownDescription: "port of the referenced service. A port name or port number is required for a IngressServiceBackend.",
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "name is the name of the port on the Service. This is a mutually exclusive setting with 'Number'.",
																						MarkdownDescription: "name is the name of the port on the Service. This is a mutually exclusive setting with 'Number'.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"number": schema.Int64Attribute{
																						Description:         "number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with 'Name'.",
																						MarkdownDescription: "number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with 'Name'.",
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
																Description:         "path is matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986. Paths must begin with a '/' and must be present when using PathType with value 'Exact' or 'Prefix'.",
																MarkdownDescription: "path is matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986. Paths must begin with a '/' and must be present when using PathType with value 'Exact' or 'Prefix'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"path_type": schema.StringAttribute{
																Description:         "pathType determines the interpretation of the path matching. PathType can be one of the following values: * Exact: Matches the URL path exactly. * Prefix: Matches based on a URL path prefix split by '/'. Matching is done on a path element by element basis. A path element refers is the list of labels in the path split by the '/' separator. A request is a match for path p if every p is an element-wise prefix of p of the request path. Note that if the last element of the path is a substring of the last element in request path, it is not a match (e.g. /foo/bar matches /foo/bar/baz, but does not match /foo/barbaz). * ImplementationSpecific: Interpretation of the Path matching is up to the IngressClass. Implementations can treat this as a separate PathType or treat it identically to Prefix or Exact path types. Implementations are required to support all path types.",
																MarkdownDescription: "pathType determines the interpretation of the path matching. PathType can be one of the following values: * Exact: Matches the URL path exactly. * Prefix: Matches based on a URL path prefix split by '/'. Matching is done on a path element by element basis. A path element refers is the list of labels in the path split by the '/' separator. A request is a match for path p if every p is an element-wise prefix of p of the request path. Note that if the last element of the path is a substring of the last element in request path, it is not a match (e.g. /foo/bar matches /foo/bar/baz, but does not match /foo/barbaz). * ImplementationSpecific: Interpretation of the Path matching is up to the IngressClass. Implementations can treat this as a separate PathType or treat it identically to Prefix or Exact path types. Implementations are required to support all path types.",
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
											Required: true,
											Optional: false,
											Computed: false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls": schema.ListNestedAttribute{
								Description:         "TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.",
								MarkdownDescription: "TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"hosts": schema.ListAttribute{
											Description:         "hosts is a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
											MarkdownDescription: "hosts is a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_name": schema.StringAttribute{
											Description:         "secretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the 'Host' header is used for routing.",
											MarkdownDescription: "secretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the 'Host' header is used for routing.",
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

					"resource_ref": schema.SingleNestedAttribute{
						Description:         "ResourceRef spec",
						MarkdownDescription: "ResourceRef spec",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "APIVersion of the referenced resource",
								MarkdownDescription: "APIVersion of the referenced resource",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referenced resource",
								MarkdownDescription: "Kind of the referenced resource",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"strategy": schema.SingleNestedAttribute{
						Description:         "Gslb Strategy spec",
						MarkdownDescription: "Gslb Strategy spec",
						Attributes: map[string]schema.Attribute{
							"dns_ttl_seconds": schema.Int64Attribute{
								Description:         "Defines DNS record TTL in seconds",
								MarkdownDescription: "Defines DNS record TTL in seconds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"primary_geo_tag": schema.StringAttribute{
								Description:         "Primary Geo Tag. Valid for failover strategy only",
								MarkdownDescription: "Primary Geo Tag. Valid for failover strategy only",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"split_brain_threshold_seconds": schema.Int64Attribute{
								Description:         "Split brain TXT record expiration in seconds",
								MarkdownDescription: "Split brain TXT record expiration in seconds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Load balancing strategy type:(roundRobin|failover)",
								MarkdownDescription: "Load balancing strategy type:(roundRobin|failover)",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"weight": schema.MapAttribute{
								Description:         "Weight is defined by map region:weight",
								MarkdownDescription: "Weight is defined by map region:weight",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *K8GbAbsaOssGslbV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_k8gb_absa_oss_gslb_v1beta1_manifest")

	var model K8GbAbsaOssGslbV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("k8gb.absa.oss/v1beta1")
	model.Kind = pointer.String("Gslb")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
