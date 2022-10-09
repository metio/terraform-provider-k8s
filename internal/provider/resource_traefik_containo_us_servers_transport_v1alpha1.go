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

type TraefikContainoUsServersTransportV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*TraefikContainoUsServersTransportV1Alpha1Resource)(nil)
)

type TraefikContainoUsServersTransportV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type TraefikContainoUsServersTransportV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		CertificatesSecrets *[]string `tfsdk:"certificates_secrets" yaml:"certificatesSecrets,omitempty"`

		DisableHTTP2 *bool `tfsdk:"disable_http2" yaml:"disableHTTP2,omitempty"`

		ForwardingTimeouts *struct {
			DialTimeout *string `tfsdk:"dial_timeout" yaml:"dialTimeout,omitempty"`

			IdleConnTimeout *string `tfsdk:"idle_conn_timeout" yaml:"idleConnTimeout,omitempty"`

			PingTimeout *string `tfsdk:"ping_timeout" yaml:"pingTimeout,omitempty"`

			ReadIdleTimeout *string `tfsdk:"read_idle_timeout" yaml:"readIdleTimeout,omitempty"`

			ResponseHeaderTimeout *string `tfsdk:"response_header_timeout" yaml:"responseHeaderTimeout,omitempty"`
		} `tfsdk:"forwarding_timeouts" yaml:"forwardingTimeouts,omitempty"`

		InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

		MaxIdleConnsPerHost *int64 `tfsdk:"max_idle_conns_per_host" yaml:"maxIdleConnsPerHost,omitempty"`

		PeerCertURI *string `tfsdk:"peer_cert_uri" yaml:"peerCertURI,omitempty"`

		RootCAsSecrets *[]string `tfsdk:"root_c_as_secrets" yaml:"rootCAsSecrets,omitempty"`

		ServerName *string `tfsdk:"server_name" yaml:"serverName,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewTraefikContainoUsServersTransportV1Alpha1Resource() resource.Resource {
	return &TraefikContainoUsServersTransportV1Alpha1Resource{}
}

func (r *TraefikContainoUsServersTransportV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_traefik_containo_us_servers_transport_v1alpha1"
}

func (r *TraefikContainoUsServersTransportV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ServersTransport is the CRD implementation of a ServersTransport. If no serversTransport is specified, the default@internal will be used. The default@internal serversTransport is created from the static configuration. More info: https://doc.traefik.io/traefik/v2.9/routing/services/#serverstransport_1",
		MarkdownDescription: "ServersTransport is the CRD implementation of a ServersTransport. If no serversTransport is specified, the default@internal will be used. The default@internal serversTransport is created from the static configuration. More info: https://doc.traefik.io/traefik/v2.9/routing/services/#serverstransport_1",
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
				Description:         "ServersTransportSpec defines the desired state of a ServersTransport.",
				MarkdownDescription: "ServersTransportSpec defines the desired state of a ServersTransport.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"certificates_secrets": {
						Description:         "CertificatesSecrets defines a list of secret storing client certificates for mTLS.",
						MarkdownDescription: "CertificatesSecrets defines a list of secret storing client certificates for mTLS.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"disable_http2": {
						Description:         "DisableHTTP2 disables HTTP/2 for connections with backend servers.",
						MarkdownDescription: "DisableHTTP2 disables HTTP/2 for connections with backend servers.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"forwarding_timeouts": {
						Description:         "ForwardingTimeouts defines the timeouts for requests forwarded to the backend servers.",
						MarkdownDescription: "ForwardingTimeouts defines the timeouts for requests forwarded to the backend servers.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"dial_timeout": {
								Description:         "DialTimeout is the amount of time to wait until a connection to a backend server can be established.",
								MarkdownDescription: "DialTimeout is the amount of time to wait until a connection to a backend server can be established.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"idle_conn_timeout": {
								Description:         "IdleConnTimeout is the maximum period for which an idle HTTP keep-alive connection will remain open before closing itself.",
								MarkdownDescription: "IdleConnTimeout is the maximum period for which an idle HTTP keep-alive connection will remain open before closing itself.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ping_timeout": {
								Description:         "PingTimeout is the timeout after which the HTTP/2 connection will be closed if a response to ping is not received.",
								MarkdownDescription: "PingTimeout is the timeout after which the HTTP/2 connection will be closed if a response to ping is not received.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"read_idle_timeout": {
								Description:         "ReadIdleTimeout is the timeout after which a health check using ping frame will be carried out if no frame is received on the HTTP/2 connection.",
								MarkdownDescription: "ReadIdleTimeout is the timeout after which a health check using ping frame will be carried out if no frame is received on the HTTP/2 connection.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"response_header_timeout": {
								Description:         "ResponseHeaderTimeout is the amount of time to wait for a server's response headers after fully writing the request (including its body, if any).",
								MarkdownDescription: "ResponseHeaderTimeout is the amount of time to wait for a server's response headers after fully writing the request (including its body, if any).",

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

					"insecure_skip_verify": {
						Description:         "InsecureSkipVerify disables SSL certificate verification.",
						MarkdownDescription: "InsecureSkipVerify disables SSL certificate verification.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"max_idle_conns_per_host": {
						Description:         "MaxIdleConnsPerHost controls the maximum idle (keep-alive) to keep per-host.",
						MarkdownDescription: "MaxIdleConnsPerHost controls the maximum idle (keep-alive) to keep per-host.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"peer_cert_uri": {
						Description:         "PeerCertURI defines the peer cert URI used to match against SAN URI during the peer certificate verification.",
						MarkdownDescription: "PeerCertURI defines the peer cert URI used to match against SAN URI during the peer certificate verification.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"root_c_as_secrets": {
						Description:         "RootCAsSecrets defines a list of CA secret used to validate self-signed certificate.",
						MarkdownDescription: "RootCAsSecrets defines a list of CA secret used to validate self-signed certificate.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"server_name": {
						Description:         "ServerName defines the server name used to contact the server.",
						MarkdownDescription: "ServerName defines the server name used to contact the server.",

						Type: types.StringType,

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

func (r *TraefikContainoUsServersTransportV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_traefik_containo_us_servers_transport_v1alpha1")

	var state TraefikContainoUsServersTransportV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel TraefikContainoUsServersTransportV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("traefik.containo.us/v1alpha1")
	goModel.Kind = utilities.Ptr("ServersTransport")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *TraefikContainoUsServersTransportV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_traefik_containo_us_servers_transport_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *TraefikContainoUsServersTransportV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_traefik_containo_us_servers_transport_v1alpha1")

	var state TraefikContainoUsServersTransportV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel TraefikContainoUsServersTransportV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("traefik.containo.us/v1alpha1")
	goModel.Kind = utilities.Ptr("ServersTransport")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *TraefikContainoUsServersTransportV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_traefik_containo_us_servers_transport_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
