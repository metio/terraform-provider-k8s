/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package gateway_nginx_org_v1alpha1

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
	_ datasource.DataSource = &GatewayNginxOrgNginxProxyV1Alpha1Manifest{}
)

func NewGatewayNginxOrgNginxProxyV1Alpha1Manifest() datasource.DataSource {
	return &GatewayNginxOrgNginxProxyV1Alpha1Manifest{}
}

type GatewayNginxOrgNginxProxyV1Alpha1Manifest struct{}

type GatewayNginxOrgNginxProxyV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		DisableHTTP2    *bool   `tfsdk:"disable_http2" json:"disableHTTP2,omitempty"`
		IpFamily        *string `tfsdk:"ip_family" json:"ipFamily,omitempty"`
		RewriteClientIP *struct {
			Mode             *string `tfsdk:"mode" json:"mode,omitempty"`
			SetIPRecursively *bool   `tfsdk:"set_ip_recursively" json:"setIPRecursively,omitempty"`
			TrustedAddresses *[]struct {
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"trusted_addresses" json:"trustedAddresses,omitempty"`
		} `tfsdk:"rewrite_client_ip" json:"rewriteClientIP,omitempty"`
		Telemetry *struct {
			Exporter *struct {
				BatchCount *int64  `tfsdk:"batch_count" json:"batchCount,omitempty"`
				BatchSize  *int64  `tfsdk:"batch_size" json:"batchSize,omitempty"`
				Endpoint   *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
				Interval   *string `tfsdk:"interval" json:"interval,omitempty"`
			} `tfsdk:"exporter" json:"exporter,omitempty"`
			ServiceName    *string `tfsdk:"service_name" json:"serviceName,omitempty"`
			SpanAttributes *[]struct {
				Key   *string `tfsdk:"key" json:"key,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"span_attributes" json:"spanAttributes,omitempty"`
		} `tfsdk:"telemetry" json:"telemetry,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GatewayNginxOrgNginxProxyV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_gateway_nginx_org_nginx_proxy_v1alpha1_manifest"
}

func (r *GatewayNginxOrgNginxProxyV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "NginxProxy is a configuration object that is attached to a GatewayClass parametersRef. It provides a wayto configure global settings for all Gateways defined from the GatewayClass.",
		MarkdownDescription: "NginxProxy is a configuration object that is attached to a GatewayClass parametersRef. It provides a wayto configure global settings for all Gateways defined from the GatewayClass.",
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
				Description:         "Spec defines the desired state of the NginxProxy.",
				MarkdownDescription: "Spec defines the desired state of the NginxProxy.",
				Attributes: map[string]schema.Attribute{
					"disable_http2": schema.BoolAttribute{
						Description:         "DisableHTTP2 defines if http2 should be disabled for all servers.Default is false, meaning http2 will be enabled for all servers.",
						MarkdownDescription: "DisableHTTP2 defines if http2 should be disabled for all servers.Default is false, meaning http2 will be enabled for all servers.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ip_family": schema.StringAttribute{
						Description:         "IPFamily specifies the IP family to be used by the NGINX.Default is 'dual', meaning the server will use both IPv4 and IPv6.",
						MarkdownDescription: "IPFamily specifies the IP family to be used by the NGINX.Default is 'dual', meaning the server will use both IPv4 and IPv6.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("dual", "ipv4", "ipv6"),
						},
					},

					"rewrite_client_ip": schema.SingleNestedAttribute{
						Description:         "RewriteClientIP defines configuration for rewriting the client IP to the original client's IP.",
						MarkdownDescription: "RewriteClientIP defines configuration for rewriting the client IP to the original client's IP.",
						Attributes: map[string]schema.Attribute{
							"mode": schema.StringAttribute{
								Description:         "Mode defines how NGINX will rewrite the client's IP address.There are two possible modes:- ProxyProtocol: NGINX will rewrite the client's IP using the PROXY protocol header.- XForwardedFor: NGINX will rewrite the client's IP using the X-Forwarded-For header.Sets NGINX directive real_ip_header: https://nginx.org/en/docs/http/ngx_http_realip_module.html#real_ip_header",
								MarkdownDescription: "Mode defines how NGINX will rewrite the client's IP address.There are two possible modes:- ProxyProtocol: NGINX will rewrite the client's IP using the PROXY protocol header.- XForwardedFor: NGINX will rewrite the client's IP using the X-Forwarded-For header.Sets NGINX directive real_ip_header: https://nginx.org/en/docs/http/ngx_http_realip_module.html#real_ip_header",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("ProxyProtocol", "XForwardedFor"),
								},
							},

							"set_ip_recursively": schema.BoolAttribute{
								Description:         "SetIPRecursively configures whether recursive search is used when selecting the client's address fromthe X-Forwarded-For header. It is used in conjunction with TrustedAddresses.If enabled, NGINX will recurse on the values in X-Forwarded-Header from the end of arrayto start of array and select the first untrusted IP.For example, if X-Forwarded-For is [11.11.11.11, 22.22.22.22, 55.55.55.1],and TrustedAddresses is set to 55.55.55.1/32, NGINX will rewrite the client IP to 22.22.22.22.If disabled, NGINX will select the IP at the end of the array.In the previous example, 55.55.55.1 would be selected.Sets NGINX directive real_ip_recursive: https://nginx.org/en/docs/http/ngx_http_realip_module.html#real_ip_recursive",
								MarkdownDescription: "SetIPRecursively configures whether recursive search is used when selecting the client's address fromthe X-Forwarded-For header. It is used in conjunction with TrustedAddresses.If enabled, NGINX will recurse on the values in X-Forwarded-Header from the end of arrayto start of array and select the first untrusted IP.For example, if X-Forwarded-For is [11.11.11.11, 22.22.22.22, 55.55.55.1],and TrustedAddresses is set to 55.55.55.1/32, NGINX will rewrite the client IP to 22.22.22.22.If disabled, NGINX will select the IP at the end of the array.In the previous example, 55.55.55.1 would be selected.Sets NGINX directive real_ip_recursive: https://nginx.org/en/docs/http/ngx_http_realip_module.html#real_ip_recursive",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"trusted_addresses": schema.ListNestedAttribute{
								Description:         "TrustedAddresses specifies the addresses that are trusted to send correct client IP information.If a request comes from a trusted address, NGINX will rewrite the client IP information,and forward it to the backend in the X-Forwarded-For* and X-Real-IP headers.If the request does not come from a trusted address, NGINX will not rewrite the client IP information.TrustedAddresses only supports CIDR blocks: 192.33.21.1/24, fe80::1/64.To trust all addresses (not recommended for production), set to 0.0.0.0/0.If no addresses are provided, NGINX will not rewrite the client IP information.Sets NGINX directive set_real_ip_from: https://nginx.org/en/docs/http/ngx_http_realip_module.html#set_real_ip_fromThis field is required if mode is set.",
								MarkdownDescription: "TrustedAddresses specifies the addresses that are trusted to send correct client IP information.If a request comes from a trusted address, NGINX will rewrite the client IP information,and forward it to the backend in the X-Forwarded-For* and X-Real-IP headers.If the request does not come from a trusted address, NGINX will not rewrite the client IP information.TrustedAddresses only supports CIDR blocks: 192.33.21.1/24, fe80::1/64.To trust all addresses (not recommended for production), set to 0.0.0.0/0.If no addresses are provided, NGINX will not rewrite the client IP information.Sets NGINX directive set_real_ip_from: https://nginx.org/en/docs/http/ngx_http_realip_module.html#set_real_ip_fromThis field is required if mode is set.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"type": schema.StringAttribute{
											Description:         "Type specifies the type of address.Default is 'cidr' which specifies that the address is a CIDR block.",
											MarkdownDescription: "Type specifies the type of address.Default is 'cidr' which specifies that the address is a CIDR block.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("cidr"),
											},
										},

										"value": schema.StringAttribute{
											Description:         "Value specifies the address value.",
											MarkdownDescription: "Value specifies the address value.",
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

					"telemetry": schema.SingleNestedAttribute{
						Description:         "Telemetry specifies the OpenTelemetry configuration.",
						MarkdownDescription: "Telemetry specifies the OpenTelemetry configuration.",
						Attributes: map[string]schema.Attribute{
							"exporter": schema.SingleNestedAttribute{
								Description:         "Exporter specifies OpenTelemetry export parameters.",
								MarkdownDescription: "Exporter specifies OpenTelemetry export parameters.",
								Attributes: map[string]schema.Attribute{
									"batch_count": schema.Int64Attribute{
										Description:         "BatchCount is the number of pending batches per worker, spans exceeding the limit are dropped.Default: https://nginx.org/en/docs/ngx_otel_module.html#otel_exporter",
										MarkdownDescription: "BatchCount is the number of pending batches per worker, spans exceeding the limit are dropped.Default: https://nginx.org/en/docs/ngx_otel_module.html#otel_exporter",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"batch_size": schema.Int64Attribute{
										Description:         "BatchSize is the maximum number of spans to be sent in one batch per worker.Default: https://nginx.org/en/docs/ngx_otel_module.html#otel_exporter",
										MarkdownDescription: "BatchSize is the maximum number of spans to be sent in one batch per worker.Default: https://nginx.org/en/docs/ngx_otel_module.html#otel_exporter",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"endpoint": schema.StringAttribute{
										Description:         "Endpoint is the address of OTLP/gRPC endpoint that will accept telemetry data.Format: alphanumeric hostname with optional http scheme and optional port.",
										MarkdownDescription: "Endpoint is the address of OTLP/gRPC endpoint that will accept telemetry data.Format: alphanumeric hostname with optional http scheme and optional port.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^(?:http?:\/\/)?[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?(?:\.[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?)*(?::\d{1,5})?$`), ""),
										},
									},

									"interval": schema.StringAttribute{
										Description:         "Interval is the maximum interval between two exports.Default: https://nginx.org/en/docs/ngx_otel_module.html#otel_exporter",
										MarkdownDescription: "Interval is the maximum interval between two exports.Default: https://nginx.org/en/docs/ngx_otel_module.html#otel_exporter",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]{1,4}(ms|s|m|h)?$`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_name": schema.StringAttribute{
								Description:         "ServiceName is the 'service.name' attribute of the OpenTelemetry resource.Default is 'ngf:<gateway-namespace>:<gateway-name>'. If a value is provided by the user,then the default becomes a prefix to that value.",
								MarkdownDescription: "ServiceName is the 'service.name' attribute of the OpenTelemetry resource.Default is 'ngf:<gateway-namespace>:<gateway-name>'. If a value is provided by the user,then the default becomes a prefix to that value.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(127),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9_-]+$`), ""),
								},
							},

							"span_attributes": schema.ListNestedAttribute{
								Description:         "SpanAttributes are custom key/value attributes that are added to each span.",
								MarkdownDescription: "SpanAttributes are custom key/value attributes that are added to each span.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "Key is the key for a span attribute.Format: must have all ''' escaped and must not contain any '$' or end with an unescaped ''",
											MarkdownDescription: "Key is the key for a span attribute.Format: must have all ''' escaped and must not contain any '$' or end with an unescaped ''",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(255),
												stringvalidator.RegexMatches(regexp.MustCompile(`^([^"$\\]|\\[^$])*$`), ""),
											},
										},

										"value": schema.StringAttribute{
											Description:         "Value is the value for a span attribute.Format: must have all ''' escaped and must not contain any '$' or end with an unescaped ''",
											MarkdownDescription: "Value is the value for a span attribute.Format: must have all ''' escaped and must not contain any '$' or end with an unescaped ''",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(255),
												stringvalidator.RegexMatches(regexp.MustCompile(`^([^"$\\]|\\[^$])*$`), ""),
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
		},
	}
}

func (r *GatewayNginxOrgNginxProxyV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gateway_nginx_org_nginx_proxy_v1alpha1_manifest")

	var model GatewayNginxOrgNginxProxyV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("gateway.nginx.org/v1alpha1")
	model.Kind = pointer.String("NginxProxy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
