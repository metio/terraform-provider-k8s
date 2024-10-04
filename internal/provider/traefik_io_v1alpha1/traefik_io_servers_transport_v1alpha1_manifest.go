/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package traefik_io_v1alpha1

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
	_ datasource.DataSource = &TraefikIoServersTransportV1Alpha1Manifest{}
)

func NewTraefikIoServersTransportV1Alpha1Manifest() datasource.DataSource {
	return &TraefikIoServersTransportV1Alpha1Manifest{}
}

type TraefikIoServersTransportV1Alpha1Manifest struct{}

type TraefikIoServersTransportV1Alpha1ManifestData struct {
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
		CertificatesSecrets *[]string `tfsdk:"certificates_secrets" json:"certificatesSecrets,omitempty"`
		DisableHTTP2        *bool     `tfsdk:"disable_http2" json:"disableHTTP2,omitempty"`
		ForwardingTimeouts  *struct {
			DialTimeout           *string `tfsdk:"dial_timeout" json:"dialTimeout,omitempty"`
			IdleConnTimeout       *string `tfsdk:"idle_conn_timeout" json:"idleConnTimeout,omitempty"`
			PingTimeout           *string `tfsdk:"ping_timeout" json:"pingTimeout,omitempty"`
			ReadIdleTimeout       *string `tfsdk:"read_idle_timeout" json:"readIdleTimeout,omitempty"`
			ResponseHeaderTimeout *string `tfsdk:"response_header_timeout" json:"responseHeaderTimeout,omitempty"`
		} `tfsdk:"forwarding_timeouts" json:"forwardingTimeouts,omitempty"`
		InsecureSkipVerify  *bool     `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
		MaxIdleConnsPerHost *int64    `tfsdk:"max_idle_conns_per_host" json:"maxIdleConnsPerHost,omitempty"`
		PeerCertURI         *string   `tfsdk:"peer_cert_uri" json:"peerCertURI,omitempty"`
		RootCAsSecrets      *[]string `tfsdk:"root_c_as_secrets" json:"rootCAsSecrets,omitempty"`
		ServerName          *string   `tfsdk:"server_name" json:"serverName,omitempty"`
		Spiffe              *struct {
			Ids         *[]string `tfsdk:"ids" json:"ids,omitempty"`
			TrustDomain *string   `tfsdk:"trust_domain" json:"trustDomain,omitempty"`
		} `tfsdk:"spiffe" json:"spiffe,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TraefikIoServersTransportV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_traefik_io_servers_transport_v1alpha1_manifest"
}

func (r *TraefikIoServersTransportV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ServersTransport is the CRD implementation of a ServersTransport. If no serversTransport is specified, the default@internal will be used. The default@internal serversTransport is created from the static configuration. More info: https://doc.traefik.io/traefik/v3.2/routing/services/#serverstransport_1",
		MarkdownDescription: "ServersTransport is the CRD implementation of a ServersTransport. If no serversTransport is specified, the default@internal will be used. The default@internal serversTransport is created from the static configuration. More info: https://doc.traefik.io/traefik/v3.2/routing/services/#serverstransport_1",
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
				Description:         "ServersTransportSpec defines the desired state of a ServersTransport.",
				MarkdownDescription: "ServersTransportSpec defines the desired state of a ServersTransport.",
				Attributes: map[string]schema.Attribute{
					"certificates_secrets": schema.ListAttribute{
						Description:         "CertificatesSecrets defines a list of secret storing client certificates for mTLS.",
						MarkdownDescription: "CertificatesSecrets defines a list of secret storing client certificates for mTLS.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_http2": schema.BoolAttribute{
						Description:         "DisableHTTP2 disables HTTP/2 for connections with backend servers.",
						MarkdownDescription: "DisableHTTP2 disables HTTP/2 for connections with backend servers.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"forwarding_timeouts": schema.SingleNestedAttribute{
						Description:         "ForwardingTimeouts defines the timeouts for requests forwarded to the backend servers.",
						MarkdownDescription: "ForwardingTimeouts defines the timeouts for requests forwarded to the backend servers.",
						Attributes: map[string]schema.Attribute{
							"dial_timeout": schema.StringAttribute{
								Description:         "DialTimeout is the amount of time to wait until a connection to a backend server can be established.",
								MarkdownDescription: "DialTimeout is the amount of time to wait until a connection to a backend server can be established.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"idle_conn_timeout": schema.StringAttribute{
								Description:         "IdleConnTimeout is the maximum period for which an idle HTTP keep-alive connection will remain open before closing itself.",
								MarkdownDescription: "IdleConnTimeout is the maximum period for which an idle HTTP keep-alive connection will remain open before closing itself.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ping_timeout": schema.StringAttribute{
								Description:         "PingTimeout is the timeout after which the HTTP/2 connection will be closed if a response to ping is not received.",
								MarkdownDescription: "PingTimeout is the timeout after which the HTTP/2 connection will be closed if a response to ping is not received.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"read_idle_timeout": schema.StringAttribute{
								Description:         "ReadIdleTimeout is the timeout after which a health check using ping frame will be carried out if no frame is received on the HTTP/2 connection.",
								MarkdownDescription: "ReadIdleTimeout is the timeout after which a health check using ping frame will be carried out if no frame is received on the HTTP/2 connection.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"response_header_timeout": schema.StringAttribute{
								Description:         "ResponseHeaderTimeout is the amount of time to wait for a server's response headers after fully writing the request (including its body, if any).",
								MarkdownDescription: "ResponseHeaderTimeout is the amount of time to wait for a server's response headers after fully writing the request (including its body, if any).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"insecure_skip_verify": schema.BoolAttribute{
						Description:         "InsecureSkipVerify disables SSL certificate verification.",
						MarkdownDescription: "InsecureSkipVerify disables SSL certificate verification.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_idle_conns_per_host": schema.Int64Attribute{
						Description:         "MaxIdleConnsPerHost controls the maximum idle (keep-alive) to keep per-host.",
						MarkdownDescription: "MaxIdleConnsPerHost controls the maximum idle (keep-alive) to keep per-host.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"peer_cert_uri": schema.StringAttribute{
						Description:         "PeerCertURI defines the peer cert URI used to match against SAN URI during the peer certificate verification.",
						MarkdownDescription: "PeerCertURI defines the peer cert URI used to match against SAN URI during the peer certificate verification.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"root_c_as_secrets": schema.ListAttribute{
						Description:         "RootCAsSecrets defines a list of CA secret used to validate self-signed certificate.",
						MarkdownDescription: "RootCAsSecrets defines a list of CA secret used to validate self-signed certificate.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"server_name": schema.StringAttribute{
						Description:         "ServerName defines the server name used to contact the server.",
						MarkdownDescription: "ServerName defines the server name used to contact the server.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"spiffe": schema.SingleNestedAttribute{
						Description:         "Spiffe defines the SPIFFE configuration.",
						MarkdownDescription: "Spiffe defines the SPIFFE configuration.",
						Attributes: map[string]schema.Attribute{
							"ids": schema.ListAttribute{
								Description:         "IDs defines the allowed SPIFFE IDs (takes precedence over the SPIFFE TrustDomain).",
								MarkdownDescription: "IDs defines the allowed SPIFFE IDs (takes precedence over the SPIFFE TrustDomain).",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"trust_domain": schema.StringAttribute{
								Description:         "TrustDomain defines the allowed SPIFFE trust domain.",
								MarkdownDescription: "TrustDomain defines the allowed SPIFFE trust domain.",
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *TraefikIoServersTransportV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_traefik_io_servers_transport_v1alpha1_manifest")

	var model TraefikIoServersTransportV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("traefik.io/v1alpha1")
	model.Kind = pointer.String("ServersTransport")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
