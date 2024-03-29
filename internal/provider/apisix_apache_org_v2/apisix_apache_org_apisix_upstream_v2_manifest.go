/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apisix_apache_org_v2

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
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
	_ datasource.DataSource = &ApisixApacheOrgApisixUpstreamV2Manifest{}
)

func NewApisixApacheOrgApisixUpstreamV2Manifest() datasource.DataSource {
	return &ApisixApacheOrgApisixUpstreamV2Manifest{}
}

type ApisixApacheOrgApisixUpstreamV2Manifest struct{}

type ApisixApacheOrgApisixUpstreamV2ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		Discovery *struct {
			Args        *map[string]string `tfsdk:"args" json:"args,omitempty"`
			ServiceName *string            `tfsdk:"service_name" json:"serviceName,omitempty"`
			Type        *string            `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"discovery" json:"discovery,omitempty"`
		ExternalNodes *[]struct {
			Name   *string `tfsdk:"name" json:"name,omitempty"`
			Port   *int64  `tfsdk:"port" json:"port,omitempty"`
			Type   *string `tfsdk:"type" json:"type,omitempty"`
			Weight *int64  `tfsdk:"weight" json:"weight,omitempty"`
		} `tfsdk:"external_nodes" json:"externalNodes,omitempty"`
		HealthCheck *struct {
			Active *struct {
				Concurrency *int64 `tfsdk:"concurrency" json:"concurrency,omitempty"`
				Healthy     *struct {
					HttpCodes *[]string `tfsdk:"http_codes" json:"httpCodes,omitempty"`
					Interval  *string   `tfsdk:"interval" json:"interval,omitempty"`
					Successes *int64    `tfsdk:"successes" json:"successes,omitempty"`
				} `tfsdk:"healthy" json:"healthy,omitempty"`
				Host           *string   `tfsdk:"host" json:"host,omitempty"`
				HttpPath       *string   `tfsdk:"http_path" json:"httpPath,omitempty"`
				Port           *int64    `tfsdk:"port" json:"port,omitempty"`
				RequestHeaders *[]string `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
				StrictTLS      *bool     `tfsdk:"strict_tls" json:"strictTLS,omitempty"`
				Timeout        *float64  `tfsdk:"timeout" json:"timeout,omitempty"`
				Type           *string   `tfsdk:"type" json:"type,omitempty"`
				Unhealthy      *struct {
					HttpCodes    *[]string `tfsdk:"http_codes" json:"httpCodes,omitempty"`
					HttpFailures *int64    `tfsdk:"http_failures" json:"httpFailures,omitempty"`
					Interval     *string   `tfsdk:"interval" json:"interval,omitempty"`
					TcpFailures  *int64    `tfsdk:"tcp_failures" json:"tcpFailures,omitempty"`
					Timeouts     *int64    `tfsdk:"timeouts" json:"timeouts,omitempty"`
				} `tfsdk:"unhealthy" json:"unhealthy,omitempty"`
			} `tfsdk:"active" json:"active,omitempty"`
			Passive *struct {
				Healthy *struct {
					HttpCodes *[]string `tfsdk:"http_codes" json:"httpCodes,omitempty"`
					Successes *int64    `tfsdk:"successes" json:"successes,omitempty"`
				} `tfsdk:"healthy" json:"healthy,omitempty"`
				Type      *string `tfsdk:"type" json:"type,omitempty"`
				Unhealthy *struct {
					HttpCodes    *[]string `tfsdk:"http_codes" json:"httpCodes,omitempty"`
					HttpFailures *int64    `tfsdk:"http_failures" json:"httpFailures,omitempty"`
					TcpFailures  *int64    `tfsdk:"tcp_failures" json:"tcpFailures,omitempty"`
					Timeouts     *int64    `tfsdk:"timeouts" json:"timeouts,omitempty"`
				} `tfsdk:"unhealthy" json:"unhealthy,omitempty"`
			} `tfsdk:"passive" json:"passive,omitempty"`
		} `tfsdk:"health_check" json:"healthCheck,omitempty"`
		IngressClassName *string `tfsdk:"ingress_class_name" json:"ingressClassName,omitempty"`
		Loadbalancer     *struct {
			HashOn *string `tfsdk:"hash_on" json:"hashOn,omitempty"`
			Key    *string `tfsdk:"key" json:"key,omitempty"`
			Type   *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"loadbalancer" json:"loadbalancer,omitempty"`
		PassHost          *string `tfsdk:"pass_host" json:"passHost,omitempty"`
		PortLevelSettings *[]struct {
			HealthCheck *struct {
				Active *struct {
					Concurrency *int64 `tfsdk:"concurrency" json:"concurrency,omitempty"`
					Healthy     *struct {
						HttpCodes *[]string `tfsdk:"http_codes" json:"httpCodes,omitempty"`
						Interval  *string   `tfsdk:"interval" json:"interval,omitempty"`
						Successes *int64    `tfsdk:"successes" json:"successes,omitempty"`
					} `tfsdk:"healthy" json:"healthy,omitempty"`
					Host           *string   `tfsdk:"host" json:"host,omitempty"`
					HttpPath       *string   `tfsdk:"http_path" json:"httpPath,omitempty"`
					Port           *int64    `tfsdk:"port" json:"port,omitempty"`
					RequestHeaders *[]string `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
					StrictTLS      *bool     `tfsdk:"strict_tls" json:"strictTLS,omitempty"`
					Timeout        *float64  `tfsdk:"timeout" json:"timeout,omitempty"`
					Type           *string   `tfsdk:"type" json:"type,omitempty"`
					Unhealthy      *struct {
						HttpCodes    *[]string `tfsdk:"http_codes" json:"httpCodes,omitempty"`
						HttpFailures *int64    `tfsdk:"http_failures" json:"httpFailures,omitempty"`
						Interval     *string   `tfsdk:"interval" json:"interval,omitempty"`
						TcpFailures  *int64    `tfsdk:"tcp_failures" json:"tcpFailures,omitempty"`
						Timeout      *string   `tfsdk:"timeout" json:"timeout,omitempty"`
					} `tfsdk:"unhealthy" json:"unhealthy,omitempty"`
				} `tfsdk:"active" json:"active,omitempty"`
				Passive *struct {
					Healthy *struct {
						HttpCodes *[]string `tfsdk:"http_codes" json:"httpCodes,omitempty"`
						Successes *int64    `tfsdk:"successes" json:"successes,omitempty"`
					} `tfsdk:"healthy" json:"healthy,omitempty"`
					Type      *string `tfsdk:"type" json:"type,omitempty"`
					Unhealthy *struct {
						HttpCodes    *[]string `tfsdk:"http_codes" json:"httpCodes,omitempty"`
						HttpFailures *int64    `tfsdk:"http_failures" json:"httpFailures,omitempty"`
						TcpFailures  *int64    `tfsdk:"tcp_failures" json:"tcpFailures,omitempty"`
						Timeout      *string   `tfsdk:"timeout" json:"timeout,omitempty"`
					} `tfsdk:"unhealthy" json:"unhealthy,omitempty"`
				} `tfsdk:"passive" json:"passive,omitempty"`
			} `tfsdk:"health_check" json:"healthCheck,omitempty"`
			Loadbalancer *struct {
				HashOn *string `tfsdk:"hash_on" json:"hashOn,omitempty"`
				Key    *string `tfsdk:"key" json:"key,omitempty"`
				Type   *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"loadbalancer" json:"loadbalancer,omitempty"`
			Port    *int64  `tfsdk:"port" json:"port,omitempty"`
			Retries *int64  `tfsdk:"retries" json:"retries,omitempty"`
			Scheme  *string `tfsdk:"scheme" json:"scheme,omitempty"`
			Timeout *struct {
				Connect *string `tfsdk:"connect" json:"connect,omitempty"`
				Read    *string `tfsdk:"read" json:"read,omitempty"`
				Send    *string `tfsdk:"send" json:"send,omitempty"`
			} `tfsdk:"timeout" json:"timeout,omitempty"`
		} `tfsdk:"port_level_settings" json:"portLevelSettings,omitempty"`
		Retries *int64  `tfsdk:"retries" json:"retries,omitempty"`
		Scheme  *string `tfsdk:"scheme" json:"scheme,omitempty"`
		Subsets *[]struct {
			Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Name   *string            `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"subsets" json:"subsets,omitempty"`
		Timeout *struct {
			Connect *string `tfsdk:"connect" json:"connect,omitempty"`
			Read    *string `tfsdk:"read" json:"read,omitempty"`
			Send    *string `tfsdk:"send" json:"send,omitempty"`
		} `tfsdk:"timeout" json:"timeout,omitempty"`
		TlsSecret *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"tls_secret" json:"tlsSecret,omitempty"`
		UpstreamHost *string `tfsdk:"upstream_host" json:"upstreamHost,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ApisixApacheOrgApisixUpstreamV2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apisix_apache_org_apisix_upstream_v2_manifest"
}

func (r *ApisixApacheOrgApisixUpstreamV2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"discovery": schema.SingleNestedAttribute{
						Description:         "Discovery is used to configure service discovery for upstream",
						MarkdownDescription: "Discovery is used to configure service discovery for upstream",
						Attributes: map[string]schema.Attribute{
							"args": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"external_nodes": schema.ListNestedAttribute{
						Description:         "ExternalNodes contains external nodes the Upstream should use If this field is set, the upstream will use these nodes directly without any further resolves",
						MarkdownDescription: "ExternalNodes contains external nodes the Upstream should use If this field is set, the upstream will use these nodes directly without any further resolves",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"port": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"weight": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
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

					"health_check": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"active": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"concurrency": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"healthy": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"http_codes": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"interval": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"successes": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(254),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"host": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^\*?[0-9a-zA-Z-._]+$`), ""),
										},
									},

									"http_path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},

									"port": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
											int64validator.AtMost(65535),
										},
									},

									"request_headers": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"strict_tls": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"timeout": schema.Float64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Float64{
											float64validator.AtLeast(0),
										},
									},

									"type": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("http", "https", "tcp"),
										},
									},

									"unhealthy": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"http_codes": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_failures": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(254),
												},
											},

											"interval": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tcp_failures": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(254),
												},
											},

											"timeouts": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
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

							"passive": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"healthy": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"http_codes": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"successes": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(254),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("http", "https", "tcp"),
										},
									},

									"unhealthy": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"http_codes": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_failures": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(254),
												},
											},

											"tcp_failures": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(254),
												},
											},

											"timeouts": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress_class_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"loadbalancer": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"hash_on": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("vars", "vars_combinations", "header", "cookie", "consumer"),
								},
							},

							"key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("roundrobin", "chash", "ewma", "least_conn"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"pass_host": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("pass", "node", "rewrite"),
						},
					},

					"port_level_settings": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"health_check": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"active": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"concurrency": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
													},
												},

												"healthy": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"http_codes": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"interval": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"successes": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
																int64validator.AtMost(254),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"host": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^\*?[0-9a-zA-Z-._]+$`), ""),
													},
												},

												"http_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
													},
												},

												"port": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
														int64validator.AtMost(65535),
													},
												},

												"request_headers": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"strict_tls": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"timeout": schema.Float64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Float64{
														float64validator.AtLeast(0),
													},
												},

												"type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("http", "https", "tcp"),
													},
												},

												"unhealthy": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"http_codes": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"http_failures": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
																int64validator.AtMost(254),
															},
														},

														"interval": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"tcp_failures": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
																int64validator.AtMost(254),
															},
														},

														"timeout": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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

										"passive": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"healthy": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"http_codes": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"successes": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
																int64validator.AtMost(254),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("http", "https", "tcp"),
													},
												},

												"unhealthy": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"http_codes": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"http_failures": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
																int64validator.AtMost(254),
															},
														},

														"tcp_failures": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
																int64validator.AtMost(254),
															},
														},

														"timeout": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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

								"loadbalancer": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"hash_on": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("vars", "vars_combinations", "header", "cookie", "consumer"),
											},
										},

										"key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("roundrobin", "chash", "ewma", "least_conn"),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"port": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(1),
										int64validator.AtMost(65535),
									},
								},

								"retries": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
									},
								},

								"scheme": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("http", "grpc"),
									},
								},

								"timeout": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"connect": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"read": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"send": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"retries": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"scheme": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("http", "grpc", "https", "grpcs"),
						},
					},

					"subsets": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"labels": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"timeout": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"connect": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"read": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"send": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tls_secret": schema.SingleNestedAttribute{
						Description:         "ApisixSecret describes the Kubernetes Secret name and namespace.",
						MarkdownDescription: "ApisixSecret describes the Kubernetes Secret name and namespace.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"namespace": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
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

					"upstream_host": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^\*?[0-9a-zA-Z-._]+$`), ""),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *ApisixApacheOrgApisixUpstreamV2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apisix_apache_org_apisix_upstream_v2_manifest")

	var model ApisixApacheOrgApisixUpstreamV2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("apisix.apache.org/v2")
	model.Kind = pointer.String("ApisixUpstream")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
