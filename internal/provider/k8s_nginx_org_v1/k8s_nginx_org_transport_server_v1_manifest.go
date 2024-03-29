/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package k8s_nginx_org_v1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &K8SNginxOrgTransportServerV1Manifest{}
)

func NewK8SNginxOrgTransportServerV1Manifest() datasource.DataSource {
	return &K8SNginxOrgTransportServerV1Manifest{}
}

type K8SNginxOrgTransportServerV1Manifest struct{}

type K8SNginxOrgTransportServerV1ManifestData struct {
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
		Action *struct {
			Pass *string `tfsdk:"pass" json:"pass,omitempty"`
		} `tfsdk:"action" json:"action,omitempty"`
		Host             *string `tfsdk:"host" json:"host,omitempty"`
		IngressClassName *string `tfsdk:"ingress_class_name" json:"ingressClassName,omitempty"`
		Listener         *struct {
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
		} `tfsdk:"listener" json:"listener,omitempty"`
		ServerSnippets    *string `tfsdk:"server_snippets" json:"serverSnippets,omitempty"`
		SessionParameters *struct {
			Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
		} `tfsdk:"session_parameters" json:"sessionParameters,omitempty"`
		StreamSnippets *string `tfsdk:"stream_snippets" json:"streamSnippets,omitempty"`
		Tls            *struct {
			Secret *string `tfsdk:"secret" json:"secret,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
		UpstreamParameters *struct {
			ConnectTimeout      *string `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
			NextUpstream        *bool   `tfsdk:"next_upstream" json:"nextUpstream,omitempty"`
			NextUpstreamTimeout *string `tfsdk:"next_upstream_timeout" json:"nextUpstreamTimeout,omitempty"`
			NextUpstreamTries   *int64  `tfsdk:"next_upstream_tries" json:"nextUpstreamTries,omitempty"`
			UdpRequests         *int64  `tfsdk:"udp_requests" json:"udpRequests,omitempty"`
			UdpResponses        *int64  `tfsdk:"udp_responses" json:"udpResponses,omitempty"`
		} `tfsdk:"upstream_parameters" json:"upstreamParameters,omitempty"`
		Upstreams *[]struct {
			Backup      *string `tfsdk:"backup" json:"backup,omitempty"`
			BackupPort  *int64  `tfsdk:"backup_port" json:"backupPort,omitempty"`
			FailTimeout *string `tfsdk:"fail_timeout" json:"failTimeout,omitempty"`
			HealthCheck *struct {
				Enable   *bool   `tfsdk:"enable" json:"enable,omitempty"`
				Fails    *int64  `tfsdk:"fails" json:"fails,omitempty"`
				Interval *string `tfsdk:"interval" json:"interval,omitempty"`
				Jitter   *string `tfsdk:"jitter" json:"jitter,omitempty"`
				Match    *struct {
					Expect *string `tfsdk:"expect" json:"expect,omitempty"`
					Send   *string `tfsdk:"send" json:"send,omitempty"`
				} `tfsdk:"match" json:"match,omitempty"`
				Passes  *int64  `tfsdk:"passes" json:"passes,omitempty"`
				Port    *int64  `tfsdk:"port" json:"port,omitempty"`
				Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"health_check" json:"healthCheck,omitempty"`
			LoadBalancingMethod *string `tfsdk:"load_balancing_method" json:"loadBalancingMethod,omitempty"`
			MaxConns            *int64  `tfsdk:"max_conns" json:"maxConns,omitempty"`
			MaxFails            *int64  `tfsdk:"max_fails" json:"maxFails,omitempty"`
			Name                *string `tfsdk:"name" json:"name,omitempty"`
			Port                *int64  `tfsdk:"port" json:"port,omitempty"`
			Service             *string `tfsdk:"service" json:"service,omitempty"`
		} `tfsdk:"upstreams" json:"upstreams,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *K8SNginxOrgTransportServerV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_k8s_nginx_org_transport_server_v1_manifest"
}

func (r *K8SNginxOrgTransportServerV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TransportServer defines the TransportServer resource.",
		MarkdownDescription: "TransportServer defines the TransportServer resource.",
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
				Description:         "TransportServerSpec is the spec of the TransportServer resource.",
				MarkdownDescription: "TransportServerSpec is the spec of the TransportServer resource.",
				Attributes: map[string]schema.Attribute{
					"action": schema.SingleNestedAttribute{
						Description:         "TransportServerAction defines an action.",
						MarkdownDescription: "TransportServerAction defines an action.",
						Attributes: map[string]schema.Attribute{
							"pass": schema.StringAttribute{
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

					"host": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ingress_class_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"listener": schema.SingleNestedAttribute{
						Description:         "TransportServerListener defines a listener for a TransportServer.",
						MarkdownDescription: "TransportServerListener defines a listener for a TransportServer.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"protocol": schema.StringAttribute{
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

					"server_snippets": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"session_parameters": schema.SingleNestedAttribute{
						Description:         "SessionParameters defines session parameters.",
						MarkdownDescription: "SessionParameters defines session parameters.",
						Attributes: map[string]schema.Attribute{
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

					"stream_snippets": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tls": schema.SingleNestedAttribute{
						Description:         "TransportServerTLS defines TransportServerTLS configuration for a TransportServer.",
						MarkdownDescription: "TransportServerTLS defines TransportServerTLS configuration for a TransportServer.",
						Attributes: map[string]schema.Attribute{
							"secret": schema.StringAttribute{
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

					"upstream_parameters": schema.SingleNestedAttribute{
						Description:         "UpstreamParameters defines parameters for an upstream.",
						MarkdownDescription: "UpstreamParameters defines parameters for an upstream.",
						Attributes: map[string]schema.Attribute{
							"connect_timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"next_upstream": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"next_upstream_timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"next_upstream_tries": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"udp_requests": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"udp_responses": schema.Int64Attribute{
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

					"upstreams": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"backup": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"backup_port": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"fail_timeout": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"health_check": schema.SingleNestedAttribute{
									Description:         "TransportServerHealthCheck defines the parameters for active Upstream HealthChecks.",
									MarkdownDescription: "TransportServerHealthCheck defines the parameters for active Upstream HealthChecks.",
									Attributes: map[string]schema.Attribute{
										"enable": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"fails": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
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

										"jitter": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"match": schema.SingleNestedAttribute{
											Description:         "TransportServerMatch defines the parameters of a custom health check.",
											MarkdownDescription: "TransportServerMatch defines the parameters of a custom health check.",
											Attributes: map[string]schema.Attribute{
												"expect": schema.StringAttribute{
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

										"passes": schema.Int64Attribute{
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

								"load_balancing_method": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_conns": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_fails": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

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

								"service": schema.StringAttribute{
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *K8SNginxOrgTransportServerV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_k8s_nginx_org_transport_server_v1_manifest")

	var model K8SNginxOrgTransportServerV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("k8s.nginx.org/v1")
	model.Kind = pointer.String("TransportServer")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
