/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package networking_karmada_io_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &NetworkingKarmadaIoMultiClusterIngressV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &NetworkingKarmadaIoMultiClusterIngressV1Alpha1DataSource{}
)

func NewNetworkingKarmadaIoMultiClusterIngressV1Alpha1DataSource() datasource.DataSource {
	return &NetworkingKarmadaIoMultiClusterIngressV1Alpha1DataSource{}
}

type NetworkingKarmadaIoMultiClusterIngressV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type NetworkingKarmadaIoMultiClusterIngressV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *NetworkingKarmadaIoMultiClusterIngressV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_networking_karmada_io_multi_cluster_ingress_v1alpha1"
}

func (r *NetworkingKarmadaIoMultiClusterIngressV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "MultiClusterIngress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend. The structure of MultiClusterIngress is same as Ingress, indicates the Ingress in multi-clusters.",
		MarkdownDescription: "MultiClusterIngress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend. The structure of MultiClusterIngress is same as Ingress, indicates the Ingress in multi-clusters.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "Spec is the desired state of the MultiClusterIngress.",
				MarkdownDescription: "Spec is the desired state of the MultiClusterIngress.",
				Attributes: map[string]schema.Attribute{
					"default_backend": schema.SingleNestedAttribute{
						Description:         "defaultBackend is the backend that should handle requests that don't match any rule. If Rules are not specified, DefaultBackend must be specified. If DefaultBackend is not set, the handling of requests that do not match any of the rules will be up to the Ingress controller.",
						MarkdownDescription: "defaultBackend is the backend that should handle requests that don't match any rule. If Rules are not specified, DefaultBackend must be specified. If DefaultBackend is not set, the handling of requests that do not match any of the rules will be up to the Ingress controller.",
						Attributes: map[string]schema.Attribute{
							"resource": schema.SingleNestedAttribute{
								Description:         "resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, a service.Name and service.Port must not be specified. This is a mutually exclusive setting with 'Service'.",
								MarkdownDescription: "resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, a service.Name and service.Port must not be specified. This is a mutually exclusive setting with 'Service'.",
								Attributes: map[string]schema.Attribute{
									"api_group": schema.StringAttribute{
										Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
										MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"kind": schema.StringAttribute{
										Description:         "Kind is the type of resource being referenced",
										MarkdownDescription: "Kind is the type of resource being referenced",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "Name is the name of resource being referenced",
										MarkdownDescription: "Name is the name of resource being referenced",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"service": schema.SingleNestedAttribute{
								Description:         "service references a service as a backend. This is a mutually exclusive setting with 'Resource'.",
								MarkdownDescription: "service references a service as a backend. This is a mutually exclusive setting with 'Resource'.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "name is the referenced service. The service must exist in the same namespace as the Ingress object.",
										MarkdownDescription: "name is the referenced service. The service must exist in the same namespace as the Ingress object.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"port": schema.SingleNestedAttribute{
										Description:         "port of the referenced service. A port name or port number is required for a IngressServiceBackend.",
										MarkdownDescription: "port of the referenced service. A port name or port number is required for a IngressServiceBackend.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "name is the name of the port on the Service. This is a mutually exclusive setting with 'Number'.",
												MarkdownDescription: "name is the name of the port on the Service. This is a mutually exclusive setting with 'Number'.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"number": schema.Int64Attribute{
												Description:         "number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with 'Name'.",
												MarkdownDescription: "number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with 'Name'.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"ingress_class_name": schema.StringAttribute{
						Description:         "ingressClassName is the name of an IngressClass cluster resource. Ingress controller implementations use this field to know whether they should be serving this Ingress resource, by a transitive connection (controller -> IngressClass -> Ingress resource). Although the 'kubernetes.io/ingress.class' annotation (simple constant name) was never formally defined, it was widely supported by Ingress controllers to create a direct binding between Ingress controller and Ingress resources. Newly created Ingress resources should prefer using the field. However, even though the annotation is officially deprecated, for backwards compatibility reasons, ingress controllers should still honor that annotation if present.",
						MarkdownDescription: "ingressClassName is the name of an IngressClass cluster resource. Ingress controller implementations use this field to know whether they should be serving this Ingress resource, by a transitive connection (controller -> IngressClass -> Ingress resource). Although the 'kubernetes.io/ingress.class' annotation (simple constant name) was never formally defined, it was widely supported by Ingress controllers to create a direct binding between Ingress controller and Ingress resources. Newly created Ingress resources should prefer using the field. However, even though the annotation is officially deprecated, for backwards compatibility reasons, ingress controllers should still honor that annotation if present.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"rules": schema.ListNestedAttribute{
						Description:         "rules is a list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.",
						MarkdownDescription: "rules is a list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"host": schema.StringAttribute{
									Description:         "host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the 'host' part of the URI as defined in RFC 3986: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to the IP in the Spec of the parent Ingress. 2. The ':' delimiter is not respected because ports are not allowed. Currently the port of an Ingress is implicitly :80 for http and :443 for https. Both these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue.  host can be 'precise' which is a domain name without the terminating dot of a network host (e.g. 'foo.bar.com') or 'wildcard', which is a domain name prefixed with a single wildcard label (e.g. '*.foo.com'). The wildcard character '*' must appear by itself as the first DNS label and matches only a single label. You cannot have a wildcard label by itself (e.g. Host == '*'). Requests will be matched against the Host field in the following way: 1. If host is precise, the request matches this rule if the http host header is equal to Host. 2. If host is a wildcard, then the request matches this rule if the http host header is to equal to the suffix (removing the first label) of the wildcard rule.",
									MarkdownDescription: "host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the 'host' part of the URI as defined in RFC 3986: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to the IP in the Spec of the parent Ingress. 2. The ':' delimiter is not respected because ports are not allowed. Currently the port of an Ingress is implicitly :80 for http and :443 for https. Both these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue.  host can be 'precise' which is a domain name without the terminating dot of a network host (e.g. 'foo.bar.com') or 'wildcard', which is a domain name prefixed with a single wildcard label (e.g. '*.foo.com'). The wildcard character '*' must appear by itself as the first DNS label and matches only a single label. You cannot have a wildcard label by itself (e.g. Host == '*'). Requests will be matched against the Host field in the following way: 1. If host is precise, the request matches this rule if the http host header is equal to Host. 2. If host is a wildcard, then the request matches this rule if the http host header is to equal to the suffix (removing the first label) of the wildcard rule.",
									Required:            false,
									Optional:            false,
									Computed:            true,
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
																		Optional:            false,
																		Computed:            true,
																	},

																	"kind": schema.StringAttribute{
																		Description:         "Kind is the type of resource being referenced",
																		MarkdownDescription: "Kind is the type of resource being referenced",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name is the name of resource being referenced",
																		MarkdownDescription: "Name is the name of resource being referenced",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"service": schema.SingleNestedAttribute{
																Description:         "service references a service as a backend. This is a mutually exclusive setting with 'Resource'.",
																MarkdownDescription: "service references a service as a backend. This is a mutually exclusive setting with 'Resource'.",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "name is the referenced service. The service must exist in the same namespace as the Ingress object.",
																		MarkdownDescription: "name is the referenced service. The service must exist in the same namespace as the Ingress object.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"port": schema.SingleNestedAttribute{
																		Description:         "port of the referenced service. A port name or port number is required for a IngressServiceBackend.",
																		MarkdownDescription: "port of the referenced service. A port name or port number is required for a IngressServiceBackend.",
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "name is the name of the port on the Service. This is a mutually exclusive setting with 'Number'.",
																				MarkdownDescription: "name is the name of the port on the Service. This is a mutually exclusive setting with 'Number'.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"number": schema.Int64Attribute{
																				Description:         "number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with 'Name'.",
																				MarkdownDescription: "number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with 'Name'.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"path": schema.StringAttribute{
														Description:         "path is matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986. Paths must begin with a '/' and must be present when using PathType with value 'Exact' or 'Prefix'.",
														MarkdownDescription: "path is matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986. Paths must begin with a '/' and must be present when using PathType with value 'Exact' or 'Prefix'.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"path_type": schema.StringAttribute{
														Description:         "pathType determines the interpretation of the path matching. PathType can be one of the following values: * Exact: Matches the URL path exactly. * Prefix: Matches based on a URL path prefix split by '/'. Matching is done on a path element by element basis. A path element refers is the list of labels in the path split by the '/' separator. A request is a match for path p if every p is an element-wise prefix of p of the request path. Note that if the last element of the path is a substring of the last element in request path, it is not a match (e.g. /foo/bar matches /foo/bar/baz, but does not match /foo/barbaz). * ImplementationSpecific: Interpretation of the Path matching is up to the IngressClass. Implementations can treat this as a separate PathType or treat it identically to Prefix or Exact path types. Implementations are required to support all path types.",
														MarkdownDescription: "pathType determines the interpretation of the path matching. PathType can be one of the following values: * Exact: Matches the URL path exactly. * Prefix: Matches based on a URL path prefix split by '/'. Matching is done on a path element by element basis. A path element refers is the list of labels in the path split by the '/' separator. A request is a match for path p if every p is an element-wise prefix of p of the request path. Note that if the last element of the path is a substring of the last element in request path, it is not a match (e.g. /foo/bar matches /foo/bar/baz, but does not match /foo/barbaz). * ImplementationSpecific: Interpretation of the Path matching is up to the IngressClass. Implementations can treat this as a separate PathType or treat it identically to Prefix or Exact path types. Implementations are required to support all path types.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"tls": schema.ListNestedAttribute{
						Description:         "tls represents the TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.",
						MarkdownDescription: "tls represents the TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"hosts": schema.ListAttribute{
									Description:         "hosts is a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
									MarkdownDescription: "hosts is a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"secret_name": schema.StringAttribute{
									Description:         "secretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the 'Host' header is used for routing.",
									MarkdownDescription: "secretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the 'Host' header is used for routing.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *NetworkingKarmadaIoMultiClusterIngressV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *NetworkingKarmadaIoMultiClusterIngressV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_networking_karmada_io_multi_cluster_ingress_v1alpha1")

	var data NetworkingKarmadaIoMultiClusterIngressV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "networking.karmada.io", Version: "v1alpha1", Resource: "MultiClusterIngress"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse NetworkingKarmadaIoMultiClusterIngressV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("networking.karmada.io/v1alpha1")
	data.Kind = pointer.String("MultiClusterIngress")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
