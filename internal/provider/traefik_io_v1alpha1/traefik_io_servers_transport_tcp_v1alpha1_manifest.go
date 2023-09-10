/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package traefik_io_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &TraefikIoServersTransportTcpV1Alpha1Manifest{}
)

func NewTraefikIoServersTransportTcpV1Alpha1Manifest() datasource.DataSource {
	return &TraefikIoServersTransportTcpV1Alpha1Manifest{}
}

type TraefikIoServersTransportTcpV1Alpha1Manifest struct{}

type TraefikIoServersTransportTcpV1Alpha1ManifestData struct {
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
		DialKeepAlive    *string `tfsdk:"dial_keep_alive" json:"dialKeepAlive,omitempty"`
		DialTimeout      *string `tfsdk:"dial_timeout" json:"dialTimeout,omitempty"`
		TerminationDelay *string `tfsdk:"termination_delay" json:"terminationDelay,omitempty"`
		Tls              *struct {
			CertificatesSecrets *[]string `tfsdk:"certificates_secrets" json:"certificatesSecrets,omitempty"`
			InsecureSkipVerify  *bool     `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
			PeerCertURI         *string   `tfsdk:"peer_cert_uri" json:"peerCertURI,omitempty"`
			RootCAsSecrets      *[]string `tfsdk:"root_c_as_secrets" json:"rootCAsSecrets,omitempty"`
			ServerName          *string   `tfsdk:"server_name" json:"serverName,omitempty"`
			Spiffe              *struct {
				Ids         *[]string `tfsdk:"ids" json:"ids,omitempty"`
				TrustDomain *string   `tfsdk:"trust_domain" json:"trustDomain,omitempty"`
			} `tfsdk:"spiffe" json:"spiffe,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TraefikIoServersTransportTcpV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_traefik_io_servers_transport_tcp_v1alpha1_manifest"
}

func (r *TraefikIoServersTransportTcpV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ServersTransportTCP is the CRD implementation of a TCPServersTransport. If no tcpServersTransport is specified, a default one named default@internal will be used. The default@internal tcpServersTransport can be configured in the static configuration. More info: https://doc.traefik.io/traefik/v3.0/routing/services/#serverstransport_3",
		MarkdownDescription: "ServersTransportTCP is the CRD implementation of a TCPServersTransport. If no tcpServersTransport is specified, a default one named default@internal will be used. The default@internal tcpServersTransport can be configured in the static configuration. More info: https://doc.traefik.io/traefik/v3.0/routing/services/#serverstransport_3",
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
				Description:         "ServersTransportTCPSpec defines the desired state of a ServersTransportTCP.",
				MarkdownDescription: "ServersTransportTCPSpec defines the desired state of a ServersTransportTCP.",
				Attributes: map[string]schema.Attribute{
					"dial_keep_alive": schema.StringAttribute{
						Description:         "DialKeepAlive is the interval between keep-alive probes for an active network connection. If zero, keep-alive probes are sent with a default value (currently 15 seconds), if supported by the protocol and operating system. Network protocols or operating systems that do not support keep-alives ignore this field. If negative, keep-alive probes are disabled.",
						MarkdownDescription: "DialKeepAlive is the interval between keep-alive probes for an active network connection. If zero, keep-alive probes are sent with a default value (currently 15 seconds), if supported by the protocol and operating system. Network protocols or operating systems that do not support keep-alives ignore this field. If negative, keep-alive probes are disabled.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"dial_timeout": schema.StringAttribute{
						Description:         "DialTimeout is the amount of time to wait until a connection to a backend server can be established.",
						MarkdownDescription: "DialTimeout is the amount of time to wait until a connection to a backend server can be established.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"termination_delay": schema.StringAttribute{
						Description:         "TerminationDelay defines the delay to wait before fully terminating the connection, after one connected peer has closed its writing capability.",
						MarkdownDescription: "TerminationDelay defines the delay to wait before fully terminating the connection, after one connected peer has closed its writing capability.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tls": schema.SingleNestedAttribute{
						Description:         "TLS defines the TLS configuration",
						MarkdownDescription: "TLS defines the TLS configuration",
						Attributes: map[string]schema.Attribute{
							"certificates_secrets": schema.ListAttribute{
								Description:         "CertificatesSecrets defines a list of secret storing client certificates for mTLS.",
								MarkdownDescription: "CertificatesSecrets defines a list of secret storing client certificates for mTLS.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"insecure_skip_verify": schema.BoolAttribute{
								Description:         "InsecureSkipVerify disables TLS certificate verification.",
								MarkdownDescription: "InsecureSkipVerify disables TLS certificate verification.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"peer_cert_uri": schema.StringAttribute{
								Description:         "MaxIdleConnsPerHost controls the maximum idle (keep-alive) to keep per-host. PeerCertURI defines the peer cert URI used to match against SAN URI during the peer certificate verification.",
								MarkdownDescription: "MaxIdleConnsPerHost controls the maximum idle (keep-alive) to keep per-host. PeerCertURI defines the peer cert URI used to match against SAN URI during the peer certificate verification.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"root_c_as_secrets": schema.ListAttribute{
								Description:         "RootCAsSecrets defines a list of CA secret used to validate self-signed certificates.",
								MarkdownDescription: "RootCAsSecrets defines a list of CA secret used to validate self-signed certificates.",
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

func (r *TraefikIoServersTransportTcpV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_traefik_io_servers_transport_tcp_v1alpha1_manifest")

	var model TraefikIoServersTransportTcpV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("traefik.io/v1alpha1")
	model.Kind = pointer.String("ServersTransportTCP")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
