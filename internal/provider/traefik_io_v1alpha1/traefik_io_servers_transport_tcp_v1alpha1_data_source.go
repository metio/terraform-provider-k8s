/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package traefik_io_v1alpha1

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
	_ datasource.DataSource              = &TraefikIoServersTransportTcpV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &TraefikIoServersTransportTcpV1Alpha1DataSource{}
)

func NewTraefikIoServersTransportTcpV1Alpha1DataSource() datasource.DataSource {
	return &TraefikIoServersTransportTcpV1Alpha1DataSource{}
}

type TraefikIoServersTransportTcpV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type TraefikIoServersTransportTcpV1Alpha1DataSourceData struct {
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

func (r *TraefikIoServersTransportTcpV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_traefik_io_servers_transport_tcp_v1alpha1"
}

func (r *TraefikIoServersTransportTcpV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "ServersTransportTCPSpec defines the desired state of a ServersTransportTCP.",
				MarkdownDescription: "ServersTransportTCPSpec defines the desired state of a ServersTransportTCP.",
				Attributes: map[string]schema.Attribute{
					"dial_keep_alive": schema.StringAttribute{
						Description:         "DialKeepAlive is the interval between keep-alive probes for an active network connection. If zero, keep-alive probes are sent with a default value (currently 15 seconds), if supported by the protocol and operating system. Network protocols or operating systems that do not support keep-alives ignore this field. If negative, keep-alive probes are disabled.",
						MarkdownDescription: "DialKeepAlive is the interval between keep-alive probes for an active network connection. If zero, keep-alive probes are sent with a default value (currently 15 seconds), if supported by the protocol and operating system. Network protocols or operating systems that do not support keep-alives ignore this field. If negative, keep-alive probes are disabled.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"dial_timeout": schema.StringAttribute{
						Description:         "DialTimeout is the amount of time to wait until a connection to a backend server can be established.",
						MarkdownDescription: "DialTimeout is the amount of time to wait until a connection to a backend server can be established.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"termination_delay": schema.StringAttribute{
						Description:         "TerminationDelay defines the delay to wait before fully terminating the connection, after one connected peer has closed its writing capability.",
						MarkdownDescription: "TerminationDelay defines the delay to wait before fully terminating the connection, after one connected peer has closed its writing capability.",
						Required:            false,
						Optional:            false,
						Computed:            true,
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
								Optional:            false,
								Computed:            true,
							},

							"insecure_skip_verify": schema.BoolAttribute{
								Description:         "InsecureSkipVerify disables TLS certificate verification.",
								MarkdownDescription: "InsecureSkipVerify disables TLS certificate verification.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"peer_cert_uri": schema.StringAttribute{
								Description:         "MaxIdleConnsPerHost controls the maximum idle (keep-alive) to keep per-host. PeerCertURI defines the peer cert URI used to match against SAN URI during the peer certificate verification.",
								MarkdownDescription: "MaxIdleConnsPerHost controls the maximum idle (keep-alive) to keep per-host. PeerCertURI defines the peer cert URI used to match against SAN URI during the peer certificate verification.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"root_c_as_secrets": schema.ListAttribute{
								Description:         "RootCAsSecrets defines a list of CA secret used to validate self-signed certificates.",
								MarkdownDescription: "RootCAsSecrets defines a list of CA secret used to validate self-signed certificates.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"server_name": schema.StringAttribute{
								Description:         "ServerName defines the server name used to contact the server.",
								MarkdownDescription: "ServerName defines the server name used to contact the server.",
								Required:            false,
								Optional:            false,
								Computed:            true,
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
										Optional:            false,
										Computed:            true,
									},

									"trust_domain": schema.StringAttribute{
										Description:         "TrustDomain defines the allowed SPIFFE trust domain.",
										MarkdownDescription: "TrustDomain defines the allowed SPIFFE trust domain.",
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
		},
	}
}

func (r *TraefikIoServersTransportTcpV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *TraefikIoServersTransportTcpV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_traefik_io_servers_transport_tcp_v1alpha1")

	var data TraefikIoServersTransportTcpV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "traefik.io", Version: "v1alpha1", Resource: "serverstransporttcps"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse TraefikIoServersTransportTcpV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("traefik.io/v1alpha1")
	data.Kind = pointer.String("ServersTransportTCP")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
