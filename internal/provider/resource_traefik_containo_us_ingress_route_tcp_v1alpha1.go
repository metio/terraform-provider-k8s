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

type TraefikContainoUsIngressRouteTCPV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*TraefikContainoUsIngressRouteTCPV1Alpha1Resource)(nil)
)

type TraefikContainoUsIngressRouteTCPV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type TraefikContainoUsIngressRouteTCPV1Alpha1GoModel struct {
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
		EntryPoints *[]string `tfsdk:"entry_points" yaml:"entryPoints,omitempty"`

		Routes *[]struct {
			Match *string `tfsdk:"match" yaml:"match,omitempty"`

			Middlewares *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"middlewares" yaml:"middlewares,omitempty"`

			Priority *int64 `tfsdk:"priority" yaml:"priority,omitempty"`

			Services *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

				Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

				ProxyProtocol *struct {
					Version *int64 `tfsdk:"version" yaml:"version,omitempty"`
				} `tfsdk:"proxy_protocol" yaml:"proxyProtocol,omitempty"`

				TerminationDelay *int64 `tfsdk:"termination_delay" yaml:"terminationDelay,omitempty"`

				Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
			} `tfsdk:"services" yaml:"services,omitempty"`
		} `tfsdk:"routes" yaml:"routes,omitempty"`

		Tls *struct {
			CertResolver *string `tfsdk:"cert_resolver" yaml:"certResolver,omitempty"`

			Domains *[]struct {
				Main *string `tfsdk:"main" yaml:"main,omitempty"`

				Sans *[]string `tfsdk:"sans" yaml:"sans,omitempty"`
			} `tfsdk:"domains" yaml:"domains,omitempty"`

			Options *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"options" yaml:"options,omitempty"`

			Passthrough *bool `tfsdk:"passthrough" yaml:"passthrough,omitempty"`

			SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`

			Store *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"store" yaml:"store,omitempty"`
		} `tfsdk:"tls" yaml:"tls,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewTraefikContainoUsIngressRouteTCPV1Alpha1Resource() resource.Resource {
	return &TraefikContainoUsIngressRouteTCPV1Alpha1Resource{}
}

func (r *TraefikContainoUsIngressRouteTCPV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_traefik_containo_us_ingress_route_tcp_v1alpha1"
}

func (r *TraefikContainoUsIngressRouteTCPV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "IngressRouteTCP is the CRD implementation of a Traefik TCP Router.",
		MarkdownDescription: "IngressRouteTCP is the CRD implementation of a Traefik TCP Router.",
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
				Description:         "IngressRouteTCPSpec defines the desired state of IngressRouteTCP.",
				MarkdownDescription: "IngressRouteTCPSpec defines the desired state of IngressRouteTCP.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"entry_points": {
						Description:         "EntryPoints defines the list of entry point names to bind to. Entry points have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v2.9/routing/entrypoints/ Default: all.",
						MarkdownDescription: "EntryPoints defines the list of entry point names to bind to. Entry points have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v2.9/routing/entrypoints/ Default: all.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"routes": {
						Description:         "Routes defines the list of routes.",
						MarkdownDescription: "Routes defines the list of routes.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"match": {
								Description:         "Match defines the router's rule. More info: https://doc.traefik.io/traefik/v2.9/routing/routers/#rule_1",
								MarkdownDescription: "Match defines the router's rule. More info: https://doc.traefik.io/traefik/v2.9/routing/routers/#rule_1",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"middlewares": {
								Description:         "Middlewares defines the list of references to MiddlewareTCP resources.",
								MarkdownDescription: "Middlewares defines the list of references to MiddlewareTCP resources.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name defines the name of the referenced Traefik resource.",
										MarkdownDescription: "Name defines the name of the referenced Traefik resource.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"namespace": {
										Description:         "Namespace defines the namespace of the referenced Traefik resource.",
										MarkdownDescription: "Namespace defines the namespace of the referenced Traefik resource.",

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

							"priority": {
								Description:         "Priority defines the router's priority. More info: https://doc.traefik.io/traefik/v2.9/routing/routers/#priority_1",
								MarkdownDescription: "Priority defines the router's priority. More info: https://doc.traefik.io/traefik/v2.9/routing/routers/#priority_1",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"services": {
								Description:         "Services defines the list of TCP services.",
								MarkdownDescription: "Services defines the list of TCP services.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name defines the name of the referenced Kubernetes Service.",
										MarkdownDescription: "Name defines the name of the referenced Kubernetes Service.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"namespace": {
										Description:         "Namespace defines the namespace of the referenced Kubernetes Service.",
										MarkdownDescription: "Namespace defines the namespace of the referenced Kubernetes Service.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "Port defines the port of a Kubernetes Service. This can be a reference to a named port.",
										MarkdownDescription: "Port defines the port of a Kubernetes Service. This can be a reference to a named port.",

										Type: utilities.IntOrStringType{},

										Required: true,
										Optional: false,
										Computed: false,
									},

									"proxy_protocol": {
										Description:         "ProxyProtocol defines the PROXY protocol configuration. More info: https://doc.traefik.io/traefik/v2.9/routing/services/#proxy-protocol",
										MarkdownDescription: "ProxyProtocol defines the PROXY protocol configuration. More info: https://doc.traefik.io/traefik/v2.9/routing/services/#proxy-protocol",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"version": {
												Description:         "Version defines the PROXY Protocol version to use.",
												MarkdownDescription: "Version defines the PROXY Protocol version to use.",

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

									"termination_delay": {
										Description:         "TerminationDelay defines the deadline that the proxy sets, after one of its connected peers indicates it has closed the writing capability of its connection, to close the reading capability as well, hence fully terminating the connection. It is a duration in milliseconds, defaulting to 100. A negative value means an infinite deadline (i.e. the reading capability is never closed).",
										MarkdownDescription: "TerminationDelay defines the deadline that the proxy sets, after one of its connected peers indicates it has closed the writing capability of its connection, to close the reading capability as well, hence fully terminating the connection. It is a duration in milliseconds, defaulting to 100. A negative value means an infinite deadline (i.e. the reading capability is never closed).",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"weight": {
										Description:         "Weight defines the weight used when balancing requests between multiple Kubernetes Service.",
										MarkdownDescription: "Weight defines the weight used when balancing requests between multiple Kubernetes Service.",

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

						Required: true,
						Optional: false,
						Computed: false,
					},

					"tls": {
						Description:         "TLS defines the TLS configuration on a layer 4 / TCP Route. More info: https://doc.traefik.io/traefik/v2.9/routing/routers/#tls_1",
						MarkdownDescription: "TLS defines the TLS configuration on a layer 4 / TCP Route. More info: https://doc.traefik.io/traefik/v2.9/routing/routers/#tls_1",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cert_resolver": {
								Description:         "CertResolver defines the name of the certificate resolver to use. Cert resolvers have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v2.9/https/acme/#certificate-resolvers",
								MarkdownDescription: "CertResolver defines the name of the certificate resolver to use. Cert resolvers have to be configured in the static configuration. More info: https://doc.traefik.io/traefik/v2.9/https/acme/#certificate-resolvers",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"domains": {
								Description:         "Domains defines the list of domains that will be used to issue certificates. More info: https://doc.traefik.io/traefik/v2.9/routing/routers/#domains",
								MarkdownDescription: "Domains defines the list of domains that will be used to issue certificates. More info: https://doc.traefik.io/traefik/v2.9/routing/routers/#domains",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"main": {
										Description:         "Main defines the main domain name.",
										MarkdownDescription: "Main defines the main domain name.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sans": {
										Description:         "SANs defines the subject alternative domain names.",
										MarkdownDescription: "SANs defines the subject alternative domain names.",

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

							"options": {
								Description:         "Options defines the reference to a TLSOption, that specifies the parameters of the TLS connection. If not defined, the 'default' TLSOption is used. More info: https://doc.traefik.io/traefik/v2.9/https/tls/#tls-options",
								MarkdownDescription: "Options defines the reference to a TLSOption, that specifies the parameters of the TLS connection. If not defined, the 'default' TLSOption is used. More info: https://doc.traefik.io/traefik/v2.9/https/tls/#tls-options",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name defines the name of the referenced Traefik resource.",
										MarkdownDescription: "Name defines the name of the referenced Traefik resource.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"namespace": {
										Description:         "Namespace defines the namespace of the referenced Traefik resource.",
										MarkdownDescription: "Namespace defines the namespace of the referenced Traefik resource.",

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

							"passthrough": {
								Description:         "Passthrough defines whether a TLS router will terminate the TLS connection.",
								MarkdownDescription: "Passthrough defines whether a TLS router will terminate the TLS connection.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"secret_name": {
								Description:         "SecretName is the name of the referenced Kubernetes Secret to specify the certificate details.",
								MarkdownDescription: "SecretName is the name of the referenced Kubernetes Secret to specify the certificate details.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"store": {
								Description:         "Store defines the reference to the TLSStore, that will be used to store certificates. Please note that only 'default' TLSStore can be used.",
								MarkdownDescription: "Store defines the reference to the TLSStore, that will be used to store certificates. Please note that only 'default' TLSStore can be used.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name defines the name of the referenced Traefik resource.",
										MarkdownDescription: "Name defines the name of the referenced Traefik resource.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"namespace": {
										Description:         "Namespace defines the namespace of the referenced Traefik resource.",
										MarkdownDescription: "Namespace defines the namespace of the referenced Traefik resource.",

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
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *TraefikContainoUsIngressRouteTCPV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_traefik_containo_us_ingress_route_tcp_v1alpha1")

	var state TraefikContainoUsIngressRouteTCPV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel TraefikContainoUsIngressRouteTCPV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("traefik.containo.us/v1alpha1")
	goModel.Kind = utilities.Ptr("IngressRouteTCP")

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

func (r *TraefikContainoUsIngressRouteTCPV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_traefik_containo_us_ingress_route_tcp_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *TraefikContainoUsIngressRouteTCPV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_traefik_containo_us_ingress_route_tcp_v1alpha1")

	var state TraefikContainoUsIngressRouteTCPV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel TraefikContainoUsIngressRouteTCPV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("traefik.containo.us/v1alpha1")
	goModel.Kind = utilities.Ptr("IngressRouteTCP")

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

func (r *TraefikContainoUsIngressRouteTCPV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_traefik_containo_us_ingress_route_tcp_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
