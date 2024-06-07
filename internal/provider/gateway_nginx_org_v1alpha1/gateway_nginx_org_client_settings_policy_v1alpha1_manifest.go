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
	_ datasource.DataSource = &GatewayNginxOrgClientSettingsPolicyV1Alpha1Manifest{}
)

func NewGatewayNginxOrgClientSettingsPolicyV1Alpha1Manifest() datasource.DataSource {
	return &GatewayNginxOrgClientSettingsPolicyV1Alpha1Manifest{}
}

type GatewayNginxOrgClientSettingsPolicyV1Alpha1Manifest struct{}

type GatewayNginxOrgClientSettingsPolicyV1Alpha1ManifestData struct {
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
		Body *struct {
			MaxSize *string `tfsdk:"max_size" json:"maxSize,omitempty"`
			Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
		} `tfsdk:"body" json:"body,omitempty"`
		KeepAlive *struct {
			Requests *int64  `tfsdk:"requests" json:"requests,omitempty"`
			Time     *string `tfsdk:"time" json:"time,omitempty"`
			Timeout  *struct {
				Header *string `tfsdk:"header" json:"header,omitempty"`
				Server *string `tfsdk:"server" json:"server,omitempty"`
			} `tfsdk:"timeout" json:"timeout,omitempty"`
		} `tfsdk:"keep_alive" json:"keepAlive,omitempty"`
		TargetRef *struct {
			Group *string `tfsdk:"group" json:"group,omitempty"`
			Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GatewayNginxOrgClientSettingsPolicyV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_gateway_nginx_org_client_settings_policy_v1alpha1_manifest"
}

func (r *GatewayNginxOrgClientSettingsPolicyV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClientSettingsPolicy is an Inherited Attached Policy. It provides a way to configure the behavior of the connectionbetween the client and NGINX Gateway Fabric.",
		MarkdownDescription: "ClientSettingsPolicy is an Inherited Attached Policy. It provides a way to configure the behavior of the connectionbetween the client and NGINX Gateway Fabric.",
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
				Description:         "Spec defines the desired state of the ClientSettingsPolicy.",
				MarkdownDescription: "Spec defines the desired state of the ClientSettingsPolicy.",
				Attributes: map[string]schema.Attribute{
					"body": schema.SingleNestedAttribute{
						Description:         "Body defines the client request body settings.",
						MarkdownDescription: "Body defines the client request body settings.",
						Attributes: map[string]schema.Attribute{
							"max_size": schema.StringAttribute{
								Description:         "MaxSize sets the maximum allowed size of the client request body.If the size in a request exceeds the configured value,the 413 (Request Entity Too Large) error is returned to the client.Setting size to 0 disables checking of client request body size.Default: https://nginx.org/en/docs/http/ngx_http_core_module.html#client_max_body_size.",
								MarkdownDescription: "MaxSize sets the maximum allowed size of the client request body.If the size in a request exceeds the configured value,the 413 (Request Entity Too Large) error is returned to the client.Setting size to 0 disables checking of client request body size.Default: https://nginx.org/en/docs/http/ngx_http_core_module.html#client_max_body_size.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d{1,4}(k|m|g)?$`), ""),
								},
							},

							"timeout": schema.StringAttribute{
								Description:         "Timeout defines a timeout for reading client request body. The timeout is set only for a period betweentwo successive read operations, not for the transmission of the whole request body.If a client does not transmit anything within this time, the request is terminated with the408 (Request Time-out) error.Default: https://nginx.org/en/docs/http/ngx_http_core_module.html#client_body_timeout.",
								MarkdownDescription: "Timeout defines a timeout for reading client request body. The timeout is set only for a period betweentwo successive read operations, not for the transmission of the whole request body.If a client does not transmit anything within this time, the request is terminated with the408 (Request Time-out) error.Default: https://nginx.org/en/docs/http/ngx_http_core_module.html#client_body_timeout.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d{1,4}(ms|s)?$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"keep_alive": schema.SingleNestedAttribute{
						Description:         "KeepAlive defines the keep-alive settings.",
						MarkdownDescription: "KeepAlive defines the keep-alive settings.",
						Attributes: map[string]schema.Attribute{
							"requests": schema.Int64Attribute{
								Description:         "Requests sets the maximum number of requests that can be served through one keep-alive connection.After the maximum number of requests are made, the connection is closed. Closing connections periodicallyis necessary to free per-connection memory allocations. Therefore, using too high maximum number of requestsis not recommended as it can lead to excessive memory usage.Default: https://nginx.org/en/docs/http/ngx_http_core_module.html#keepalive_requests.",
								MarkdownDescription: "Requests sets the maximum number of requests that can be served through one keep-alive connection.After the maximum number of requests are made, the connection is closed. Closing connections periodicallyis necessary to free per-connection memory allocations. Therefore, using too high maximum number of requestsis not recommended as it can lead to excessive memory usage.Default: https://nginx.org/en/docs/http/ngx_http_core_module.html#keepalive_requests.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"time": schema.StringAttribute{
								Description:         "Time defines the maximum time during which requests can be processed through one keep-alive connection.After this time is reached, the connection is closed following the subsequent request processing.Default: https://nginx.org/en/docs/http/ngx_http_core_module.html#keepalive_time.",
								MarkdownDescription: "Time defines the maximum time during which requests can be processed through one keep-alive connection.After this time is reached, the connection is closed following the subsequent request processing.Default: https://nginx.org/en/docs/http/ngx_http_core_module.html#keepalive_time.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^\d{1,4}(ms|s)?$`), ""),
								},
							},

							"timeout": schema.SingleNestedAttribute{
								Description:         "Timeout defines the keep-alive timeouts for clients.",
								MarkdownDescription: "Timeout defines the keep-alive timeouts for clients.",
								Attributes: map[string]schema.Attribute{
									"header": schema.StringAttribute{
										Description:         "Header sets the timeout in the 'Keep-Alive: timeout=time' response header field.",
										MarkdownDescription: "Header sets the timeout in the 'Keep-Alive: timeout=time' response header field.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^\d{1,4}(ms|s)?$`), ""),
										},
									},

									"server": schema.StringAttribute{
										Description:         "Server sets the timeout during which a keep-alive client connection will stay open on the server side.Setting this value to 0 disables keep-alive client connections.",
										MarkdownDescription: "Server sets the timeout during which a keep-alive client connection will stay open on the server side.Setting this value to 0 disables keep-alive client connections.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^\d{1,4}(ms|s)?$`), ""),
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

					"target_ref": schema.SingleNestedAttribute{
						Description:         "TargetRef identifies an API object to apply the policy to.Object must be in the same namespace as the policy.Support: Gateway, HTTPRoute, GRPCRoute.",
						MarkdownDescription: "TargetRef identifies an API object to apply the policy to.Object must be in the same namespace as the policy.Support: Gateway, HTTPRoute, GRPCRoute.",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "Group is the group of the target resource.",
								MarkdownDescription: "Group is the group of the target resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
								},
							},

							"kind": schema.StringAttribute{
								Description:         "Kind is kind of the target resource.",
								MarkdownDescription: "Kind is kind of the target resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of the target resource.",
								MarkdownDescription: "Name is the name of the target resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(253),
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
	}
}

func (r *GatewayNginxOrgClientSettingsPolicyV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gateway_nginx_org_client_settings_policy_v1alpha1_manifest")

	var model GatewayNginxOrgClientSettingsPolicyV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("gateway.nginx.org/v1alpha1")
	model.Kind = pointer.String("ClientSettingsPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
