/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package networking_istio_io_v1beta1

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
	_ datasource.DataSource              = &NetworkingIstioIoSidecarV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &NetworkingIstioIoSidecarV1Beta1DataSource{}
)

func NewNetworkingIstioIoSidecarV1Beta1DataSource() datasource.DataSource {
	return &NetworkingIstioIoSidecarV1Beta1DataSource{}
}

type NetworkingIstioIoSidecarV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type NetworkingIstioIoSidecarV1Beta1DataSourceData struct {
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
		Egress *[]struct {
			Bind        *string   `tfsdk:"bind" json:"bind,omitempty"`
			CaptureMode *string   `tfsdk:"capture_mode" json:"captureMode,omitempty"`
			Hosts       *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
			Port        *struct {
				Name       *string `tfsdk:"name" json:"name,omitempty"`
				Number     *int64  `tfsdk:"number" json:"number,omitempty"`
				Protocol   *string `tfsdk:"protocol" json:"protocol,omitempty"`
				TargetPort *int64  `tfsdk:"target_port" json:"targetPort,omitempty"`
			} `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"egress" json:"egress,omitempty"`
		Ingress *[]struct {
			Bind            *string `tfsdk:"bind" json:"bind,omitempty"`
			CaptureMode     *string `tfsdk:"capture_mode" json:"captureMode,omitempty"`
			DefaultEndpoint *string `tfsdk:"default_endpoint" json:"defaultEndpoint,omitempty"`
			Port            *struct {
				Name       *string `tfsdk:"name" json:"name,omitempty"`
				Number     *int64  `tfsdk:"number" json:"number,omitempty"`
				Protocol   *string `tfsdk:"protocol" json:"protocol,omitempty"`
				TargetPort *int64  `tfsdk:"target_port" json:"targetPort,omitempty"`
			} `tfsdk:"port" json:"port,omitempty"`
			Tls *struct {
				CaCertificates        *string   `tfsdk:"ca_certificates" json:"caCertificates,omitempty"`
				CipherSuites          *[]string `tfsdk:"cipher_suites" json:"cipherSuites,omitempty"`
				CredentialName        *string   `tfsdk:"credential_name" json:"credentialName,omitempty"`
				HttpsRedirect         *bool     `tfsdk:"https_redirect" json:"httpsRedirect,omitempty"`
				MaxProtocolVersion    *string   `tfsdk:"max_protocol_version" json:"maxProtocolVersion,omitempty"`
				MinProtocolVersion    *string   `tfsdk:"min_protocol_version" json:"minProtocolVersion,omitempty"`
				Mode                  *string   `tfsdk:"mode" json:"mode,omitempty"`
				PrivateKey            *string   `tfsdk:"private_key" json:"privateKey,omitempty"`
				ServerCertificate     *string   `tfsdk:"server_certificate" json:"serverCertificate,omitempty"`
				SubjectAltNames       *[]string `tfsdk:"subject_alt_names" json:"subjectAltNames,omitempty"`
				VerifyCertificateHash *[]string `tfsdk:"verify_certificate_hash" json:"verifyCertificateHash,omitempty"`
				VerifyCertificateSpki *[]string `tfsdk:"verify_certificate_spki" json:"verifyCertificateSpki,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"ingress" json:"ingress,omitempty"`
		OutboundTrafficPolicy *struct {
			EgressProxy *struct {
				Host *string `tfsdk:"host" json:"host,omitempty"`
				Port *struct {
					Number *int64 `tfsdk:"number" json:"number,omitempty"`
				} `tfsdk:"port" json:"port,omitempty"`
				Subset *string `tfsdk:"subset" json:"subset,omitempty"`
			} `tfsdk:"egress_proxy" json:"egressProxy,omitempty"`
			Mode *string `tfsdk:"mode" json:"mode,omitempty"`
		} `tfsdk:"outbound_traffic_policy" json:"outboundTrafficPolicy,omitempty"`
		WorkloadSelector *struct {
			Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"workload_selector" json:"workloadSelector,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NetworkingIstioIoSidecarV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_networking_istio_io_sidecar_v1beta1"
}

func (r *NetworkingIstioIoSidecarV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "Configuration affecting network reachability of a sidecar. See more details at: https://istio.io/docs/reference/config/networking/sidecar.html",
				MarkdownDescription: "Configuration affecting network reachability of a sidecar. See more details at: https://istio.io/docs/reference/config/networking/sidecar.html",
				Attributes: map[string]schema.Attribute{
					"egress": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"bind": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"capture_mode": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"hosts": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"port": schema.SingleNestedAttribute{
									Description:         "The port associated with the listener.",
									MarkdownDescription: "The port associated with the listener.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Label assigned to the port.",
											MarkdownDescription: "Label assigned to the port.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"number": schema.Int64Attribute{
											Description:         "A valid non-negative integer port number.",
											MarkdownDescription: "A valid non-negative integer port number.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"protocol": schema.StringAttribute{
											Description:         "The protocol exposed on the port.",
											MarkdownDescription: "The protocol exposed on the port.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"target_port": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"ingress": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"bind": schema.StringAttribute{
									Description:         "The IP(IPv4 or IPv6) to which the listener should be bound.",
									MarkdownDescription: "The IP(IPv4 or IPv6) to which the listener should be bound.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"capture_mode": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"default_endpoint": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"port": schema.SingleNestedAttribute{
									Description:         "The port associated with the listener.",
									MarkdownDescription: "The port associated with the listener.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Label assigned to the port.",
											MarkdownDescription: "Label assigned to the port.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"number": schema.Int64Attribute{
											Description:         "A valid non-negative integer port number.",
											MarkdownDescription: "A valid non-negative integer port number.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"protocol": schema.StringAttribute{
											Description:         "The protocol exposed on the port.",
											MarkdownDescription: "The protocol exposed on the port.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"target_port": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"tls": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"ca_certificates": schema.StringAttribute{
											Description:         "REQUIRED if mode is 'MUTUAL' or 'OPTIONAL_MUTUAL'.",
											MarkdownDescription: "REQUIRED if mode is 'MUTUAL' or 'OPTIONAL_MUTUAL'.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"cipher_suites": schema.ListAttribute{
											Description:         "Optional: If specified, only support the specified cipher list.",
											MarkdownDescription: "Optional: If specified, only support the specified cipher list.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"credential_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"https_redirect": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"max_protocol_version": schema.StringAttribute{
											Description:         "Optional: Maximum TLS protocol version.",
											MarkdownDescription: "Optional: Maximum TLS protocol version.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"min_protocol_version": schema.StringAttribute{
											Description:         "Optional: Minimum TLS protocol version.",
											MarkdownDescription: "Optional: Minimum TLS protocol version.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mode": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"private_key": schema.StringAttribute{
											Description:         "REQUIRED if mode is 'SIMPLE' or 'MUTUAL'.",
											MarkdownDescription: "REQUIRED if mode is 'SIMPLE' or 'MUTUAL'.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"server_certificate": schema.StringAttribute{
											Description:         "REQUIRED if mode is 'SIMPLE' or 'MUTUAL'.",
											MarkdownDescription: "REQUIRED if mode is 'SIMPLE' or 'MUTUAL'.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"subject_alt_names": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"verify_certificate_hash": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"verify_certificate_spki": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"outbound_traffic_policy": schema.SingleNestedAttribute{
						Description:         "Configuration for the outbound traffic policy.",
						MarkdownDescription: "Configuration for the outbound traffic policy.",
						Attributes: map[string]schema.Attribute{
							"egress_proxy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										Description:         "The name of a service from the service registry.",
										MarkdownDescription: "The name of a service from the service registry.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"port": schema.SingleNestedAttribute{
										Description:         "Specifies the port on the host that is being addressed.",
										MarkdownDescription: "Specifies the port on the host that is being addressed.",
										Attributes: map[string]schema.Attribute{
											"number": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"subset": schema.StringAttribute{
										Description:         "The name of a subset within the service.",
										MarkdownDescription: "The name of a subset within the service.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"mode": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"workload_selector": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"labels": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
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
	}
}

func (r *NetworkingIstioIoSidecarV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *NetworkingIstioIoSidecarV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_networking_istio_io_sidecar_v1beta1")

	var data NetworkingIstioIoSidecarV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "networking.istio.io", Version: "v1beta1", Resource: "Sidecar"}).
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

	var readResponse NetworkingIstioIoSidecarV1Beta1DataSourceData
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
	data.ApiVersion = pointer.String("networking.istio.io/v1beta1")
	data.Kind = pointer.String("Sidecar")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}